package gocleanarch

type Gateway interface {
	FindAllCodecasts() []*Codecast
	Delete(codecast *Codecast)
	Save(codecast *Codecast)
}
