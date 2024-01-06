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
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSort(t *testing.T) {
	Convey("SplitSortReversed", t, func() {
		So(SplitSortReversed("", ""), ShouldEqual, []string{})
		So(SplitSortReversed("one/two", "/"), ShouldEqual, []string{"two", "one"})
	})

	Convey("SortedByLastName", t, func() {
		So(SortedByLastName([]string{
			"First Name",
			"Another LastName",
			"More KindNames",
		}), ShouldEqual, []string{
			"More KindNames",
			"Another LastName",
			"First Name",
		})
	})

	Convey("SortByLength", t, func() {
		slice := []string{
			"first",
			"third",
			"fourth",
			"second",
			"fifth10",
			"fifth19",
			"fifth01",
		}
		sort.Sort(SortByLength(slice))
		So(slice, ShouldEqual, []string{
			"fifth01",
			"fifth10",
			"fifth19",
			"fourth",
			"second",
			"first",
			"third",
		})
	})
}
