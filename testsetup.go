package gocleanarch

import (
	"time"
)

func SetupContext() {
	CodecastRepo = &MockGateway{}
	UserRepo = &InMemoryUserGateway{}
	LicenseRepo = &InMemoryLicenseGateway{}
}

func SetupSampleData() {
	SetupContext()
	jimi := NewUser("jimi")
	dayo := NewUser("dayo")

	UserRepo.Save(jimi)
	UserRepo.Save(dayo)

	e1 := &Codecast{}
	e1.SetTile("Episode 1 - The Beginning")
	e1.SetPublicationDate(time.Now())

	e2 := &Codecast{}
	e2.SetTile("Episode 2 - The Continuation")
	e2.SetPublicationDate(e1.PublicationDate().Add(1))

	CodecastRepo.Save(e1)
	CodecastRepo.Save(e2)

	jimiE1License := NewLicense(Viewing, jimi, e1)
	jimiE2License := NewLicense(Viewing, jimi, e2)

	LicenseRepo.Save(jimiE1License)
	LicenseRepo.Save(jimiE2License)
}
