package tests

import (
	. "github.com/jimiolaniyan/gocleanarch"
	"time"
)

func SetupContext() {
	AGateway = NewMockGateway()
}

func SetupSampleData() {
	SetupContext()
	//jimi := NewUser("bob")
	//dayo := NewUser("dayo")

	e1 := &Codecast{}
	e1.SetTile("Episode 1 - The Beginning")
	e1.SetPublicationDate(time.Now())

	e2 := &Codecast{}
	e2.SetTile("Episode 2 - The Continuation")
	e2.SetPublicationDate(e1.PublicationDate().Add(1))

	//jimiE1License := NewLicense(Viewing, jimi, e1)
	//jimiE2License := NewLicense(Viewing, jimi, e2)
}
