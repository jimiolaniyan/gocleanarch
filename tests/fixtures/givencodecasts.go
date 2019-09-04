package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
	"time"
)

type GivenCodecast struct {
	Title           string
	PublicationDate string
	Permalink       string
}

func (gc *GivenCodecast) Execute() bool {
	codecast := &Codecast{}
	codecast.SetTile(gc.Title)

	t, _ := time.Parse("1/2/2006", gc.PublicationDate)

	codecast.SetPublicationDate(t)
	CodecastRepo.Save(codecast)
	return true
}
