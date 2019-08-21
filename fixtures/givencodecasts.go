package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
)

type GivenCodecast struct {
	Title           string
	PublicationDate string
}

func (gc *GivenCodecast) Execute() {
	codecast := &Codecast{}
	codecast.SetTile(gc.Title)
	codecast.SetPublicationDate(gc.PublicationDate)
	AGateway.SaveCodecast(codecast)
}
