package gocleanarch

// Gateway is an interface that defines what persistence operations are available.
// It belongs in the interface adapters layer.
type Gateway interface {
	FindAllCodecasts() []*Codecast
	Delete(codecast *Codecast)
	SaveCodecast(codecast *Codecast)
	SaveUser(user *User)
	FindUser(username string) *User
	FindCodecastByTitle(codecastTitle string) *Codecast
	SaveLicense(license *License)
	FindLicensesForUserAndCodecast(user *User, codecast *Codecast) []*License
}