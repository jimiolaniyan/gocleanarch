package main

import (
	"fmt"
	"net"
)

type HelloService struct {
}

func (hs *HelloService) Serve(c net.Conn) {
	response := "HTTP/1.1 200 OK\n" +
		"Content-Length: 21\n" +
		"\n" +
		"<h1>Hello, world</h1>"
	_, err := c.Write([]byte(response))
	if err != nil {
		fmt.Println(err)
	}
}

//func main() {
//	done := make(chan bool)
//	server, err := socketserver.NewSocketServer(8081, &HelloService{})
//	if err != nil {
//		fmt.Println(err)
//	}
//	server.Start()
//	<-done
//}
