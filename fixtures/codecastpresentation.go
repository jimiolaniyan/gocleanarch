package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
)

type codecastPresentation struct {
	gateKeeper GateKeeper
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
	return -1
}
