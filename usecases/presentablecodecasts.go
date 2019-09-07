package usecases

// PresentableCodecastSummary is a data structure that returns
// a code cast in format that is acceptable for a presenter.
// It belongs in the interface adapters layer
type PresentableCodecastSummary struct {
	Title           string
	PublicationDate string
	Permalink       string
	IsViewable      bool
	IsDownLoadable  bool
}

type PresentableCodecastDetails struct {
	PresentableCodecastSummary
	Found bool
}