package main

import (
	"fmt"
	. "github.com/jimiolaniyan/gocleanarch/socketserver"
	"github.com/jimiolaniyan/gocleanarch/view"
	"net"
	"path/filepath"
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
	fmt.Println("wrote")
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
	frontPageFilePath, err := filepath.Abs("./web/html/frontpage.html")
	codecastPath, _ := filepath.Abs("./web/html/codecast.html")

	if err != nil {
		fmt.Printf("Could not open %s: %s", "./web/html/frontpage.html", err)
	}

	if frontPageTemplate, err := view.CreateTemplate(frontPageFilePath); err == nil {
		codecastTemplate, _ := view.CreateTemplate(codecastPath)
		codecastTemplate.Replace("title", "Episode 1: The beginning")
		codecastView := codecastTemplate.View

		frontPageTemplate.Replace("codecasts", codecastView)

		return frontPageTemplate.View
	}
	return "Gunk"
}

func main() {
	done := make(chan bool)
	server, _ := NewSocketServer(8081, &MainService{})
	server.Start()
	<-done
}
