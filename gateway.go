package gocleanarch

// CodecastGateway is an interface that defines what persistence operations are available.
// It belongs in the interface adapters layer.
type CodecastGateway interface {
	FindAllCodecastsSortedChronologically() []*Codecast
	Delete(codecast *Codecast)
	Save(codecast *Codecast) *Codecast
	FindByTitle(codecastTitle string) *Codecast
}

type UserGateway interface {
	Save(user *User) *User
	FindByName(username string) *User
}

type LicenseGateway interface {
	Save(license *License)
	FindLicensesForUserAndCodecast(user *User, codecast *Codecast) []*License
}
