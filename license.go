package gocleanarch

type License struct {
	Entity
	user     *User
	codecast *Codecast
}

type Licenser interface {
	User() *User
	Codecast() *Codecast
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

type DownloadLicense struct {
	License
}

func NewDownloadLicense(user *User, codecast *Codecast) *DownloadLicense {
	return &DownloadLicense{License{user: user, codecast: codecast}}
}
