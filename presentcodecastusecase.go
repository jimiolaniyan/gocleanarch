package gocleanarch

// PresentCodecastUseCase is a use case that handles the presentation of a codecast.
// It belongs in the use case layer.
type PresentCodecastUseCase struct {
}

func (codecastUseCase *PresentCodecastUseCase) PresentCodecasts(loggedInUser *User) []*PresentableCodecast {
	var presentableCodecasts []*PresentableCodecast
	for _, codecast := range AGateway.FindAllCodecastsSortedChronologically() {
		presentableCodecasts = append(presentableCodecasts, codecastUseCase.formatCodecast(codecast, loggedInUser))
	}
	return presentableCodecasts
}

func (codecastUseCase *PresentCodecastUseCase) formatCodecast(codecast *Codecast, user *User) *PresentableCodecast {
	pc := &PresentableCodecast{}
	pc.Title = codecast.Title()
	pc.PublicationDate = codecast.PublicationDate().Format("1/2/2006")
	pc.IsViewable = codecastUseCase.IsLicensedFor(Viewing, user, codecast)
	pc.IsDownLoadable = codecastUseCase.IsLicensedFor(Downloading, user, codecast)
	return pc
}

func (codecastUseCase *PresentCodecastUseCase) IsLicensedFor(licenseType LicenseType, user *User, codecast *Codecast) bool {
	licenses := AGateway.FindLicensesForUserAndCodecast(user, codecast)
	for _, l := range licenses {
		if l.LicenseType() == licenseType {
			return true
		}
	}
	return false
}
