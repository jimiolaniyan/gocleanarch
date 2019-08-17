package gocleanarch

type MockGateway struct {
	codecasts []*Codecast
}

func NewMockGateway() *MockGateway{
	 return &MockGateway{codecasts: []*Codecast{}}
}

func (m *MockGateway) Save(codecast *Codecast) {
	m.codecasts = append(m.codecasts, codecast)
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
