package usecases

import (
	"fmt"
	"github.com/jimiolaniyan/gocleanarch"
	. "github.com/jimiolaniyan/gocleanarch/http"
)

type CodecastSummariesController struct {
	DefaultController
	UseCase   CodecastSummariesInputBoundary
	Presenter CodecastSummariesOutputBoundary
	View      CodecastSummariesView
}

func (c *CodecastSummariesController) Handle(request *ParsedRequest) string {
	user := gocleanarch.SessionKeeper.LoggedInUser()
	c.UseCase.SummarizeCodecasts(user, c.Presenter)
	c.View.Generate(c.Presenter.GetResponseModel())
	return ""
	//useCase := CodecastSummariesUseCase{}
	//jimi := gocleanarch.UserRepo.FindByName("jimi")
	//presentableCodecasts := useCase.PresentCodecasts(jimi)
	//html := CodecastSummariesViewImpl{}.toHTML(presentableCodecasts)
	//return c.MakeResponse(html)
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s", message, err)
	}
}
