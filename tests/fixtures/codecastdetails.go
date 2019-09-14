package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/usecases"
)

type CodecastDetails struct {
	useCase *usecases.CodecastDetailsUseCase
	details *usecases.PresentableCodecastDetails
}

func (cd *CodecastDetails) RequestCodecast(permalink string) bool {
	cd.details = cd.useCase.RequestCodecastDetails(SessionKeeper.LoggedInUser(), permalink)
	return cd.details != nil
}

func (cd *CodecastDetails) CodecastDetailsOfferPurchaseOf(licenseType string) bool {
	//return (strings.EqualFold(licenseType, "viewing") && !cd.details.IsViewable) ||
	//	(strings.EqualFold(licenseType, "download") && !cd.details.IsDownLoadable)
	return false
}

func (cd *CodecastDetails) CodecastDetailsTitle() string {
	//return cd.details.Title
	return ""
}

func (cd *CodecastDetails) CodecastDetailsDate() string {
	//return cd.details.PublicationDate
	return ""
}
