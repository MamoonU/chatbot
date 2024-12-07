package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

func server() error {

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {

	addClient(conn)
	defer removeClient(conn)

	fmt.Println("Client connecting: ", conn.RemoteAddr().String())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		broadcast(scanner.Text(), conn)
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}

var (
	clients []net.Conn = make([]net.Conn, 0)
	lock               = &sync.Mutex{}
)

func addClient(conn net.Conn) {

	lock.Lock()
	defer lock.Unlock()

	clients = append(clients, conn)

}

func removeClient(conn net.Conn) {

	lock.Lock()
	defer lock.Unlock()

	for i, client := range clients {

		if conn.RemoteAddr().String() == client.RemoteAddr().String() {

			clients = append(clients[:i], clients[i+1:]...)

		}
	}
}

func broadcast(message string, sender net.Conn) {

	lock.Lock()
	defer lock.Unlock()

	for _, client := range clients {

		if sender.RemoteAddr().String() != client.RemoteAddr().String() {

			fmt.Fprintf(client, "%s\n", message)

		}
	}

}
