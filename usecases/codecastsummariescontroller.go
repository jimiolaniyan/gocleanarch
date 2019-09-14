package usecases

import (
	"fmt"
	"github.com/jimiolaniyan/gocleanarch"
	. "github.com/jimiolaniyan/gocleanarch/http"
)

type CodecastSummariesController struct {
	UseCase   CodecastSummariesInputBoundary
	Presenter CodecastSummariesOutputBoundary
	View      CodecastSummariesView
}

func (c *CodecastSummariesController) Handle(request *ParsedRequest) string {
	user := gocleanarch.SessionKeeper.LoggedInUser()
	c.UseCase.SummarizeCodecasts(user, c.Presenter)
	return c.View.Generate(c.Presenter.GetViewModel())
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s", message, err)
	}
}
