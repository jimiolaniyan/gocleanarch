package gocleanarch

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PresentCodecastUsecaseTestSuite struct {
	suite.Suite
	user     *User
	codecast *Codecast
	useCase  *PresentCodecastUseCase
}

func (suite *PresentCodecastUsecaseTestSuite) SetupTest() {
	AGateway = NewMockGateway()
	suite.user = AGateway.SaveUser(&User{Username: "Shakespeare"})
	suite.codecast = AGateway.SaveCodecast(&Codecast{})
	suite.useCase = new(PresentCodecastUseCase)
}

func (suite *PresentCodecastUsecaseTestSuite) TestUserWithoutViewLicense_CannotViewCodecast() {
	//AGateway = NewMockGateway()
	suite.user.Username = "Jimi"
	licensedToView := suite.useCase.IsLicensedToViewCodecast(suite.user, suite.codecast)
	assert.False(suite.T(), licensedToView, "User should not be licenced to view as there is no license")
}

func (suite *PresentCodecastUsecaseTestSuite) TestUserWithViewLicense_CanViewCodecast() {

	viewLicence := &License{User: suite.user, Codecast: suite.codecast}

	AGateway.SaveLicense(viewLicence)

	licensedToView := suite.useCase.IsLicensedToViewCodecast(suite.user, suite.codecast)
	assert.True(suite.T(), licensedToView)
}

func (suite *PresentCodecastUsecaseTestSuite) TestUserWithoutViewLicense_CannotViewOtherUsersCodecast() {
	otherUser := AGateway.SaveUser(&User{Username: "Atwood"})

	viewLicence := &License{User: suite.user, Codecast: suite.codecast}
	AGateway.SaveLicense(viewLicence)

	licensedToView := suite.useCase.IsLicensedToViewCodecast(otherUser, suite.codecast)

	assert.False(suite.T(), licensedToView)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPresentCodecastUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(PresentCodecastUsecaseTestSuite))
}
