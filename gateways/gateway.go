package gateways

import "github.com/jimiolaniyan/gocleanarch/entities"

// CodecastGateway is an interface that defines what persistence operations can be performed on codecasts.
// It belongs in the usecases layer.
type CodecastGateway interface {
	FindAllCodecastsSortedChronologically() []*entities.Codecast
	Delete(codecast *entities.Codecast)
	Save(codecast *entities.Codecast) *entities.Codecast
	FindByTitle(codecastTitle string) *entities.Codecast
	FindByPermalink(permalink string) *entities.Codecast
}

type UserGateway interface {
	Save(user *entities.User) *entities.User
	FindByName(username string) *entities.User
}

type LicenseGateway interface {
	Save(license *entities.License)
	FindLicensesForUserAndCodecast(user *entities.User, codecast *entities.Codecast) []*entities.License
}
