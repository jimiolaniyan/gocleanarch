package gocleanarch

import (
	"github.com/segmentio/ksuid"
	"sort"
)

type InMemoryCodecastGateway struct {
	codecasts []*Codecast
}

type InMemoryUserGateway struct {
	users []*User
}

type InMemoryLicenseGateway struct {
	licenses []*License
}

func (lg *InMemoryLicenseGateway) Save(license *License) {
	lg.licenses = append(lg.licenses, license)
}

func (lg *InMemoryLicenseGateway) FindLicensesForUserAndCodecast(user *User, codecast *Codecast) []*License {
	var results []*License
	for _, license := range lg.licenses {
		if license.User().IsSame(&user.Entity) && license.Codecast().IsSame(&codecast.Entity) {
			results = append(results, license)
		}
	}
	return results
}

func (ug *InMemoryUserGateway) FindByName(username string) *User {
	for _, user := range ug.users {
		if user.username == username {
			return user
		}
	}
	return nil
}

func (ug *InMemoryUserGateway) Save(user *User) *User {
	establishId(&user.Entity)
	ug.users = append(ug.users, user)
	return user
}

func (cg *InMemoryCodecastGateway) FindAllCodecastsSortedChronologically() []*Codecast {
	sort.Slice(cg.codecasts, func(i, j int) bool {
		return cg.codecasts[i].PublicationDate().Before(cg.codecasts[j].PublicationDate())
	})
	return cg.codecasts
}

func (cg *InMemoryCodecastGateway) Delete(codecast *Codecast) {
	for i, cc := range cg.codecasts {
		if cc.title == codecast.title {
			cg.codecasts = append(cg.codecasts[:i], cg.codecasts[i+1:]...)
		}
	}
}

func establishId(e *Entity) {
	if e.Id() == "" {
		e.SetId(ksuid.New().String())
	}
}

func (cg *InMemoryCodecastGateway) Save(codecast *Codecast) *Codecast {
	establishId(&codecast.Entity)
	cg.codecasts = append(cg.codecasts, codecast)
	return codecast
}

func (cg *InMemoryCodecastGateway) FindByTitle(codecastTitle string) *Codecast {
	for _, codecast := range cg.codecasts {
		if codecast.title == codecastTitle {
			return codecast
		}
	}
	return nil
}
