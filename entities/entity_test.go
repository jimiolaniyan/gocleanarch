package entities

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EntitySuite struct {
	suite.Suite
}

func (suite *EntitySuite) TestTwoDifferentEntitiesAreNotTheSame() {
	e1 := &Entity{}
	e2 := &Entity{}

	e1.SetId("e1ID")
	e2.SetId("e2ID")

	assert.False(suite.T(), e1.IsSame(e2))
}

func (suite *EntitySuite) TestOneEntityIsTheSameAsItself() {
	e1 := &Entity{}
	e1.SetId("e1Id")

	assert.True(suite.T(), e1.IsSame(e1))
}

func (suite *EntitySuite) TestEntitiesWithTheSameIdAreTheSame() {
	e1 := &Entity{}
	e2 := &Entity{}

	e1.SetId("e1Id")
	e2.SetId("e1Id")

	assert.True(suite.T(), e1.IsSame(e2))
}

func (suite *EntitySuite) TestEntitiesWithNullIdsAreNeverTheSame() {
	e1 := &Entity{}
	e2 := &Entity{}

	assert.False(suite.T(), e1.IsSame(e2))
}

func TestEntitySuite(t *testing.T) {
	suite.Run(t, new(EntitySuite))
}
