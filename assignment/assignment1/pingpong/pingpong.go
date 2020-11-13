//pong ping using channels in golang
package main

import (
	"fmt"
	"os"
	"time"
)

func pongerfunc(pingerfunc chan<- int, pongerfunc <-chan int) {
	for {
		val := <-pongerfunc
		fmt.Println("pong")
		time.Sleep(100 * time.Millisecond)
		pingerfunc <- val + 1
	}
}
func pingerfunc(pingerfunc <-chan int, pongerfunc chan<- int) {
	for {
		<-pingerfunc
		fmt.Println("ping")
		time.Sleep(300 * time.Millisecond)
		pongerfunc <- 1
	}
}

func main() {

	pong := make(chan int)
	ping := make(chan int)
	go pingerfunc(ping, pong)
	go pongerfunc(ping, pong)
	//goroutine starts the pong-ping by sending into the pong channel

	pong <- 1
	time.Sleep(5 * time.Second)
	os.Exit(0)
}
