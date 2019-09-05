package gocleanarch

import "time"

type Codecast struct {
	Entity
	title           string
	publicationDate time.Time
	permalink string
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

func (c *Codecast) SetPermalink(permalink string) {
	c.permalink = permalink
}
