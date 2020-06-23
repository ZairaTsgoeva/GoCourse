package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	var events = make(chan bool)

	go handleUserInput(events)
	go handleCommands(events)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleUserInput(exit chan<- bool) {
	for {
		var command string
		fmt.Scanln(&command)
		switch command {
		case "exit":
			exit <- true
		}
	}
}

func handleCommands(events <-chan bool) {
	for {
		select {
		case <-events:
			println("os exit")
			os.Exit(0)
		}
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	for {
		// тут select потому что не поняла, надо ли явно завершать эти горутины, думала еще и их уведомить о завершении, но из задания непонятно
		select {
		default:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}
}
