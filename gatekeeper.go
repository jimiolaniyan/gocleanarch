package gocleanarch

type GateKeeper struct {
	loggedInUser *User
}

func (gk *GateKeeper) SetLoggedInUser(loggedInUser *User) {
	gk.loggedInUser = loggedInUser
}

func (gk *GateKeeper) LoggedInUser() *User {
	return gk.loggedInUser
}
