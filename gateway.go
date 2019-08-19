package gocleanarch

// Gateway is an interface that defines how
// persistence operations work. It belongs in the interface adapters layer.
type Gateway interface {
	FindAllCodecasts() []*Codecast
	Delete(codecast *Codecast)
	SaveCodecast(codecast *Codecast)
	SaveUser(user *User)
	FindUser(username string) *User
}