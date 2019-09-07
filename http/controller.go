package http

type Controller interface {
	Handle(request *ParsedRequest) string
}
