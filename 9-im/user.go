package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	addr := conn.RemoteAddr().String()

	user := &User{
		Name:   addr,
		Addr:   addr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage(conn)
	return user
}

func (u *User) ListenMessage(conn net.Conn) {
	for {
		msg := <-u.C
		conn.Write([]byte(msg + "\n"))
	}
}

func (u *User) online() {

	u.server.MapLock.Lock()
	u.server.UserList[u.Name] = u
	u.server.MapLock.Unlock()

	u.server.broadcast(u, "online")
}

func (u *User) offline() {
	u.server.MapLock.Lock()
	delete(u.server.UserList, u.Name)
	u.server.MapLock.Unlock()

	u.server.broadcast(u, "offline")
}

func (u *User) DoMessage(msg string) {
	if msg == "who" {
		u.server.MapLock.Lock()
		for _, user := range u.server.UserList {
			u.SendMsg(user.Name + " is on line\n")
		}
		u.server.MapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]
		if _, ok := u.server.UserList[newName]; ok {
			u.SendMsg("username exist!")
		} else {
			u.server.MapLock.Lock()
			delete(u.server.UserList, u.Name)
			u.server.UserList[newName] = u
			u.server.MapLock.Unlock()

			u.Name = newName
			u.SendMsg("username rename success")
		}
	} else if len(msg) > 3 && msg[:3] == "to|" {
		fmt.Println(msg)
		args := strings.Split(msg, "|")
		if len(args) != 3 {
			fmt.Println("input format error, format to|name|msg")
		}
		toName, msg := args[1], args[2]
		if len(msg) == 0 {
			fmt.Println("please input message")
		}
		toUser, ok := u.server.UserList[toName]
		if !ok {
			fmt.Println("input name error, can not find user")
		}
		toUser.SendMsg(u.Name + ": " + msg + "\n")
	} else {
		u.server.broadcast(u, msg)

	}

}

func (u *User) SendMsg(s string) {
	u.conn.Write([]byte(s))
}
