package gocleanarch

type Codecast struct {
	title string
	publicationDate string
}

func (c *Codecast) SetTitle(title string) {
	c.title = title
}

func (c *Codecast) SetPublicationDate(publicationDate string) {
	c.publicationDate = publicationDate
}
