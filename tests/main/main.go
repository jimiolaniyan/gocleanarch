package main

import (
	"bufio"
	"fmt"
	"github.com/jimiolaniyan/gocleanarch/http"
	. "github.com/jimiolaniyan/gocleanarch/socketserver"
	"github.com/jimiolaniyan/gocleanarch/tests/setup"
	. "github.com/jimiolaniyan/gocleanarch/usecases"
	"net"
	"strconv"
)

var router *http.Router

func main() {
	setup.LoadSampleData()

	router = http.NewRouter()
	useCase := &CodecastSummariesUseCase{}
	presenter := &CodecastSummariesPresenter{}
	view := &CodecastSummariesViewImpl{}
	controller := &CodecastSummariesController{UseCase: useCase, Presenter: presenter, View: view}
	router.AddPath("", controller)
	//router.AddPath("/episode", CodecastDetailsController{})

	done := make(chan bool)
	server, _ := NewSocketServer(8081, &MainService{})
	server.Start()
	<-done
}

type MainService struct {
}

func (ms *MainService) Serve(c net.Conn) {
	defer c.Close()
	r := bufio.NewReader(c)
	line, err := r.ReadString(byte('\n'))
	checkError(err, "")

	request := new(http.RequestParser).Parse(line)
	response := router.Route(request)
	_, err = c.Write([]byte(makeResponse(response)))

	checkError(err, "")
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s", message, err)
	}
}

func makeResponse(content string) string {
	return "HTTP/1.1 200 OK\n" +
		"Content-Length: " + strconv.Itoa(len(content)) + "\n" +
		"\n" +
		content
}