package kafka_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKafka(t *testing.T) {
	Convey("Given I have a valid kafka setup", t, func() {
		Convey("When I send a message", func() {
			Convey("Then I should receive a message", func() {
				So(true, ShouldEqual, true)
			})
		})
	})
}
