package gocleanarch

type Gateway interface {
	FindAllCodecasts() []*Codecast
	Delete(codecast *Codecast)
	SaveCodecast(codecast *Codecast)
	SaveUser(user *User)
}
