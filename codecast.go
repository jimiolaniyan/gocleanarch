package gocleanarch

import "time"

// Codecast is a simple data structure for a codecast. It belongs in the Entities layer.
type Codecast struct {
	Entity
	title           string
	publicationDate time.Time
}

func (c *Codecast) Title() string {
	return c.title
}

func (c *Codecast) PublicationDate() time.Time {
	return c.publicationDate
}

func (c *Codecast) SetTile(title string) {
	c.title = title
}

func (c *Codecast) SetPublicationDate(publicationDate time.Time) {
	c.publicationDate = publicationDate
}
