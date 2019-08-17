package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
)

type GivenCodecast struct {
	Title string
	PublicationDate string
}

//func (gc *GivenCodecast) SetTitle(title string)  {
//	gc.title = title
//}
//
//func (gc *GivenCodecast) SetPublished(publicationDate string)  {
//	gc.publicationDate = publicationDate
//}

func (gc *GivenCodecast) Execute() {
	codecast := &Codecast{}
	codecast.SetTitle(gc.Title)
	codecast.SetPublicationDate(gc.PublicationDate)
	AGateway.Save(codecast)
}
