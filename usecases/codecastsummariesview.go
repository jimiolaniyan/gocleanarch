package usecases

import (
	"fmt"
	"github.com/jimiolaniyan/gocleanarch/view"
	"path/filepath"
	"strings"
)

type CodecastSummariesView interface {
	Generate(model *CodecastSummaryResponseModel) string
}

type CodecastSummariesViewImpl struct {

}

func (c *CodecastSummariesViewImpl) Generate(model *CodecastSummaryResponseModel) string {
	return ""
}

func (c CodecastSummariesViewImpl) toHTML(presentableCodecasts []*PresentableCodecastSummary) string {
	frontPageFilePath, err := filepath.Abs("./web/html/frontpage.html")
	checkError(err, fmt.Sprintf("Could not open %s", "./web/html/frontpage.html"))

	codecastPath, err := filepath.Abs("./web/html/codecast.html")
	checkError(err, fmt.Sprintf("Could not open %s", "./web/html/codecast.html"))

	if frontPageTemplate, err := view.CreateTemplate(frontPageFilePath); err == nil {
		var codecastLines strings.Builder

		for _, pc := range presentableCodecasts {
			codecastTemplate, _ := view.CreateTemplate(codecastPath)
			codecastTemplate.Replace("title", pc.Title)
			codecastTemplate.Replace("publicationDate", pc.PublicationDate)
			codecastTemplate.Replace("permalink", pc.Permalink)

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
