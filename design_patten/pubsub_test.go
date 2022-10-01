package design_patten

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

type (
	subscriber chan interface{}
	topicFun   func(v interface{}) bool
)

type Publisher struct {
	m           sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFun
}

func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFun),
	}
}

func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic topicFun) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

func (p *Publisher) Publish(v interface{}) {
	p.m.RLocker()
	defer p.m.RLocker()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

func (p *Publisher) sendTopic(sub subscriber, topic topicFun, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}

}

func TestPubSub(t *testing.T) {
	t.Run("test pub sub patten", func(t *testing.T) {
		p := NewPublisher(100*time.Millisecond, 10)
		defer p.Close()

		all := p.Subscribe()
		golang := p.SubscribeTopic(func(v interface{}) bool {
			if s, ok := v.(string); ok {
				return strings.Contains(s, "golang")
			}
			return false
		})

		p.Publish("hello world")
		p.Publish("hello golang")

		go func() {
			for msg := range all {
				fmt.Println("all: ", msg)
			}
		}()

		go func() {
			for msg := range golang {
				fmt.Println("golang: ", msg)
			}
		}()

		time.Sleep(3 * time.Second)
	})
}
