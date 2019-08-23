package fixtures

import . "github.com/jimiolaniyan/gocleanarch"

type QueryResponse struct {
	Title        string
	Picture      string
	Description  string
	Viewable     bool
	Downloadable bool
}

func Query() []QueryResponse {
	loggedInUser := CodecastPresentation.gateKeeper.LoggedInUser()
	useCase := new(PresentCodecastUseCase)
	presentableCodecasts := useCase.PresentCodecasts(loggedInUser)
	var response []QueryResponse

	for _, pc := range presentableCodecasts {
		response = append(response, QueryResponse{
			Title:        pc.Title,
			Picture:      pc.Title,
			Description:  pc.Title,
			Viewable:     pc.IsViewable,
			Downloadable: false,
		})
	}

	return response
}
