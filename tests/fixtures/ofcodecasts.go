package fixtures

type QueryResponse struct {
	Title        string
	Picture      string
	Description  string
	Viewable     bool
	Downloadable bool
}

func Query() []QueryResponse {
	var response []QueryResponse
	viewableCodecastSummaries := LoadViewableCodecasts()
	for _, summary := range viewableCodecastSummaries {
		response = append(response, QueryResponse{
			Title:        summary.Title,
			Picture:      summary.Title,
			Description:  summary.Title,
			Viewable:     summary.IsViewable,
			Downloadable: summary.IsDownloadable,
		})
	}

	return response
}
