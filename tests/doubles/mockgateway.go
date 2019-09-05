package doubles

import (
	"github.com/jimiolaniyan/gocleanarch/entities"
	"github.com/segmentio/ksuid"
	"sort"
)

type InMemoryLicenseGateway struct {
	licenses []*entities.License
}

type InMemoryUserGateway struct {
	users []*entities.User
}

type InMemoryCodecastGateway struct {
	codecasts []*entities.Codecast
}

func (lg *InMemoryLicenseGateway) Save(license *entities.License) {
	lg.licenses = append(lg.licenses, license)
}

func (lg *InMemoryLicenseGateway) FindLicensesForUserAndCodecast(user *entities.User, codecast *entities.Codecast) []*entities.License {
	var results []*entities.License
	for _, license := range lg.licenses {
		if license.User().IsSame(&user.Entity) && license.Codecast().IsSame(&codecast.Entity) {
			results = append(results, license)
		}
	}
	return results
}

func (ug *InMemoryUserGateway) FindByName(username string) *entities.User {
	for _, user := range ug.users {
		if user.Username() == username {
			return user
		}
	}
	return nil
}

func (ug *InMemoryUserGateway) Save(user *entities.User) *entities.User {
	establishId(&user.Entity)
	ug.users = append(ug.users, user)
	return user
}

func (cg *InMemoryCodecastGateway) FindAllCodecastsSortedChronologically() []*entities.Codecast {
	sort.Slice(cg.codecasts, func(i, j int) bool {
		return cg.codecasts[i].PublicationDate().Before(cg.codecasts[j].PublicationDate())
	})
	return cg.codecasts
}

func (cg *InMemoryCodecastGateway) Delete(codecast *entities.Codecast) {
	for i, cc := range cg.codecasts {
		if cc.Title() == codecast.Title() {
			cg.codecasts = append(cg.codecasts[:i], cg.codecasts[i+1:]...)
		}
	}
}

func establishId(e *entities.Entity) {
	if e.Id() == "" {
		e.SetId(ksuid.New().String())
	}
}

func (cg *InMemoryCodecastGateway) Save(codecast *entities.Codecast) *entities.Codecast {
	establishId(&codecast.Entity)
	cg.codecasts = append(cg.codecasts, codecast)
	return codecast
}

func (cg *InMemoryCodecastGateway) FindByTitle(codecastTitle string) *entities.Codecast {
	for _, codecast := range cg.codecasts {
		if codecast.Title() == codecastTitle {
			return codecast
		}
	}
	return nil
}

func (cg *InMemoryCodecastGateway) FindByPermalink(permalink string) *entities.Codecast {
	for _, codecast := range cg.codecasts {
		if codecast.Permalink() == permalink {
			return codecast
		}
	}
	return nil
}
