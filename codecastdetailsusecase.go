package gocleanarch

type CodecastDetailsUseCase struct {
}

func (useCase *CodecastDetailsUseCase) RequestCodecastDetails(user *User, permalink string) *PresentableCodecastDetails {
	details := &PresentableCodecastDetails{}

	codecast := CodecastRepo.FindByPermalink(permalink)
	details.Title = codecast.title
	details.PublicationDate = codecast.PublicationDate().Format("1/02/2006")
	return details
}
