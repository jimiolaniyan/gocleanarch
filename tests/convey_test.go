package tests

import (
	"github.com/jimiolaniyan/gocleanarch/tests/fixtures"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGivenNoCodecastsPresentNoCodecasts(t *testing.T) {
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

func TestPresentViewableCodecastsInChronologicalOrder(t *testing.T) {
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

func TestPresentDownloadableCodecastsInChronologicalOrder(t *testing.T) {
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
				{Title: "A", Picture: "A", Description: "A", Viewable: false, Downloadable: true},
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

func TestShowEpisode(t *testing.T) {
	user := "U"
	permalink := "episode-1"

	Convey("Given Codecasts", t, func() {
		gc := fixtures.GivenCodecast{Title: "A", PublicationDate: "3/1/2014", Permalink: permalink}
		ok := gc.Execute()
		So(ok, ShouldEqual, true)
	})

	Convey("And a user U", t, func() {
		ok := fixtures.CodecastPresentation.AddUser(user)
		So(ok, ShouldEqual, true)
	})

	Convey("With U logged in", t, func() {
		ok := fixtures.CodecastPresentation.LoginUser(user)
		So(ok, ShouldEqual, true)

		Convey("When the user requests details for codecast episode-1", func() {
			details := fixtures.CodecastDetails{}
			ok := details.RequestCodecast(permalink)
			So(ok, ShouldBeTrue)

			Convey("Then the presented title is A, published 3/01/2014", func() {
				title := details.CodecastDetailsTitle()
				date := details.CodecastDetailsDate()

				So(title, ShouldEqual, "A")
				So(date, ShouldEqual, "3/01/2014")
			})

			Convey("With option to purchase viewing license", func() {
				requestViewing := details.CodecastDetailsOfferPurchaseOf("viewing")
				So(requestViewing, ShouldBeTrue)
			})

			Convey("And option to purchase download license", func() {
				downloadViewing := details.CodecastDetailsOfferPurchaseOf("download")
				So(downloadViewing, ShouldBeTrue)
			})
		})

	})
}

func createCodeCasts() bool {
	var codecast1 = fixtures.GivenCodecast{Title: "A", PublicationDate: "3/1/2014"}
	var codecast2 = fixtures.GivenCodecast{Title: "B", PublicationDate: "3/2/2014"}
	var codecast3 = fixtures.GivenCodecast{Title: "C", PublicationDate: "2/18/2014"}

	res1 := codecast1.Execute()
	res2 := codecast2.Execute()
	res3 := codecast3.Execute()

	return res1 == res2 == res3 == true
}
