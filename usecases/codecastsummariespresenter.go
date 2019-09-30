package usecases

type CodecastSummariesPresenter struct {
	ViewModel *CodecastSummariesViewModel
}

type ViewableCodecastSummary struct {
	Title           string
	Permalink       string
	PublicationDate string
	IsViewable      bool
	IsDownloadable  bool
}

type CodecastSummariesViewModel struct {
	ViewableCodecastSummaries []*ViewableCodecastSummary
}

func (c *CodecastSummariesPresenter) GetViewModel() *CodecastSummariesViewModel {
	return c.ViewModel
}

func (c *CodecastSummariesPresenter) Present(responseModel *CodecastSummariesResponseModel) {
	c.ViewModel = &CodecastSummariesViewModel{}
	for _, summary := range responseModel.CodecastSummaries {
		c.ViewModel.ViewableCodecastSummaries = append(c.ViewModel.ViewableCodecastSummaries, makeViewable(summary))
	}
}

func makeViewable(summary *codecastSummary) *ViewableCodecastSummary {
	return &ViewableCodecastSummary{
		Title:           summary.Title,
		Permalink:       summary.Permalink,
		PublicationDate: summary.PublicationDate.Format("1/02/2006"),
		IsViewable:      summary.IsViewable,
		IsDownloadable:  summary.IsDownloadable,
	}
}
