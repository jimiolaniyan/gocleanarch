package usecases

import (
	"github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/entities"
	"github.com/jimiolaniyan/gocleanarch/tests/setup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)


type CodecastDetailsUseCaseSuite struct {
	suite.Suite
	user     *entities.User
	codecast *entities.Codecast
	useCase  *CodecastDetailsUseCase
}

func (suite *CodecastDetailsUseCaseSuite) SetupTest() {
	setup.LoadContext()
	suite.user = gocleanarch.UserRepo.Save(entities.NewUser("User"))
	suite.codecast = gocleanarch.CodecastRepo.Save(&entities.Codecast{})
	suite.useCase = new(CodecastDetailsUseCase)
}

func (suite *CodecastDetailsUseCaseSuite) TestCreatesCodecastDetailsPresentation() {
	suite.codecast.SetTile("Codecast")
	suite.codecast.SetPermalink("permalink-a")
	suite.codecast.SetPublicationDate(time.Date(2015, 1, 02, 0, 0, 0, 0, time.UTC))

	useCase := &CodecastDetailsUseCase{}
	details := useCase.RequestCodecastDetails(gocleanarch.SessionKeeper.LoggedInUser(), "permalink-a")

	assert.Equal(suite.T(), "Codecast", details.Title)
	assert.Equal(suite.T(), "1/02/2015", details.PublicationDate)
}

func (suite *CodecastDetailsUseCaseSuite) TestDoesntCrashOnMissingCodecast() {
	useCase := &CodecastDetailsUseCase{}
	details := useCase.RequestCodecastDetails(gocleanarch.SessionKeeper.LoggedInUser(), "missing")

	assert.False(suite.T(), details.Found)
}

func TestCodecastDetailsUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CodecastDetailsUseCaseSuite))
}