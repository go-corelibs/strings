// Copyright (c) 2024  The Go-Curses Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strings

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCases(t *testing.T) {
	Convey("PathToSnake", t, func() {
		So(PathToSnake(""), ShouldEqual, "")
		So(PathToSnake("One"), ShouldEqual, "one")
		So(PathToSnake("/One/"), ShouldEqual, "one")
		So(PathToSnake("/One/Two"), ShouldEqual, "one__two")
	})

	Convey("ToTitleWords", t, func() {
		So(ToTitleWords(""), ShouldEqual, "")
		So(ToTitleWords("one"), ShouldEqual, "One")
		So(ToTitleWords("one two"), ShouldEqual, "One Two")
		So(ToTitleWords("one1two"), ShouldEqual, "One1two")
		So(ToTitleWords("one1 two"), ShouldEqual, "One1 Two")
		So(ToTitleWords("one 1two"), ShouldEqual, "One 1two")
	})

	Convey("ToSpaced", t, func() {
		So(ToSpaced("one-two"), ShouldEqual, "one two")
		So(ToSpaced("OneTwo"), ShouldEqual, "one two")
	})

	Convey("ToSpacedTitle", t, func() {
		So(ToSpacedTitle("one-two"), ShouldEqual, "One Two")
		So(ToSpacedTitle("OneTwo"), ShouldEqual, "One Two")
	})

	Convey("ToSpacedCamel", t, func() {
		So(ToSpacedCamel("one-twoTwo"), ShouldEqual, "One Two Two")
		So(ToSpacedCamel("One_TwoTwo"), ShouldEqual, "One Two Two")
	})

	Convey("ToDeepKey", t, func() {
		So(ToDeepKey(".OneThing.another_thing"), ShouldEqual, ".one-thing.another-thing")
		So(ToDeepKey(".one-thing.TwoThings"), ShouldEqual, ".one-thing.two-things")
	})

	Convey("ToDeepVar", t, func() {
		So(ToDeepVar(".OneThing.another_thing"), ShouldEqual, ".OneThing.AnotherThing")
		So(ToDeepVar(".one-thing.TwoThings"), ShouldEqual, ".OneThing.TwoThings")
	})
}
