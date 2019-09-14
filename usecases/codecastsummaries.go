package usecases

import (
	. "github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/entities"
	"time"
)

// CodecastSummariesUseCase is a use case that handles the presentation of a codecast.
// It belongs in the use case layer.
type CodecastSummariesUseCase struct {
}

type codecastSummary struct {
	Title           string
	PublicationDate time.Time
	Permalink       string
	IsViewable      bool
	IsDownloadable  bool
}

type CodecastSummariesResponseModel struct {
	CodecastSummaries []*codecastSummary
}

func (c *CodecastSummariesResponseModel) addCodecastSummary(summary *codecastSummary) {
	c.CodecastSummaries = append(c.CodecastSummaries, summary)
}

func (c *CodecastSummariesUseCase) IsLicensedFor(licenseType entities.LicenseType, user *entities.User, codecast *entities.Codecast) bool {
	licenses := LicenseRepo.FindLicensesForUserAndCodecast(user, codecast)
	for _, l := range licenses {
		if l.LicenseType() == licenseType {
			return true
		}
	}
	return false
}

func (c *CodecastSummariesUseCase) SummarizeCodecasts(loggedInUser *entities.User, presenter CodecastSummariesOutputBoundary) {
	responseModel := &CodecastSummariesResponseModel{}

	for _, codecast := range CodecastRepo.FindAllCodecastsSortedChronologically() {
		responseModel.addCodecastSummary(c.summarizeCodecast(codecast, loggedInUser))
	}

	presenter.Present(responseModel)
}

func (c *CodecastSummariesUseCase) summarizeCodecast(codecast *entities.Codecast, user *entities.User) *codecastSummary {
	return &codecastSummary{
		Title:           codecast.Title(),
		PublicationDate: codecast.PublicationDate(),
		Permalink:       codecast.Permalink(),
		IsViewable:      c.IsLicensedFor(entities.Viewing, user, codecast),
		IsDownloadable:  c.IsLicensedFor(entities.Downloading, user, codecast),
	}
}
