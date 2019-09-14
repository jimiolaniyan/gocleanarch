package usecases

type CodecastSummariesPresenter struct {
	ViewModel *CodecastSummariesViewModel
}

type ViewableCodecastSummary struct {
	Title string
	Permalink string
	PublicationDate string
	IsViewable bool
	IsDownloadable bool
}

type CodecastSummariesViewModel struct {
	ViewableCodecastSummaries []*ViewableCodecastSummary
}

//func (c CodecastSummariesPresenter) FormatSummaryFields(pc *CodecastSummariesResponseModel, codecast *entities.Codecast, user *entities.User) {
//	pc.Title = codecast.Title()
//	pc.PublicationDate = codecast.PublicationDate().Format("1/02/2006")
//	pc.IsViewable = new(CodecastSummariesUseCase).IsLicensedFor(entities.Viewing, user, codecast)
//	pc.IsDownloadable =  new(CodecastSummariesUseCase).IsLicensedFor(entities.Downloading, user, codecast)
//	pc.Permalink = codecast.Permalink()
//}

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
