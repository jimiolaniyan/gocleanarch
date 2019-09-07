package usecases

import (
	"github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/entities"
)


type CodecastDetailsUseCase struct {
}

func (useCase *CodecastDetailsUseCase) RequestCodecastDetails(user *entities.User, permalink string) *PresentableCodecastDetails {
	details := &PresentableCodecastDetails{}
	var found bool
	if codecast := gocleanarch.CodecastRepo.FindByPermalink(permalink); codecast != nil {
		found = true
		new(CodecastSummariesUseCase).FormatSummaryFields(&details.PresentableCodecastSummary, codecast, user)
	}
	details.Found = found
	return details
}
