package tests

import (
	"github.com/jimiolaniyan/gocleanarch/fixtures"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var codecastPresentation = fixtures.NewCodecastPresentation()

func TestClearCodecastsSpec(t *testing.T) {
	Convey("Given no codecasts", t, func() {
		createCodeCasts()
		Convey("When ClearCodecasts called", func() {
			res := codecastPresentation.ClearCodecasts()

			Convey("It should be cleared", func() {
				So(res, ShouldEqual, true)
			})
		})
	})
}

func createCodeCasts() {
	var codecast1 = fixtures.GivenCodecast{Title: "A", PublicationDate: "3/1/2014"}
	var codecast2 = fixtures.GivenCodecast{Title: "B", PublicationDate: "3/2/2014"}
	var codecast3 = fixtures.GivenCodecast{Title: "C", PublicationDate: "2/18/2014"}

	codecast1.Execute()
	codecast2.Execute()
	codecast3.Execute()
}
