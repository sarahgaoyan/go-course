package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	serverIP   string
	serverPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIP string, serverPort int) *Client {
	client := &Client{
		serverIP:   serverIP,
		serverPort: serverPort,
		flag:       999,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("conn dial error: ", err)
		return nil
	}
	client.conn = conn
	return client
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "server ip")
	flag.IntVar(&serverPort, "port", 8888, "server port")
}

func (c *Client) menu() bool {
	fmt.Println("Please select the mode(0-3):")
	fmt.Println("1. rename")
	fmt.Println("2. public chat")
	fmt.Println("3. private chat")
	fmt.Println("0. exit")

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("invalid input, input 0-3")
		return false
	}

	if choice < 0 || choice > 3 {
		fmt.Println("invalid input, input 0-3")
		return false
	}
	c.flag = choice
	return true
}

func (c *Client) receiveMsg() {
	io.Copy(os.Stdout, c.conn)
}

func (c *Client) rename() bool {
	fmt.Println("please input the name:")
	fmt.Scanln(&c.Name)
	msg := "rename|" + c.Name + "\n"
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("update name error : ", err)
		return false
	}
	return true
}

func (c *Client) publicChat() {
	fmt.Println("please input the message, or exit for exit")
	var msg string
	fmt.Scanln(&msg)
	for msg != "exit" {
		_, err := c.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("send message error : ", err)
		}

		msg = ""
		fmt.Scanln(&msg)
	}
}

func (c *Client) privateChat() {
	fmt.Println("please select the name, or exit for exit")
	c.findOnlineUsers()
	var toUser string
	fmt.Scanln(&toUser)

	for toUser != "exit" {
		fmt.Println("please input message, or exit for exit")
		var msg string
		fmt.Scanln(&msg)
		for msg != "exit" {
			pMsg := "to|" + toUser + "|" + msg + "\n"
			_, err := c.conn.Write([]byte(pMsg))
			if err != nil {
				fmt.Println("send message error : ", err)
				break
			}

			msg = ""
			fmt.Println("please input message, or exit for exit")
			fmt.Scanln(&msg)
		}

		fmt.Println("please select the name, or exit for exit")
		c.findOnlineUsers()
		fmt.Scanln(&toUser)
	}
}

func (c *Client) run() {
	for c.flag != 0 {
		for c.menu() {
			switch c.flag {
			case 1:
				c.rename()
				break
			case 2:
				c.publicChat()
				break
			case 3:
				c.privateChat()
				break
			default:
				break
			}
		}

	}
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("connection to server failed")
		return
	}

	go client.receiveMsg()

	fmt.Println("connection to server succeed")

	client.run()
}

func (c *Client) findOnlineUsers() {
	msg := "who\n"
	c.conn.Write([]byte(msg))
}
