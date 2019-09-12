package usecases

import (
	"fmt"
	. "github.com/jimiolaniyan/gocleanarch/http"
)

type CodecastSummariesController struct {
	DefaultController
	InputBoundary CodecastSummaryInputBoundary
}

func (c *CodecastSummariesController) Handle(request *ParsedRequest) string {
	c.InputBoundary.SummarizeCodecasts()
	return ""
	//useCase := CodecastSummariesUseCase{}
	//jimi := gocleanarch.UserRepo.FindByName("jimi")
	//presentableCodecasts := useCase.PresentCodecasts(jimi)
	//html := CodecastSummariesView{}.toHTML(presentableCodecasts)
	//return c.MakeResponse(html)
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s", message, err)
	}
}
