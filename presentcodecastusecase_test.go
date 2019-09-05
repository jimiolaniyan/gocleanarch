package gocleanarch

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type PresentCodecastUsecaseSuite struct {
	suite.Suite
	user     *User
	codecast *Codecast
	useCase  *CodecastSummaryUseCase
}

func (suite *PresentCodecastUsecaseSuite) SetupTest() {
	SetupContext()
	suite.user = UserRepo.Save(NewUser("Shakespeare"))
	suite.codecast = CodecastRepo.Save(&Codecast{})
	suite.useCase = new(CodecastSummaryUseCase)
}

func (suite *PresentCodecastUsecaseSuite) TestUserWithoutViewLicense_CannotViewCodecast() {
	licensedToView := suite.useCase.IsLicensedFor(Viewing, suite.user, suite.codecast)
	assert.False(suite.T(), licensedToView, "User should not be licenced to view as there is no license")
}

func (suite *PresentCodecastUsecaseSuite) TestUserWithViewLicense_CanViewCodecast() {
	viewLicence := NewLicense(Viewing, suite.user, suite.codecast)
	LicenseRepo.Save(viewLicence)
	licensedToView := suite.useCase.IsLicensedFor(Viewing, suite.user, suite.codecast)

	assert.True(suite.T(), licensedToView)
}

func (suite *PresentCodecastUsecaseSuite) TestUserWithoutViewLicense_CannotViewOtherUsersCodecast() {
	otherUser := UserRepo.Save(NewUser("Atwood"))
	viewLicence := NewLicense(Viewing, suite.user, suite.codecast)
	LicenseRepo.Save(viewLicence)
	licensedToView := suite.useCase.IsLicensedFor(Viewing, otherUser, suite.codecast)

	assert.False(suite.T(), licensedToView)
}

func (suite *PresentCodecastUsecaseSuite) TestPresentingNoCodecasts() {
	CodecastRepo.Delete(suite.codecast)
	presentableCodeCasts := suite.useCase.PresentCodecasts(suite.user)

	assert.True(suite.T(), len(presentableCodeCasts) == 0)
}

func (suite *PresentCodecastUsecaseSuite) TestPresentOneCodecast() {
	suite.codecast.SetTile("Some Title")
	date := time.Date(2011, 5, 22, 00, 00, 00, 000, time.UTC)
	suite.codecast.SetPublicationDate(date)
	presentableCodeCasts := suite.useCase.PresentCodecasts(suite.user)
	pc := presentableCodeCasts[0]

	assert.True(suite.T(), len(presentableCodeCasts) == 1)
	assert.True(suite.T(), "Some Title" == pc.Title)
	assert.True(suite.T(), date.Format("1/2/2006") == pc.PublicationDate)
}

func (suite *PresentCodecastUsecaseSuite) TestPresentedCodecastIsNotViewableIfNoLicense() {
	presentableCodeCasts := suite.useCase.PresentCodecasts(suite.user)
	assert.False(suite.T(), presentableCodeCasts[0].IsViewable)
}

func (suite *PresentCodecastUsecaseSuite) TestPresentedCodecastIsViewableIfLicenseExists() {
	LicenseRepo.Save(NewLicense(Viewing, suite.user, suite.codecast))
	presentableCodeCasts := suite.useCase.PresentCodecasts(suite.user)
	assert.True(suite.T(), presentableCodeCasts[0].IsViewable)
}

func (suite *PresentCodecastUsecaseSuite) TestPresentedCodecastIsDownloadableIfDownloadLicenseExists() {
	license := NewLicense(Downloading, suite.user, suite.codecast)
	LicenseRepo.Save(license)
	presentableCodeCasts := suite.useCase.PresentCodecasts(suite.user)
	presentableCodecast := presentableCodeCasts[0]
	assert.True(suite.T(), presentableCodecast.IsDownLoadable)
	assert.False(suite.T(), presentableCodecast.IsViewable)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPresentCodecastUsecaseSuite(t *testing.T) {
	suite.Run(t, new(PresentCodecastUsecaseSuite))
}
