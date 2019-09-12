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

type PresentCodecastUseCaseSuite struct {
	suite.Suite
	user     *entities.User
	codecast *entities.Codecast
	useCase  *CodecastSummariesUseCase
}

func (suite *PresentCodecastUseCaseSuite) SetupTest() {
	setup.LoadContext()
	suite.user = gocleanarch.UserRepo.Save(entities.NewUser("Shakespeare"))
	suite.codecast = gocleanarch.CodecastRepo.Save(&entities.Codecast{})
	suite.useCase = new(CodecastSummariesUseCase)
}

func (suite *PresentCodecastUseCaseSuite) TestUserWithoutViewLicense_CannotViewCodecast() {
	licensedToView := suite.useCase.IsLicensedFor(entities.Viewing, suite.user, suite.codecast)
	assert.False(suite.T(), licensedToView, "User should not be licenced to view as there is no license")
}

func (suite *PresentCodecastUseCaseSuite) TestUserWithViewLicense_CanViewCodecast() {
	viewLicence := entities.NewLicense(entities.Viewing, suite.user, suite.codecast)
	gocleanarch.LicenseRepo.Save(viewLicence)
	licensedToView := suite.useCase.IsLicensedFor(entities.Viewing, suite.user, suite.codecast)

	assert.True(suite.T(), licensedToView)
}

func (suite *PresentCodecastUseCaseSuite) TestUserWithoutViewLicense_CannotViewOtherUsersCodecast() {
	otherUser := gocleanarch.UserRepo.Save(entities.NewUser("Atwood"))
	viewLicence := entities.NewLicense(entities.Viewing, suite.user, suite.codecast)
	gocleanarch.LicenseRepo.Save(viewLicence)
	licensedToView := suite.useCase.IsLicensedFor(entities.Viewing, otherUser, suite.codecast)

	assert.False(suite.T(), licensedToView)
}

func (suite *PresentCodecastUseCaseSuite) TestPresentingNoCodecasts() {
	gocleanarch.CodecastRepo.Delete(suite.codecast)
	presentableCodeCasts := suite.useCase.PresentCodecasts(suite.user)

	assert.True(suite.T(), len(presentableCodeCasts) == 0)
}

func (suite *PresentCodecastUseCaseSuite) TestPresentOneCodecast() {
	suite.codecast.SetTile("Some Title")
	suite.codecast.SetPublicationDate(time.Date(2011, 5, 22, 00, 00, 00, 000, time.UTC))
	suite.codecast.SetPermalink("permalink")

	presentableCodeCasts := suite.useCase.PresentCodecasts(suite.user)
	pc := presentableCodeCasts[0]

	assert.True(suite.T(), len(presentableCodeCasts) == 1)
	assert.True(suite.T(), "Some Title" == pc.Title)
	assert.True(suite.T(), "5/22/2011" == pc.PublicationDate)
	assert.True(suite.T(), "permalink" == pc.Permalink)
}

func (suite *PresentCodecastUseCaseSuite) TestPresentedCodecastIsNotViewableIfNoLicense() {
	presentableCodeCasts := suite.useCase.PresentCodecasts(suite.user)
	assert.False(suite.T(), presentableCodeCasts[0].IsViewable)
}

func (suite *PresentCodecastUseCaseSuite) TestPresentedCodecastIsViewableIfLicenseExists() {
	gocleanarch.LicenseRepo.Save(entities.NewLicense(entities.Viewing, suite.user, suite.codecast))
	presentableCodeCasts := suite.useCase.PresentCodecasts(suite.user)
	assert.True(suite.T(), presentableCodeCasts[0].IsViewable)
}

func (suite *PresentCodecastUseCaseSuite) TestPresentedCodecastIsDownloadableIfDownloadLicenseExists() {
	license := entities.NewLicense(entities.Downloading, suite.user, suite.codecast)
	gocleanarch.LicenseRepo.Save(license)
	presentableCodeCasts := suite.useCase.PresentCodecasts(suite.user)
	presentableCodecast := presentableCodeCasts[0]
	assert.True(suite.T(), presentableCodecast.IsDownLoadable)
	assert.False(suite.T(), presentableCodecast.IsViewable)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPresentCodecastUsecaseSuite(t *testing.T) {
	suite.Run(t, new(PresentCodecastUseCaseSuite))
}
