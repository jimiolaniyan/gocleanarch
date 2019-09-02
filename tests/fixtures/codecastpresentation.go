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
	AGateway = NewMockGateway()
	return &codecastPresentation{gateKeeper: new(GateKeeper)}
}

func (c *codecastPresentation) ClearCodecasts() bool {
	var codecasts = AGateway.FindAllCodecastsSortedChronologically()

	// TODO not a perfect solution
	for i := len(codecasts) - 1; i >= 0; i-- {
		AGateway.Delete(codecasts[i])
	}
	return len(AGateway.FindAllCodecastsSortedChronologically()) == 0
}

func (c *codecastPresentation) LoginUser(username string) bool {
	user := AGateway.FindUserByName(username)
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
	AGateway.SaveUser(NewUser(username))
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
	user := AGateway.FindUserByName(username)
	codecast := AGateway.FindCodecastByTitle(codecastTitle)
	var license = NewLicense(Viewing, user, codecast)
	AGateway.SaveLicense(license)
	return c.useCase.IsLicensedFor(Viewing, user, codecast)
}

func (c *codecastPresentation) CreateLicenceForDownloading(username string, codecastTitle string) bool {
	user := AGateway.FindUserByName(username)
	codecast := AGateway.FindCodecastByTitle(codecastTitle)
	var license = NewLicense(Downloading, user, codecast)
	AGateway.SaveLicense(license)
	return c.useCase.IsLicensedFor(Downloading, user, codecast)
}
