package gocleanarch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoDifferentUsersAreNotTheSame(t *testing.T) {
	u1 := &User{Username:"u1"}
	u2 := &User{Username:"u2"}

	u1.SetId("u1ID")
	u2.SetId("u2ID")

	assert.False(t, u1.IsSame(u2))
}

func TestOneUserIsTheSameAsItself(t *testing.T) {
	u1 := &User{Username:"u1"}
	assert.True(t, u1.IsSame(u1))
}

func TestUsersWithTheSameIdAreTheSame(t *testing.T) {
	u1 := &User{Username:"u1"}
	u2 := &User{Username:"u2"}

	u1.SetId("u1Id")
	u2.SetId("u1Id")

	assert.True(t, u1.IsSame(u2))
}

