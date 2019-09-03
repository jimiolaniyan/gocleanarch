package main

import (
	"fmt"
	"github.com/jimiolaniyan/gocleanarch"
	. "github.com/jimiolaniyan/gocleanarch/socketserver"
	"github.com/jimiolaniyan/gocleanarch/tests"
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
	useCase := gocleanarch.PresentCodecastUseCase{}
	jimi := gocleanarch.AGateway.FindUserByName("jimi")
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

			//staged
			codecastTemplate.Replace("thumbnail", "https://cleancoders.com/images/portraits/robert-martin.jpg")
			codecastTemplate.Replace("author", "Jimi")
			codecastTemplate.Replace("duration", "58 mins.")
			codecastTemplate.Replace("contentActions", "Buying options go here.")
			codecastLines.WriteString(codecastTemplate.View)
		}

		fmt.Println(codecastLines.String())

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

func main() {
	tests.SetupSampleData()
	done := make(chan bool)
	server, _ := NewSocketServer(8081, &MainService{})
	server.Start()
	<-done
}
