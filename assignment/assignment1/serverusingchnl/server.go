package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

func main() {
	ln, err := net.Listen("tcp", ":5050")
	check(err, "Welcome to server :).")

	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")

		go func() {
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("See you soon.\n")
					break
				}

				conn.Write([]byte("Hi, " + name))
			}
		}()
	}
}