package gocleanarch

type License struct {
	Entity
	user        *User
	codecast    *Codecast
	licenseType LicenseType
}

type LicenseType int

const (
	Viewing LicenseType = 1 + iota
	Downloading
)

func NewLicense(lType LicenseType, user *User, codecast *Codecast) *License {
	return &License{licenseType: lType, user: user, codecast: codecast}
}

func (l *License) User() *User {
	return l.user
}

func (l *License) Codecast() *Codecast {
	return l.codecast
}

func (l *License) LicenseType() LicenseType {
	return l.licenseType
}
