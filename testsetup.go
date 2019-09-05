package gocleanarch

import (
	"github.com/jimiolaniyan/gocleanarch/entities"
	"time"
)

func SetupContext() {
	CodecastRepo = &InMemoryCodecastGateway{}
	UserRepo = &InMemoryUserGateway{}
	LicenseRepo = &InMemoryLicenseGateway{}
	SessionKeeper = &GateKeeper{}
}

func SetupSampleData() {
	SetupContext()
	jimi := entities.NewUser("jimi")
	dayo := entities.NewUser("dayo")

	UserRepo.Save(jimi)
	UserRepo.Save(dayo)

	e1 := &entities.Codecast{}
	e1.SetTile("Episode 1 - The Beginning")
	e1.SetPublicationDate(time.Now())

	e2 := &entities.Codecast{}
	e2.SetTile("Episode 2 - The Continuation")
	e2.SetPublicationDate(e1.PublicationDate().Add(1))

	CodecastRepo.Save(e1)
	CodecastRepo.Save(e2)

	jimiE1License := entities.NewLicense(entities.Viewing, jimi, e1)
	jimiE2License := entities.NewLicense(entities.Viewing, jimi, e2)

	LicenseRepo.Save(jimiE1License)
	LicenseRepo.Save(jimiE2License)
}
