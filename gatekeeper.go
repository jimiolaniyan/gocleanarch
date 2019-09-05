package gocleanarch

import "github.com/jimiolaniyan/gocleanarch/entities"

type GateKeeper struct {
	loggedInUser *entities.User
}

func (gk *GateKeeper) SetLoggedInUser(loggedInUser *entities.User) {
	gk.loggedInUser = loggedInUser
}

func (gk *GateKeeper) LoggedInUser() *entities.User {
	return gk.loggedInUser
}
