package gocleanarch

import (
	"github.com/segmentio/ksuid"
)

// MockGateway is a mock implementation of the Gateway
type MockGateway struct {
	codecasts []*Codecast
	users     []*User
	licenses  []*License
}

func NewMockGateway() *MockGateway {
	return &MockGateway{codecasts: []*Codecast{}}
}

func (m *MockGateway) FindAllCodecasts() []*Codecast {
	return m.codecasts
}

func (m *MockGateway) Delete(codecast *Codecast) {
	for i, cc := range m.codecasts {
		if cc.title == codecast.title {
			m.codecasts = append(m.codecasts[:i], m.codecasts[i+1:]...)
		}
	}
}

func (m *MockGateway) SaveUser(user *User) *User {
	establishId(&user.Entity)
	m.users = append(m.users, user)
	return user
}

func establishId(e *Entity) {
	if e.Id() == "" {
		e.SetId(ksuid.New().String())
	}
}

func (m *MockGateway) SaveLicense(license *License) {
	m.licenses = append(m.licenses, license)
}

func (m *MockGateway) SaveCodecast(codecast *Codecast) *Codecast {
	establishId(&codecast.Entity)
	m.codecasts = append(m.codecasts, codecast)
	return codecast
}

func (m *MockGateway) FindUser(username string) *User {
	for _, user := range m.users {
		if user.username == username {
			return user
		}
	}
	return nil
}

func (m *MockGateway) FindCodecastByTitle(codecastTitle string) *Codecast {
	for _, codecast := range m.codecasts {
		if codecast.title == codecastTitle {
			return codecast
		}
	}
	return nil
}

func (m *MockGateway) FindLicensesForUserAndCodecast(user *User, codecast *Codecast) []*License {
	var results []*License
	for _, license := range m.licenses {
		if license.User().IsSame(&user.Entity) && license.Codecast().IsSame(&codecast.Entity) {
			results = append(results, license)
		}
	}
	return results
}
