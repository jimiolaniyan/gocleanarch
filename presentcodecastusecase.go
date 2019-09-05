package gocleanarch

// CodecastSummaryUseCase is a use case that handles the presentation of a codecast.
// It belongs in the use case layer.
type CodecastSummaryUseCase struct {
}

func (codecastUseCase *CodecastSummaryUseCase) PresentCodecasts(loggedInUser *User) []*PresentableCodecastSummary {
	var presentableCodecasts []*PresentableCodecastSummary
	for _, codecast := range CodecastRepo.FindAllCodecastsSortedChronologically() {
		presentableCodecasts = append(presentableCodecasts, codecastUseCase.formatCodecast(codecast, loggedInUser))
	}
	return presentableCodecasts
}

func (codecastUseCase *CodecastSummaryUseCase) formatCodecast(codecast *Codecast, user *User) *PresentableCodecastSummary {
	pc := &PresentableCodecastSummary{}
	codecastUseCase.FormatSummaryFields(pc, codecast, user)
	return pc
}

func (codecastUseCase *CodecastSummaryUseCase) FormatSummaryFields(pc *PresentableCodecastSummary, codecast *Codecast, user *User) {
	pc.Title = codecast.Title()
	pc.PublicationDate = codecast.PublicationDate().Format("1/02/2006")
	pc.IsViewable = codecastUseCase.IsLicensedFor(Viewing, user, codecast)
	pc.IsDownLoadable = codecastUseCase.IsLicensedFor(Downloading, user, codecast)
}

func (codecastUseCase *CodecastSummaryUseCase) IsLicensedFor(licenseType LicenseType, user *User, codecast *Codecast) bool {
	licenses := LicenseRepo.FindLicensesForUserAndCodecast(user, codecast)
	for _, l := range licenses {
		if l.LicenseType() == licenseType {
			return true
		}
	}
	return false
}
