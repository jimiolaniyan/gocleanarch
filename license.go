package gocleanarch

type License struct {
	Entity
	user     *User
	codecast *Codecast
}

func NewLicense(user *User, codecast *Codecast) *License {
	return &License{user: user, codecast: codecast}
}

func (l *License) User() *User {
	return l.user
}

func (l *License) Codecast() *Codecast {
	return l.codecast
}
