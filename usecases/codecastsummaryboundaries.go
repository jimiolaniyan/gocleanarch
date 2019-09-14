package usecases

import "github.com/jimiolaniyan/gocleanarch/entities"

type CodecastSummariesInputBoundary interface {
	SummarizeCodecasts(loggedInUser *entities.User, presenter CodecastSummariesOutputBoundary)
}

type CodecastSummariesOutputBoundary interface {
	GetViewModel() *CodecastSummariesViewModel
	Present(responseModel *CodecastSummariesResponseModel)
}
