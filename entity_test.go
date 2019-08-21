package gocleanarch

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
}

func (suite *UserTestSuite) TestTwoDifferentEntitiesAreNotTheSame() {
	e1 := &Entity{}
	e2 := &Entity{}

	e1.SetId("e1ID")
	e2.SetId("e2ID")

	assert.False(suite.T(), e1.IsSame(e2))
}

func (suite *UserTestSuite) TestOneEntityIsTheSameAsItself() {
	e1 := &Entity{}
	e1.SetId("e1Id")

	assert.True(suite.T(), e1.IsSame(e1))
}

func (suite *UserTestSuite) TestEntitiesWithTheSameIdAreTheSame() {
	e1 := &Entity{}
	e2 := &Entity{}

	e1.SetId("e1Id")
	e2.SetId("e1Id")

	assert.True(suite.T(), e1.IsSame(e2))
}

func (suite *UserTestSuite) TestEntitiesWithNullIdsAreNeverTheSame() {
	e1 := &Entity{}
	e2 := &Entity{}

	assert.False(suite.T(), e1.IsSame(e2))
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
