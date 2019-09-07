package http

import "strings"

type RequestParser struct {
}

type ParsedRequest struct {
	Method string
	Path   string
}

func (p *RequestParser) Parse(requestString string) *ParsedRequest {
	req := &ParsedRequest{}
	if requestString == "" {
		return req
	}

	parts := strings.Split(requestString, " ")

	if len(parts) >= 1 {
		req.Method = parts[0]
	}

	if len(parts) >= 2 {
		req.Path = parts[1]
	}

	return req
}
