package gocleanarch

type CodecastDetailsUseCase struct {
}

func (useCase *CodecastDetailsUseCase) RequestCodecastDetails(user *User, permalink string) *PresentableCodecastDetails {
	details := &PresentableCodecastDetails{}

	codecast := CodecastRepo.FindByPermalink(permalink)
	new(PresentCodecastUseCase).DoFormatCodecast(&details.PresentableCodecast, codecast, user)
	return details
}
