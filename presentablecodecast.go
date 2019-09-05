package gocleanarch

// PresentableCodecastSummary is a data structure that returns
// a code cast in format that is acceptable for a presenter.
// It belongs in the interface adapters layer
type PresentableCodecastSummary struct {
	IsViewable      bool
	Title           string
	PublicationDate string
	IsDownLoadable  bool
}
