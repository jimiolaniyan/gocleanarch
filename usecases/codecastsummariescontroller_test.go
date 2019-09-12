package usecases

import (
	. "github.com/jimiolaniyan/gocleanarch/http"
	"github.com/jimiolaniyan/gocleanarch/tests/setup"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CodecastSummaryInputBoundarySpy struct {
	summarizeCodecastsWasCalled bool
}

func (c *CodecastSummaryInputBoundarySpy) SummarizeCodecasts() {
	c.summarizeCodecastsWasCalled = true
}

func TestFrontPage(t *testing.T) {
	setup.LoadContext()
	inputBoundarySpy := &CodecastSummaryInputBoundarySpy{}
	request := &ParsedRequest{Method: "GET", Path: "bla"}
	controller := CodecastSummariesController{InputBoundary: inputBoundarySpy}

	controller.Handle(request)

	assert.True(t, inputBoundarySpy.summarizeCodecastsWasCalled)
}
