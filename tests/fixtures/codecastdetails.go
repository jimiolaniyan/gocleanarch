package fixtures

import (
	"fmt"
	. "github.com/jimiolaniyan/gocleanarch"
)

type CodecastDetails struct {
	useCase CodecastDetailsUseCase
	details *PresentableCodecastDetails
}

func (cd *CodecastDetails) RequestCodecast(permalink string) bool {
	cd.details = cd.useCase.RequestCodecastDetails(SessionKeeper.LoggedInUser(), permalink)
	return cd.details != nil
}

func (cd *CodecastDetails) CodecastDetailsOfferPurchaseOf(licenseType string) bool {
	return false
}

func (cd *CodecastDetails) CodecastDetailsTitle() string {
	return cd.details.Title
}

func (cd *CodecastDetails) CodecastDetailsDate() string {
	fmt.Println(cd.details.PublicationDate)
	return cd.details.PublicationDate
}
