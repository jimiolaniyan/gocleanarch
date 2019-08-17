package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
)

type GivenCodecast struct {
	Title string
	PublicationDate string
}

func (gc *GivenCodecast) Execute() {
	codecast := &Codecast{Title:gc.Title, PublicationDate:gc.PublicationDate}
	AGateway.SaveCodecast(codecast)
}
