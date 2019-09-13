package usecases

import (
	"github.com/jimiolaniyan/gocleanarch"
	"github.com/jimiolaniyan/gocleanarch/entities"
	. "github.com/jimiolaniyan/gocleanarch/http"
	"github.com/jimiolaniyan/gocleanarch/tests/setup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CodecastSummaryInputBoundarySpy struct {
	requestedUser               *entities.User
	summarizeCodecastsWasCalled bool
	OutputBoundary              CodecastSummariesOutputBoundary
}

func (c *CodecastSummaryInputBoundarySpy) SummarizeCodecasts(loggedInUser *entities.User, presenter CodecastSummariesOutputBoundary) {
	c.summarizeCodecastsWasCalled = true
	c.requestedUser = loggedInUser
	c.OutputBoundary = presenter
}

type CodecastSummaryOutputBoundarySpy struct {
	ResponseModel *CodecastSummaryResponseModel
}

func (c *CodecastSummaryOutputBoundarySpy) GetResponseModel() *CodecastSummaryResponseModel {
	return c.ResponseModel
}

type CodecastSummaryViewSpy struct {
	generateViewWasCalled bool
	ResponseModel         *CodecastSummaryResponseModel
}

func (c *CodecastSummaryViewSpy) Generate(model *CodecastSummaryResponseModel) string {
	c.ResponseModel = model
	c.generateViewWasCalled = true
	return  ""
}

type CodecastSummariesControllerTestSuite struct {
	suite.Suite
	useCaseSpy   *CodecastSummaryInputBoundarySpy
	presenterSpy *CodecastSummaryOutputBoundarySpy
	viewSpy      *CodecastSummaryViewSpy
	controller   *CodecastSummariesController
}

func (suite *CodecastSummariesControllerTestSuite) SetupTest() {
	setup.LoadSampleData()
	suite.useCaseSpy = &CodecastSummaryInputBoundarySpy{}
	suite.presenterSpy = &CodecastSummaryOutputBoundarySpy{}
	suite.viewSpy = &CodecastSummaryViewSpy{}
	suite.controller = &CodecastSummariesController{UseCase: suite.useCaseSpy, Presenter: suite.presenterSpy, View: suite.viewSpy}
}

func (suite *CodecastSummariesControllerTestSuite) TestInputBoundaryInvocation() {
	request := &ParsedRequest{Method: "GET", Path: "bla"}
	loggedInUser := gocleanarch.UserRepo.FindByName("jimi")

	suite.controller.Handle(request)

	assert.True(suite.T(), suite.useCaseSpy.summarizeCodecastsWasCalled)
	assert.Equal(suite.T(), loggedInUser, suite.useCaseSpy.requestedUser)
	assert.Equal(suite.T(), suite.presenterSpy, suite.useCaseSpy.OutputBoundary)
}

func (suite *CodecastSummariesControllerTestSuite) TestControllerSendsTheResponseModelToTheView() {
	request := &ParsedRequest{Method: "GET", Path: "bla"}
	suite.presenterSpy.ResponseModel = &CodecastSummaryResponseModel{}

	suite.controller.Handle(request)

	assert.True(suite.T(), suite.viewSpy.generateViewWasCalled)
	assert.Equal(suite.T(), suite.presenterSpy.ResponseModel, suite.viewSpy.ResponseModel)
}

func TestCodecastSummariesControllerSuite(t *testing.T) {
	suite.Run(t, new(CodecastSummariesControllerTestSuite))
}
