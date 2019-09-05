package gocleanarch

type CodecastDetailsUseCase struct {
}

func (useCase *CodecastDetailsUseCase) RequestCodecastDetails(user *User, permalink string) *PresentableCodecastDetails {
	details := &PresentableCodecastDetails{}
	var found bool
	if codecast := CodecastRepo.FindByPermalink(permalink); codecast != nil {
		found = true
		new(CodecastSummaryUseCase).FormatSummaryFields(&details.PresentableCodecastSummary, codecast, user)
	}
	details.Found = found
	return details
}
