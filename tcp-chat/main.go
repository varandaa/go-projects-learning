package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type clientMessage struct {
	clientConn net.Conn
	msg        string
}

var wg = sync.WaitGroup{}
var clients []net.Conn = make([]net.Conn, 0, 5)
var c = make(chan clientMessage)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Server started on port 8000")
	wg.Add(1)
	go broadcastMessage(c)
	for {
		conn, err := ln.Accept()
		fmt.Printf("Ip %v connected!\n", conn.RemoteAddr().String())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		clients = append(clients, conn)
		wg.Add(1)
		go clientHandler(conn, c)
	}
}

func clientHandler(clientConn net.Conn, c chan clientMessage) {
	for {
		msg, err := bufio.NewReader(clientConn).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var cm clientMessage
		cm.clientConn = clientConn
		cm.msg = msg
		c <- cm
	}
}

func broadcastMessage(c chan clientMessage) {
	for {
		cm := <-c
		msg := cm.msg
		clientConn := cm.clientConn
		for _, client := range clients {
			if client != clientConn {
				client.Write([]byte(msg))
			}
		}
	}
}
