package gocleanarch

// Codecast is a simple data structure for a codecast. It belongs in the Entities layer.
type Codecast struct {
	Entity
	title           string
	publicationDate string
}

func (c *Codecast) Title() string {
	return c.title
}

func (c *Codecast) PublicationDate() string {
	return c.publicationDate
}

func (c *Codecast) SetTile(title string) {
	c.title = title
}

func (c *Codecast) SetPublicationDate(publicationDate string) {
	c.publicationDate = publicationDate
}
