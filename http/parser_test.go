package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyRequest(t *testing.T) {
	parser := RequestParser{}
	parsedRequest := parser.Parse("")

	assert.Equal(t, "", parsedRequest.Method)
	assert.Equal(t, "", parsedRequest.Path)
}

func TestNonEmptyRequest(t *testing.T) {
	parser := RequestParser{}
	parsedRequest := parser.Parse("GET /foo/bar HTTP/1.1")

	assert.Equal(t, "GET", parsedRequest.Method)
	assert.Equal(t, "/foo/bar", parsedRequest.Path)
}

func TestPartialRequest(t *testing.T) {
	parser := RequestParser{}
	parsedRequest := parser.Parse("GET")

	assert.Equal(t, "GET", parsedRequest.Method)
	assert.Equal(t, "", parsedRequest.Path)
}