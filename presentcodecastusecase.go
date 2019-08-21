package gocleanarch

// PresentCodecastUseCase is a use case that handles the presentation of a codecast.
// It belongs in the use case layer.
type PresentCodecastUseCase struct {
}

func (codecastUseCase *PresentCodecastUseCase) PresentCodecasts(user *User) []*PresentableCodecast {
	return []*PresentableCodecast{}
}

func (codecastUseCase *PresentCodecastUseCase) IsLicensedToViewCodecast(user *User, codecast *Codecast) bool {
	licenses := AGateway.FindLicensesForUserAndCodecast(user, codecast)
	return len(licenses) > 0
}
