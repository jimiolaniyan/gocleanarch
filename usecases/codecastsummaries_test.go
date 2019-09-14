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
	presenterSpy *CodecastSummaryOutputBoundarySpy
}

func (suite *PresentCodecastUseCaseSuite) SetupTest() {
	setup.LoadContext()
	suite.user = gocleanarch.UserRepo.Save(entities.NewUser("Shakespeare"))
	suite.codecast = gocleanarch.CodecastRepo.Save(&entities.Codecast{})
	suite.useCase = new(CodecastSummariesUseCase)
	suite.presenterSpy = &CodecastSummaryOutputBoundarySpy{}
}

func (suite *PresentCodecastUseCaseSuite) TestUseCaseWiring() {
	suite.useCase.SummarizeCodecasts(suite.user, suite.presenterSpy)
	assert.NotNil(suite.T(), suite.presenterSpy.ResponseModel)
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
	suite.useCase.SummarizeCodecasts(suite.user, suite.presenterSpy)

	assert.True(suite.T(), len(suite.presenterSpy.ResponseModel.CodecastSummaries) == 0)
}

func (suite *PresentCodecastUseCaseSuite) TestPresentOneCodecast() {
	suite.codecast.SetTile("Some Title")
	date := time.Date(2011, 5, 22, 00, 00, 00, 000, time.UTC)
	suite.codecast.SetPublicationDate(date)
	suite.codecast.SetPermalink("permalink")

	presenterSpy := &CodecastSummaryOutputBoundarySpy{}
	suite.useCase.SummarizeCodecasts(suite.user, presenterSpy)

	summaries := presenterSpy.ResponseModel.CodecastSummaries;
	summary := summaries[0]

	assert.True(suite.T(), len(summaries) == 1)
	assert.True(suite.T(), "Some Title" == summary.Title)
	assert.True(suite.T(), date == summary.PublicationDate)
	assert.True(suite.T(), "permalink" == summary.Permalink)
}

func (suite *PresentCodecastUseCaseSuite) TestPresentedCodecastIsNotViewableIfNoLicense() {
	suite.useCase.SummarizeCodecasts(suite.user, suite.presenterSpy)
	summary := suite.presenterSpy.ResponseModel.CodecastSummaries[0]
	assert.False(suite.T(), summary.IsViewable)
}

func (suite *PresentCodecastUseCaseSuite) TestPresentedCodecastIsViewableIfLicenseExists() {
	gocleanarch.LicenseRepo.Save(entities.NewLicense(entities.Viewing, suite.user, suite.codecast))
	suite.useCase.SummarizeCodecasts(suite.user, suite.presenterSpy)
	summary := suite.presenterSpy.ResponseModel.CodecastSummaries[0]
	assert.True(suite.T(), summary.IsViewable)
}

func (suite *PresentCodecastUseCaseSuite) TestPresentedCodecastIsDownloadableIfDownloadLicenseExists() {
	license := entities.NewLicense(entities.Downloading, suite.user, suite.codecast)
	gocleanarch.LicenseRepo.Save(license)
	suite.useCase.SummarizeCodecasts(suite.user, suite.presenterSpy)
	summary := suite.presenterSpy.ResponseModel.CodecastSummaries[0]
	assert.True(suite.T(), summary.IsDownloadable)
	assert.False(suite.T(), summary.IsViewable)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPresentCodecastUsecaseSuite(t *testing.T) {
	suite.Run(t, new(PresentCodecastUseCaseSuite))
}
