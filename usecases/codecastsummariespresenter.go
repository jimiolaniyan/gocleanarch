package usecases

import "github.com/jimiolaniyan/gocleanarch/entities"

type CodecastSummariesPresenter struct {

}

func (c CodecastSummariesPresenter) FormatSummaryFields(pc *PresentableCodecastSummary, codecast *entities.Codecast, user *entities.User) {
	pc.Title = codecast.Title()
	pc.PublicationDate = codecast.PublicationDate().Format("1/02/2006")
	pc.IsViewable = new(CodecastSummariesUseCase).IsLicensedFor(entities.Viewing, user, codecast)
	pc.IsDownLoadable =  new(CodecastSummariesUseCase).IsLicensedFor(entities.Downloading, user, codecast)
	pc.Permalink = codecast.Permalink()
}

func (c CodecastSummariesPresenter) FormatCodecast(codecast *entities.Codecast, user *entities.User) *PresentableCodecastSummary {
	pc := &PresentableCodecastSummary{}
	CodecastSummariesPresenter{}.FormatSummaryFields(pc, codecast, user)
	return pc
}