package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
	"strings"
)

type CodecastDetails struct {
	useCase *CodecastDetailsUseCase
	details *PresentableCodecastDetails
}

func (cd *CodecastDetails) RequestCodecast(permalink string) bool {
	cd.details = cd.useCase.RequestCodecastDetails(SessionKeeper.LoggedInUser(), permalink)
	return cd.details != nil
}

func (cd *CodecastDetails) CodecastDetailsOfferPurchaseOf(licenseType string) bool {
	return (strings.EqualFold(licenseType, "viewing") && !cd.details.IsViewable) ||
		(strings.EqualFold(licenseType, "download") && !cd.details.IsDownLoadable)
}

func (cd *CodecastDetails) CodecastDetailsTitle() string {
	return cd.details.Title
}

func (cd *CodecastDetails) CodecastDetailsDate() string {
	return cd.details.PublicationDate
}
