package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/usecases"
)

type QueryResponse struct {
	Title        string
	Picture      string
	Description  string
	Viewable     bool
	Downloadable bool
}

func Query() []QueryResponse {
	loggedInUser := SessionKeeper.LoggedInUser()
	useCase := new(usecases.CodecastSummaryUseCase)
	presentableCodecasts := useCase.PresentCodecasts(loggedInUser)
	var response []QueryResponse

	for _, pc := range presentableCodecasts {
		response = append(response, QueryResponse{
			Title:        pc.Title,
			Picture:      pc.Title,
			Description:  pc.Title,
			Viewable:     pc.IsViewable,
			Downloadable: pc.IsDownLoadable,
		})
	}

	return response
}
