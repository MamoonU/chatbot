package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	mode     string
	nickname string
)

func main() {

	flag.StringVar(&mode, "mode", "", "Select Mode")
	flag.StringVar(&nickname, "nickname", "", "Select Nickname")
	flag.Parse()

	if nickname == "" && mode == "client" {
		fmt.Println("Error: Nickname Required")
		return
	}

	if mode == "" {
		fmt.Println("Error 404")
		return
	}
	if mode == "server" {
		fmt.Println("Server Mode Enabled")
		err := server()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if mode == "client" {
		fmt.Println("Client Mode Enabled")
		err := client()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("Error")
		return
	}

}

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

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}

func client() error {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintf(conn, "[%s] %s\n", nickname, scanner.Text())
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
	//fmt.Fprintf(conn, "Client Connection Test\n")
}
