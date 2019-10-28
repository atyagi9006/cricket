package random

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testBuilder func(bool) func()

func TestGenerator(t *testing.T) {

	testcases := map[string]testBuilder{

		"For a empty list of probability should get a err": func(withAssertions bool) func() {
			return func() {

				prob := make([]int, 0)
				expected := fmt.Errorf("Invalid probability : %v", prob)
				randomNums, err := Generator(prob)
				if withAssertions {
					So(err, ShouldNotBeNil)
					So(randomNums, ShouldEqual, -1)
					So(err.Error(), ShouldEqual, expected.Error())
				}
			}
		},
		"For a nil list of probability should get a err": func(withAssertions bool) func() {
			return func() {
				var prob []int
				expected := fmt.Errorf("Invalid probability : %v", prob)
				randomNums, err := Generator(prob)
				if withAssertions {
					So(err, ShouldNotBeNil)
					So(randomNums, ShouldEqual, -1)
					So(err.Error(), ShouldEqual, expected.Error())
				}
			}
		},
		"For single element in probablity list same sholud be returned": func(withAssertions bool) func() {
			return func() {

				prob := []int{5}
				randomNum, err := Generator(prob)
				if withAssertions {
					So(err, ShouldBeNil)
					So(randomNum, ShouldEqual, len(prob))

				}
			}
		},
		"Given that  for a list of probabilities get a random number within a length": func(withAssertions bool) func() {
			return func() {
				prob := []int{5, 2, 50, 10, 13, 11, 9}
				randomNums, err := Generator(prob)
				if withAssertions {
					So(err, ShouldBeNil)
				}
				if withAssertions {
					So(randomNums, ShouldBeLessThanOrEqualTo, len(prob))
				}
			}
		},
	}
	Convey("Ramdon Generator Tests. ", t, func() {
		for testname, builder := range testcases {
			Convey(testname, builder(true))
		}
	})

}
