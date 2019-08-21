package gocleanarch

// PresentCodecastUseCase is a use case that handles the presentation of a codecast.
// It belongs in the use case layer.
type PresentCodecastUseCase struct {
}

func (codecastUseCase *PresentCodecastUseCase) PresentCodecasts(loggedInUser *User) []*PresentableCodecast {
	//allCodecasts := AGateway.FindAllCodecasts()
	var presentableCodecasts []*PresentableCodecast

	for _, codecast := range AGateway.FindAllCodecasts() {
		pc := &PresentableCodecast{}
		pc.Title = codecast.Title()
		pc.PublicationDate = codecast.PublicationDate()
		pc.IsViewable = codecastUseCase.IsLicensedToViewCodecast(loggedInUser, codecast)
		presentableCodecasts = append(presentableCodecasts, pc)
	}

	return presentableCodecasts
}

func (codecastUseCase *PresentCodecastUseCase) IsLicensedToViewCodecast(user *User, codecast *Codecast) bool {
	licenses := AGateway.FindLicensesForUserAndCodecast(user, codecast)
	return len(licenses) > 0
}
