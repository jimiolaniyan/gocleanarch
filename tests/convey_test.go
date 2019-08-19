package tests

import (
	"github.com/jimiolaniyan/gocleanarch/fixtures"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var codecastPresentation = fixtures.NewCodecastPresentation()

func TestPresentNoCodecastsSpec(t *testing.T) {
	user := "U"
	Convey("Given no codecasts", t, func() {
		res := codecastPresentation.ClearCodecasts()
		So(res, ShouldEqual, true)
	})

	Convey("And a user U", t, func() {
		status := codecastPresentation.AddUser(user)
		So(status, ShouldEqual, true)
	})

	Convey("With U logged in", t, func() {
		out := codecastPresentation.LoginUser(user)
		So(out, ShouldEqual, true)

		Convey("Then presentation user will be U", func() {
			currentUser := codecastPresentation.PresentationUser()
			So(currentUser, ShouldEqual, user)
		})


		Convey("And there will be no codecasts presented", func() {
			count := codecastPresentation.CountOfCodecastsPresented()
			So(count, ShouldEqual, 0)
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
