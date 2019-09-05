package gocleanarch

// CodecastGateway is an interface that defines what persistence operations can be performed on codecasts.
// It belongs in the usecases layer.
type CodecastGateway interface {
	FindAllCodecastsSortedChronologically() []*Codecast
	Delete(codecast *Codecast)
	Save(codecast *Codecast) *Codecast
	FindByTitle(codecastTitle string) *Codecast
	FindByPermalink(permalink string) *Codecast
}

type UserGateway interface {
	Save(user *User) *User
	FindByName(username string) *User
}

type LicenseGateway interface {
	Save(license *License)
	FindLicensesForUserAndCodecast(user *User, codecast *Codecast) []*License
}
