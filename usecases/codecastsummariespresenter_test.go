package usecases

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestValidateViewModel(t *testing.T) {
	rm := &CodecastSummariesResponseModel{}
	date := time.Date(2015, 5, 22, 00, 00, 00, 000, time.UTC)
	summary := &codecastSummary{
		Title:           "Title",
		PublicationDate: date,
		Permalink:       "permalink",
		IsViewable:      true,
		IsDownloadable:  false,
	}
	rm.addCodecastSummary(summary)

	presenter := CodecastSummariesPresenter{}
	presenter.Present(rm)

	viewModel := presenter.GetViewModel()
	viewableCodecastSummary := viewModel.ViewableCodecastSummaries[0]
	assert.Equal(t, viewableCodecastSummary.Title, summary.Title)
}