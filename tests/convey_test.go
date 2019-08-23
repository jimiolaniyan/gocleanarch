package tests

import (
	"github.com/jimiolaniyan/gocleanarch/fixtures"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPresentNoCodecastsSpec(t *testing.T) {
	user := "U"
	Convey("Given no codecasts", t, func() {
		res := fixtures.CodecastPresentation.ClearCodecasts()
		So(res, ShouldEqual, true)
	})

	Convey("And a user U", t, func() {
		status := fixtures.CodecastPresentation.AddUser(user)
		So(status, ShouldEqual, true)
	})

	Convey("With U logged in", t, func() {
		out := fixtures.CodecastPresentation.LoginUser(user)
		So(out, ShouldEqual, true)

		Convey("Then presentation user will be U", func() {
			currentUser := fixtures.CodecastPresentation.PresentationUser()
			So(currentUser, ShouldEqual, user)
		})

		Convey("And there will be no codecasts presented", func() {
			count := fixtures.CodecastPresentation.CountOfCodecastsPresented()
			So(count, ShouldEqual, 0)
		})
		Reset(func() {
			fixtures.CodecastPresentation.LogOutUser()
		})
	})

}

func TestPresentViewableCodecastsSpec(t *testing.T) {
	user := "U"
	codecast := "A"
	Convey("Given codecasts", t, func() {
		res := createCodeCasts()
		So(res, ShouldEqual, true)
	})

	Convey("And a user U", t, func() {
		status := fixtures.CodecastPresentation.AddUser(user)
		So(status, ShouldEqual, true)
	})

	Convey("With U logged in", t, func() {
		out := fixtures.CodecastPresentation.LoginUser(user)
		So(out, ShouldEqual, true)
	})

	Convey("And with license for U able to view A", t, func() {
		status := fixtures.CodecastPresentation.CreateLicenceForViewing(user, codecast)
		So(status, ShouldEqual, true)

		Convey("Then the following codecasts will be presented for U", func() {
			currentUser := fixtures.CodecastPresentation.PresentationUser()
			So(currentUser, ShouldEqual, user)

			presentedCodecasts := fixtures.Query()
			expected := []fixtures.QueryResponse{
				{Title: "C", Picture: "C", Description: "C", Viewable: false},
				{Title: "A", Picture: "A", Description: "A", Viewable: true},
				{Title: "B", Picture: "B", Description: "B", Viewable: false},
			}
			So(presentedCodecasts, ShouldResemble, expected)

			Reset(func() {
				fixtures.CodecastPresentation.LogOutUser()
				fixtures.CodecastPresentation.ClearCodecasts()
			})
		})
	})
}

func TestPresentDownloadableCodecastsSpec(t *testing.T) {
	user := "U"
	codecast := "A"
	Convey("Given codecasts", t, func() {
		res := createCodeCasts()
		So(res, ShouldEqual, true)
	})

	Convey("And a user U", t, func() {
		status := fixtures.CodecastPresentation.AddUser(user)
		So(status, ShouldEqual, true)
	})

	Convey("With U logged in", t, func() {
		out := fixtures.CodecastPresentation.LoginUser(user)
		So(out, ShouldEqual, true)
	})

	Convey("And with license for U able to download A", t, func() {
		status := fixtures.CodecastPresentation.CreateLicenceForDownloading(user, codecast)
		So(status, ShouldEqual, true)

		Convey("Then the following codecasts will be presented for U", func() {
			currentUser := fixtures.CodecastPresentation.PresentationUser()
			So(currentUser, ShouldEqual, user)

			presentedCodecasts := fixtures.Query()
			expected := []fixtures.QueryResponse{
				{Title: "C", Picture: "C", Description: "C", Viewable: false},
				{Title: "A", Picture: "A", Description: "A", Viewable: true, Downloadable: true},
				{Title: "B", Picture: "B", Description: "B", Viewable: false},
			}
			So(presentedCodecasts, ShouldResemble, expected)
		})
	})
}

func createCodeCasts() bool {
	var codecast1 = fixtures.GivenCodecast{Title: "A", PublicationDate: "3/1/2014"}
	var codecast2 = fixtures.GivenCodecast{Title: "B", PublicationDate: "3/2/2014"}
	var codecast3 = fixtures.GivenCodecast{Title: "C", PublicationDate: "2/18/2014"}

	codecast1.Execute()
	codecast2.Execute()
	codecast3.Execute()
	return true
}
