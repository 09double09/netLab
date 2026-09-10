package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	text := make(chan string)
	done := make(chan struct{})

	go inputListener(text)
	go connectListener(text, done)

	<-done
	println("exit")
}

func inputListener(text chan<- string) {
	defer close(text)
	rd := bufio.NewScanner(os.Stdin)
	for rd.Scan() {
		if rd.Text() == "exit" {
			return
		}
		text <- rd.Text()
	}
}

func connectListener(text <-chan string, done chan<- struct{}) {
	defer close(done)
	conn, err := net.Dial("tcp", "localhost:3002")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for input := range text {
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			log.Println("write error:", err)
			return
		}
		fmt.Println("write:" + input)
	}
}