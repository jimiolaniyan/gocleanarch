package gocleanarch

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)


type CodecastDetailsUseCaseSuite struct {
	suite.Suite
	user     *User
	codecast *Codecast
	useCase  *CodecastDetailsUseCase
}

func (suite *CodecastDetailsUseCaseSuite) SetupTest() {
	SetupContext()
	suite.user = UserRepo.Save(NewUser("User"))
	suite.codecast = CodecastRepo.Save(&Codecast{})
	suite.useCase = new(CodecastDetailsUseCase)
}

func (suite *CodecastDetailsUseCaseSuite) TestCreatesCodecastDetailsPresentation() {
	suite.codecast.SetTile("Codecast")
	suite.codecast.SetPermalink("permalink-a")
	suite.codecast.SetPublicationDate(time.Date(2015, 1, 02, 0, 0, 0, 0, time.UTC))

	useCase := &CodecastDetailsUseCase{}
	details := useCase.RequestCodecastDetails(SessionKeeper.loggedInUser, "permalink-a")

	assert.Equal(suite.T(), "Codecast", details.Title)
	assert.Equal(suite.T(), "1/02/2015", details.PublicationDate)
}

func (suite *CodecastDetailsUseCaseSuite) TestDoesntCrashOnMissingCodecast() {
	useCase := &CodecastDetailsUseCase{}
	details := useCase.RequestCodecastDetails(SessionKeeper.loggedInUser, "missing")

	assert.False(suite.T(), details.Found)
}

func TestCodecastDetailsUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CodecastDetailsUseCaseSuite))
}