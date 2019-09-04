package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
)

var CodecastPresentation = NewCodecastPresentation()

type codecastPresentation struct {
	gateKeeper *GateKeeper
	useCase    PresentCodecastUseCase
}

func NewCodecastPresentation() *codecastPresentation {
	SetupContext()
	return &codecastPresentation{gateKeeper: new(GateKeeper)}
}

func (c *codecastPresentation) ClearCodecasts() bool {
	var codecasts = CodecastRepo.FindAllCodecastsSortedChronologically()

	// TODO not a perfect solution
	for i := len(codecasts) - 1; i >= 0; i-- {
		CodecastRepo.Delete(codecasts[i])
	}
	return len(CodecastRepo.FindAllCodecastsSortedChronologically()) == 0
}

func (c *codecastPresentation) LoginUser(username string) bool {
	user := UserRepo.FindByName(username)
	if user != nil {
		c.gateKeeper.SetLoggedInUser(user)
		return true
	} else {
		return false
	}
}

func (c *codecastPresentation) LogOutUser() {
	c.gateKeeper.SetLoggedInUser(nil)
}

func (c *codecastPresentation) AddUser(username string) bool {
	UserRepo.Save(NewUser(username))
	return true
}

func (c *codecastPresentation) PresentationUser() string {
	return c.gateKeeper.LoggedInUser().Username()
}

func (c *codecastPresentation) CountOfCodecastsPresented() int {
	presentations := c.useCase.PresentCodecasts(c.gateKeeper.LoggedInUser())
	return len(presentations)
}

func (c *codecastPresentation) CreateLicenceForViewing(username string, codecastTitle string) bool {
	user := UserRepo.FindByName(username)
	codecast := CodecastRepo.FindByTitle(codecastTitle)
	var license = NewLicense(Viewing, user, codecast)
	LicenseRepo.Save(license)
	return c.useCase.IsLicensedFor(Viewing, user, codecast)
}

func (c *codecastPresentation) CreateLicenceForDownloading(username string, codecastTitle string) bool {
	user := UserRepo.FindByName(username)
	codecast := CodecastRepo.FindByTitle(codecastTitle)
	var license = NewLicense(Downloading, user, codecast)
	LicenseRepo.Save(license)
	return c.useCase.IsLicensedFor(Downloading, user, codecast)
}
