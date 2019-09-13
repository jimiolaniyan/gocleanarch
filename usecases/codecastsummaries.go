package usecases

import (
	. "github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/entities"
)

// CodecastSummariesUseCase is a use case that handles the presentation of a codecast.
// It belongs in the use case layer.
type CodecastSummariesUseCase struct {
}

func (c *CodecastSummariesUseCase) SummarizeCodecasts(*entities.User, CodecastSummariesOutputBoundary) {

}

type CodecastSummariesResponseModel struct {
	Title           string
	PublicationDate string
	Permalink       string
	IsViewable      bool
	IsDownLoadable  bool
}

func (c *CodecastSummariesUseCase) PresentCodecasts(loggedInUser *entities.User) []*CodecastSummariesResponseModel {
	var presentableCodecasts []*CodecastSummariesResponseModel
	for _, codecast := range CodecastRepo.FindAllCodecastsSortedChronologically() {
		presentableCodecasts = append(presentableCodecasts, CodecastSummariesPresenter{}.FormatCodecast(codecast, loggedInUser))
	}
	return presentableCodecasts
}



func (codecastUseCase *CodecastSummariesUseCase) IsLicensedFor(licenseType entities.LicenseType, user *entities.User, codecast *entities.Codecast) bool {
	licenses := LicenseRepo.FindLicensesForUserAndCodecast(user, codecast)
	for _, l := range licenses {
		if l.LicenseType() == licenseType {
			return true
		}
	}
	return false
}

