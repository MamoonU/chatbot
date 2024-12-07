package main

import (
	"flag"
	"fmt"
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
