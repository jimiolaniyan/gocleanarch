package main

import (
	"bufio"
	"fmt"
	"github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/http"
	. "github.com/jimiolaniyan/gocleanarch/socketserver"
	"github.com/jimiolaniyan/gocleanarch/tests/setup"
	"github.com/jimiolaniyan/gocleanarch/usecases"
	"github.com/jimiolaniyan/gocleanarch/view"
	"net"
	"path/filepath"
	"strconv"
	"strings"
)

type Main struct {
	Server  *SocketServer
	Service SocketService
}

type MainService struct {
}

func (ms *MainService) Serve(c net.Conn) {
	defer c.Close()
	r := bufio.NewReader(c)
	line, err := r.ReadString(byte('\n'))
	checkError(err, "")

	request := new(http.RequestParser).Parse(line)
	if response := router.Route(request); response != "" {
		_, err = c.Write([]byte(response))
	}

	_ , err = c.Write([]byte("HTTP/1.1 404 Not Found\n"))

	checkError(err, "")
}

type CodecastSummariesController struct {
}

func (c *CodecastSummariesController) Handle(request *http.ParsedRequest) string {
	frontPage := getFrontPage()
	return makeResponse(frontPage)
}

var router *http.Router

func main() {
	setup.SetupSampleData()

	router = http.NewRouter()
	router.AddPath("", &CodecastSummariesController{})
	//router.AddPath("/episode", CodecastDetailsController{})

	done := make(chan bool)
	server, _ := NewSocketServer(8081, &MainService{})
	server.Start()
	<-done
}

func makeResponse(content string) (response string) {
	response = "HTTP/1.1 200 OK\n" +
		"Content-Length: " + strconv.Itoa(len(content)) + "\n" +
		"\n" +
		content
	return
}

func getFrontPage() string {
	useCase := usecases.CodecastSummariesUseCase{}
	jimi := gocleanarch.UserRepo.FindByName("jimi")
	presentableCodecasts := useCase.PresentCodecasts(jimi)

	frontPageFilePath, err := filepath.Abs("./web/html/frontpage.html")
	checkError(err, fmt.Sprintf("Could not open %s", "./web/html/frontpage.html"))

	codecastPath, err := filepath.Abs("./web/html/codecast.html")
	checkError(err, fmt.Sprintf("Could not open %s", "./web/html/codecast.html"))

	if frontPageTemplate, err := view.CreateTemplate(frontPageFilePath); err == nil {
		var codecastLines strings.Builder

		for _, pc := range presentableCodecasts {
			codecastTemplate, _ := view.CreateTemplate(codecastPath)
			codecastTemplate.Replace("title", pc.Title)
			codecastTemplate.Replace("publicationDate", pc.PublicationDate)
			codecastTemplate.Replace("permalink", pc.Permalink)

			//staged
			codecastTemplate.Replace("thumbnail", "https://cleancoders.com/images/portraits/robert-martin.jpg")
			codecastTemplate.Replace("author", "Jimi")
			codecastTemplate.Replace("duration", "58 mins.")
			codecastTemplate.Replace("contentActions", "Buying options go here.")
			codecastLines.WriteString(codecastTemplate.View)
		}

		frontPageTemplate.Replace("codecasts", codecastLines.String())

		return frontPageTemplate.View
	}
	return "Gunk"
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s", message, err)
	}
}
