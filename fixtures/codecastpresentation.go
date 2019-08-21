package fixtures

import (
	"fmt"
	. "github.com/jimiolaniyan/gocleanarch"
)

type codecastPresentation struct {
	gateKeeper GateKeeper
	useCase PresentCodecastUseCase
}

func NewCodecastPresentation() *codecastPresentation {
	AGateway = NewMockGateway()
	return &codecastPresentation{}
}

func (c *codecastPresentation) ClearCodecasts() bool {
	var codecasts = AGateway.FindAllCodecasts()

	// TODO not a perfect solution
	for i := len(codecasts)-1; i >= 0; i-- {
		AGateway.Delete(codecasts[i])
	}
	return len(AGateway.FindAllCodecasts()) == 0
}

func (c *codecastPresentation) LoginUser(username string) bool {
	user := AGateway.FindUser(username)
	if user != nil {
		c.gateKeeper.SetLoggedInUser(user)
		return true
	} else {
		return false
	}
}

func (c *codecastPresentation) AddUser(username string) bool {
	AGateway.SaveUser(&User{Username:username})
	return true
}

func (c *codecastPresentation) PresentationUser() string {
	return c.gateKeeper.LoggedInUser().Username
}

func (c *codecastPresentation) CountOfCodecastsPresented() int {
	presentations := c.useCase.PresentCodecasts(c.gateKeeper.LoggedInUser())
	return len(presentations)
}

func (c *codecastPresentation) CreateLicenceForViewing(username string, codecastTitle string) bool {
	user := AGateway.FindUser(username)
	codecast := AGateway.FindCodecastByTitle(codecastTitle)
	var license = &License{User:user, Codecast:codecast}
	AGateway.SaveLicense(license)
	fmt.Println(AGateway)
	return c.useCase.IsLicensedToViewCodecast(user, codecast)
}
