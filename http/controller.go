package http

import "strconv"

type Controller interface {
	Handle(request *ParsedRequest) string
	MakeResponse(content string) string
}

type DefaultController struct {
	Controller
}

func (c *DefaultController) MakeResponse(content string) string {
	return "HTTP/1.1 200 OK\n" +
		"Content-Length: " + strconv.Itoa(len(content)) + "\n" +
		"\n" +
		content
}
