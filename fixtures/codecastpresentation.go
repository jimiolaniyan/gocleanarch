package fixtures

import (
	. "github.com/jimiolaniyan/gocleanarch"
)

type codecastPresentation struct {
}

func NewCodecastPresentation() *codecastPresentation {
	AGateway = NewMockGateway()
	return &codecastPresentation{}
}

func (c codecastPresentation) ClearCodecasts() bool {
	var codecasts = AGateway.FindAllCodecasts()

	// TODO not a perfect solution
	for i := len(codecasts)-1; i >= 0; i-- {
		AGateway.Delete(codecasts[i])
	}
	return len(AGateway.FindAllCodecasts()) == 0
}
