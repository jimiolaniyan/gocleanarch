package gocleanarch

import (
	"reflect"
)

// PresentCodecastUseCase is a use case that handles the presentation of a codecast.
// It belongs in the use case layer.
type PresentCodecastUseCase struct {
}

func (codecastUseCase *PresentCodecastUseCase) PresentCodecasts(loggedInUser *User) []*PresentableCodecast {
	var presentableCodecasts []*PresentableCodecast

	for _, codecast := range AGateway.FindAllCodecastsSortedChronologically() {
		pc := &PresentableCodecast{}
		pc.Title = codecast.Title()
		pc.PublicationDate = codecast.PublicationDate().Format("1/2/2006")
		pc.IsViewable = codecastUseCase.IsLicensedToViewCodecast(loggedInUser, codecast)
		pc.IsDownLoadable = codecastUseCase.IsLicensedToDownloadCodecast(loggedInUser, codecast)
		presentableCodecasts = append(presentableCodecasts, pc)
	}

	return presentableCodecasts
}

func (codecastUseCase *PresentCodecastUseCase) IsLicensedToViewCodecast(user *User, codecast *Codecast) bool {
	licenses := AGateway.FindLicensesForUserAndCodecast(user, codecast)
	for _, l := range licenses {
		if reflect.TypeOf(l) == reflect.TypeOf(&ViewableLicense{}) {
			return true
		}
	}
	return false
}

func (codecastUseCase *PresentCodecastUseCase) IsLicensedToDownloadCodecast(user *User, codecast *Codecast) bool {
	licenses := AGateway.FindLicensesForUserAndCodecast(user, codecast)
	for _, l := range licenses {
		if reflect.TypeOf(l) == reflect.TypeOf(&DownloadLicense{}) {
			return true
		}
	}
	return false
}
