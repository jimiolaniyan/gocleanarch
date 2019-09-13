package usecases

import "github.com/jimiolaniyan/gocleanarch/entities"

type CodecastSummariesInputBoundary interface {
	SummarizeCodecasts(*entities.User, CodecastSummariesOutputBoundary)
}

type CodecastSummariesOutputBoundary interface {
	GetViewModel() *CodecastSummariesViewModel
}
