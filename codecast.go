package gocleanarch

// Codecast is a simple data structure for a codecast. It belongs in the Entities layer.
type Codecast struct {
	Title           string
	PublicationDate string
}

func (c *Codecast) isSame(codecast *Codecast) bool {
	return true
}