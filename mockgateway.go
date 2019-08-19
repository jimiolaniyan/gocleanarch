package gocleanarch

type MockGateway struct {
	codecasts []*Codecast
	users []*User
}

func NewMockGateway() *MockGateway{
	 return &MockGateway{codecasts: []*Codecast{}}
}

func (m *MockGateway) SaveCodecast(codecast *Codecast) {
	m.codecasts = append(m.codecasts, codecast)
}

func (m *MockGateway) FindAllCodecasts() []*Codecast {
	return m.codecasts
}

func (m *MockGateway) Delete(codecast *Codecast) {
	for i, cc := range m.codecasts {
		if cc.Title == codecast.Title {
			m.codecasts = append(m.codecasts[:i], m.codecasts[i+1:]...)
		}
	}
}

func (m *MockGateway) SaveUser(user *User) {
	m.users = append(m.users, user)
}

func (m *MockGateway) FindUser(username string) *User {
	for _, user := range m.users {
		if user.Username == username {
			return user
		}
	}
	return nil
}