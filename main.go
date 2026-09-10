package main

import (
	"bufio"
	"fmt"
	"net"
)


func main(){
	ls, err := net.Listen("tcp","localhost:3002")
	if err != nil {
		panic(err)
	}
	fmt.Println("tcp listen at port 80")
	defer ls.Close()
	c := make(chan bool)
	listen(ls,c)
	<-c
}

func listen(listener net.Listener, c chan<- bool ){
	
	for {
		conn, err:=listener.Accept()
		println("conneting")
		if err != nil {
			close(c)
			panic(err)
		}
		
		go func (c net.Conn)  {
			scanner := bufio.NewScanner(c)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		
			if err := scanner.Err(); err != nil {
				fmt.Println("read error:", err)
			}
		}(conn)
	}
}