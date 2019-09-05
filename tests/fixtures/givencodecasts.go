package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/entities"
	"time"
)

type GivenCodecast struct {
	Title           string
	PublicationDate string
	Permalink       string
}

func (gc *GivenCodecast) Execute() bool {
	codecast := &entities.Codecast{}
	codecast.SetTile(gc.Title)

	date, _ := time.Parse("1/2/2006", gc.PublicationDate)

	codecast.SetPublicationDate(date)
	codecast.SetPermalink(gc.Permalink)
	CodecastRepo.Save(codecast)
	return true
}
