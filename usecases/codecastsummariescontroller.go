package usecases

import (
	"fmt"
	"github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/http"
)

type CodecastSummariesController struct {
	http.DefaultController
}

func (c *CodecastSummariesController) Handle(request *http.ParsedRequest) string {
	useCase := CodecastSummariesUseCase{}
	jimi := gocleanarch.UserRepo.FindByName("jimi")
	presentableCodecasts := useCase.PresentCodecasts(jimi)
	html := CodecastSummariesView{}.toHTML(presentableCodecasts)
	return c.MakeResponse(html)
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s", message, err)
	}
}
