package usecases

import (
	"fmt"
	"github.com/jimiolaniyan/gocleanarch/view"
	"path/filepath"
	"strings"
)

type CodecastSummariesView interface {
	Generate(model *CodecastSummariesViewModel) string
}

type CodecastSummariesViewImpl struct {

}

func (c *CodecastSummariesViewImpl) Generate(model *CodecastSummariesViewModel) string {
	return c.toHTML(model.ViewableCodecastSummaries)
}

func (c CodecastSummariesViewImpl) toHTML(viewableCodecastSummaries []*ViewableCodecastSummary) string {
	frontPageFilePath, err := filepath.Abs("./web/html/frontpage.html")
	checkError(err, fmt.Sprintf("Could not open %s", "./web/html/frontpage.html"))

	codecastPath, err := filepath.Abs("./web/html/codecast.html")
	checkError(err, fmt.Sprintf("Could not open %s", "./web/html/codecast.html"))

	if frontPageTemplate, err := view.CreateTemplate(frontPageFilePath); err == nil {
		var codecastLines strings.Builder

		for _, viewableCodecastSummary := range viewableCodecastSummaries {
			codecastTemplate, _ := view.CreateTemplate(codecastPath)
			codecastTemplate.Replace("title", viewableCodecastSummary.Title)
			codecastTemplate.Replace("publicationDate", viewableCodecastSummary.PublicationDate)
			codecastTemplate.Replace("permalink", viewableCodecastSummary.Permalink)

			//staged
			codecastTemplate.Replace("thumbnail", "https://cleancoders.com/images/portraits/robert-martin.jpg")
			codecastTemplate.Replace("author", "Jimi")
			codecastTemplate.Replace("duration", "58 mins.")
			codecastTemplate.Replace("contentActions", "Buying options go here.")
			codecastLines.WriteString(codecastTemplate.View)
		}

		frontPageTemplate.Replace("codecasts", codecastLines.String())

		return frontPageTemplate.View
	}
	return "Gunk"
}
