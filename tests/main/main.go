package main

import (
	"fmt"
	. "github.com/jimiolaniyan/gocleanarch/socketserver"
	"net"
	"strconv"
)

type Main struct {
	Server  *SocketServer
	Service SocketService
}

type MainService struct {
}

func (ms *MainService) Serve(c net.Conn) {
	frontPage := getFrontPage()
	response := makeResponse(frontPage)
	_, err := c.Write([]byte(response))
	if err != nil {
		fmt.Println(err)
	}
}

func makeResponse(content string) (response string) {
	response = "HTTP/1.1 200 OK\n" +
		"Content-Length: " + strconv.Itoa(len(content)) + "\n" +
		"\n" +
		content
	return
}

func getFrontPage() string {
	return "Gunk"
}

func main() {
	done := make(chan bool)
	server, _ := NewSocketServer(8081, &MainService{})
	server.Start()
	<-done
}
