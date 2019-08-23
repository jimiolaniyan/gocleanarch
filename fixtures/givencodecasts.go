package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
	"time"
)

type GivenCodecast struct {
	Title           string
	PublicationDate string
}

func (gc *GivenCodecast) Execute() {
	codecast := &Codecast{}
	codecast.SetTile(gc.Title)

	t, _ := time.Parse("01/02/2006", gc.PublicationDate)
	codecast.SetPublicationDate(t)
	AGateway.SaveCodecast(codecast)
}
