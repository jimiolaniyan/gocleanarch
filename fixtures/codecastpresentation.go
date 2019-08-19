package fixtures

import (
	"fmt"
	. "github.com/jimiolaniyan/gocleanarch"
)

type codecastPresentation struct {
	GateKeeper GateKeeper
}

func NewCodecastPresentation() *codecastPresentation {
	AGateway = NewMockGateway()
	return &codecastPresentation{}
}

func (c codecastPresentation) ClearCodecasts() bool {
	var codecasts = AGateway.FindAllCodecasts()

	// TODO not a perfect solution
	for i := len(codecasts)-1; i >= 0; i-- {
		AGateway.Delete(codecasts[i])
	}
	return len(AGateway.FindAllCodecasts()) == 0
}

func (c codecastPresentation) LoginUser(username string) bool {
	user := AGateway.FindUser(username)
	fmt.Println(user)
	if user != nil {
		c.GateKeeper.SetLoggedInUser(user)
		return true
	} else {
		return false
	}
}

func (c codecastPresentation) AddUser(username string) bool {
	AGateway.SaveUser(&User{Username:username})
	return true
}
