package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	UserList map[string]*User
	MapLock  sync.RWMutex

	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:       ip,
		Port:     port,
		UserList: make(map[string]*User),
		MapLock:  sync.RWMutex{},
		Message:  make(chan string),
	}
	return server
}

func (s *Server) start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("listen err: ", err)
	}

	defer listen.Close()

	go s.ListenMessage()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Print("conn error: ", err)
			continue
		}

		go s.handler(conn)
	}
}

func (s *Server) handler(conn net.Conn) {
	user := NewUser(conn, s)
	user.online()

	var isLive = make(chan bool)

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read error: ", err)
				return
			}
			msg := string(buf[:n-1])
			user.DoMessage(msg)
			isLive <- true
		}
	}()

	for {
		select {
		case <-isLive:
		case <-time.After(time.Second * 300):
			user.SendMsg("timeout\n")
			close(user.C)
			conn.Close()
			return
		}
	}
}

func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message
		s.MapLock.Lock()
		for _, user := range s.UserList {
			user.C <- msg
		}
		s.MapLock.Unlock()
	}
}

func (s *Server) broadcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + msg
	s.Message <- sendMsg
}
