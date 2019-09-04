package gocleanarch

import (
	"time"
)

func SetupContext() {
	AGateway = &MockGateway{}
	AUserGateway = &InMemoryUserGateway{}
}

func SetupSampleData() {
	SetupContext()
	jimi := NewUser("jimi")
	dayo := NewUser("dayo")

	AUserGateway.SaveUser(jimi)
	AUserGateway.SaveUser(dayo)

	e1 := &Codecast{}
	e1.SetTile("Episode 1 - The Beginning")
	e1.SetPublicationDate(time.Now())

	e2 := &Codecast{}
	e2.SetTile("Episode 2 - The Continuation")
	e2.SetPublicationDate(e1.PublicationDate().Add(1))

	AGateway.SaveCodecast(e1)
	AGateway.SaveCodecast(e2)

	jimiE1License := NewLicense(Viewing, jimi, e1)
	jimiE2License := NewLicense(Viewing, jimi, e2)

	AGateway.SaveLicense(jimiE1License)
	AGateway.SaveLicense(jimiE2License)
}
