package view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoReplacement(t *testing.T) {
	content := "some static content"
	template := Template{content}
	assert.Equal(t, content, template.View)
}

func TestSimpleReplacement(t *testing.T) {
	template := Template{View: "replace ${this}."}
	template.Replace("this", "that")
	assert.Equal(t, "replace that.", template.View)
}
