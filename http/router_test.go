package http

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RouterTestSuite struct {
	suite.Suite
	router *Router
}

var actualRequest = &ParsedRequest{}

func (suite *RouterTestSuite) SetupTest() {
	suite.router = NewRouter()
}

func (suite *RouterTestSuite) TestSimplePath() {
	req := &ParsedRequest{Method: "GET", Path: "/it"}
	suite.router.AddPath("it", &TestController{})

	suite.router.Route(req)

	assert.Equal(suite.T(), actualRequest, req)
}

func (suite *RouterTestSuite) TestPathWithDynamicData() {
	req := &ParsedRequest{Method: "GET", Path: "/a/b/c"}
	suite.router.AddPath("a", &TestController{})

	suite.router.Route(req)

	assert.Equal(suite.T(), actualRequest, req)
}

func (suite *RouterTestSuite) TestRootPath() {
	req := &ParsedRequest{Method: "GET", Path: "/"}
	suite.router.AddPath("", &TestController{})

	suite.router.Route(req)

	assert.Equal(suite.T(), actualRequest, req)
}

func (suite *RouterTestSuite) Test404() {
	request := &ParsedRequest{Method: "GET", Path: "/something-missing"}

	result := suite.router.Route(request)

	assert.Equal(suite.T(), "HTTP/1.1 404 Not Found", result)
}

type TestController struct {
}

func (t *TestController) Handle(request *ParsedRequest) string {
	actualRequest = request
	return ""
}

func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}
