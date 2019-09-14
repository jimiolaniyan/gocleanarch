package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/entities"
	"github.com/jimiolaniyan/gocleanarch/tests/setup"
	"github.com/jimiolaniyan/gocleanarch/usecases"
)

var CodecastPresentation = NewCodecastPresentation()

type codecastPresentation struct {
	useCase usecases.CodecastSummariesUseCase
}

func NewCodecastPresentation() *codecastPresentation {
	setup.LoadContext()
	return &codecastPresentation{}
}

func (c *codecastPresentation) ClearCodecasts() bool {
	var codecasts = CodecastRepo.FindAllCodecastsSortedChronologically()

	for i := len(codecasts) - 1; i >= 0; i-- {
		CodecastRepo.Delete(codecasts[i])
	}
	return len(CodecastRepo.FindAllCodecastsSortedChronologically()) == 0
}

func (c *codecastPresentation) LoginUser(username string) bool {
	user := UserRepo.FindByName(username)
	if user != nil {
		SessionKeeper.SetLoggedInUser(user)
		return true
	}
	return false
}

func (c *codecastPresentation) LogOutUser() {
	SessionKeeper.SetLoggedInUser(nil)
}

func (c *codecastPresentation) AddUser(username string) bool {
	UserRepo.Save(entities.NewUser(username))
	return true
}

func (c *codecastPresentation) PresentationUser() string {
	return SessionKeeper.LoggedInUser().Username()
}

func (c *codecastPresentation) CountOfCodecastsPresented() int {
	return len(LoadViewableCodecasts())
}

func (c *codecastPresentation) CreateLicenseForType(username string, codecastTitle string, licenseType entities.LicenseType) bool {
	user := UserRepo.FindByName(username)
	codecast := CodecastRepo.FindByTitle(codecastTitle)
	var license = entities.NewLicense(licenseType, user, codecast)
	LicenseRepo.Save(license)
	return c.useCase.IsLicensedFor(licenseType, user, codecast)
}

func (c *codecastPresentation) CreateLicenceForViewing(username string, codecastTitle string) bool {
	return c.CreateLicenseForType(username, codecastTitle, entities.Viewing)
}

func (c *codecastPresentation) CreateLicenceForDownloading(username string, codecastTitle string) bool {
	return c.CreateLicenseForType(username, codecastTitle, entities.Downloading)
}

func LoadViewableCodecasts() []*usecases.ViewableCodecastSummary {
	loggedInUser := SessionKeeper.LoggedInUser()
	useCase := new(usecases.CodecastSummariesUseCase)
	presenter := &usecases.CodecastSummariesPresenter{}
	useCase.SummarizeCodecasts(loggedInUser, presenter)
	viewableCodecastSummaries := presenter.ViewModel.ViewableCodecastSummaries
	return viewableCodecastSummaries
}
