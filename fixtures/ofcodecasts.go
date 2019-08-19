package fixtures

type QueryResponse struct {
	Title        string
	Picture      string
	Description  string
	Viewable     bool
	downloadable bool
}

func Query() []QueryResponse {
	return []QueryResponse{}
}
