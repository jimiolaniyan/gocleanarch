package usecases

import (
	. "github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/entities"
)

// CodecastSummariesUseCase is a use case that handles the presentation of a codecast.
// It belongs in the use case layer.
type CodecastSummariesUseCase struct {
}

type CodecastSummaryResponseModel struct {
}

func (codecastUseCase *CodecastSummariesUseCase) PresentCodecasts(loggedInUser *entities.User) []*PresentableCodecastSummary {
	var presentableCodecasts []*PresentableCodecastSummary
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
