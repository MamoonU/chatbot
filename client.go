package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func client() error {

	conn, err := net.Dial("tcp", "35.179.140.138:8080")
	if err != nil {
		return err
	}

	go messageReader(conn)

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

func messageReader(conn net.Conn) {

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("\r%s\n", scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println("error noob: ", scanner.Err())
	}

}
