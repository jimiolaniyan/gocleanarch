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

func TestAddUserSpec(t *testing.T) {
	Convey("Given a new user U", t, func() {
		user := "User"
		Convey("When the U is added", func() {
			status := codecastPresentation.AddUser(user)
			Convey("Then U should be in the system", func() {
				So(status, ShouldEqual, true)
			})
		})
	})
}

func TestLoginUserSpec(t *testing.T) {
	Convey("Given a user U", t, func() {
		user := "U"
		codecastPresentation.AddUser(user)
		Convey("When the U logs in", func() {
			res := codecastPresentation.LoginUser(user)
			Convey("Then U should have a valid session", func() {
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
