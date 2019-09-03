package view

import (
	"io/ioutil"
	"strings"
)

type Template struct {
	View string
}

func (t *Template) Replace(tagName string, replacement string) {
	t.View = strings.Replace(t.View, "${"+tagName+"}", replacement, 1)
}

func CreateTemplate(templateResource string) (*Template, error) {
	file, err := ioutil.ReadFile(templateResource)

	if err != nil {
		return nil, err
	}
	return &Template{string(file)}, nil
}
