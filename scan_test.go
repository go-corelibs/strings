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
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestScan(t *testing.T) {

	Convey("Scan", t, func() {
		checks := []struct {
			src, text string
			b, a      string
			ok        bool
		}{
			{
				"one two more", " ",
				"one", "two more",
				true,
			},
			{
				"one {{ two more }}", "{{",
				"one ", " two more }}",
				true,
			},
			{
				"one { nope }", "}}",
				"one { nope }", "",
				false,
			},
			{
				`"two {{more}}" \}} }} after`, "}}",
				`"two {{more}}" \}} `, ` after`,
				true,
			},
		}

		for _, check := range checks {
			b, a, ok := Scan(check.src, check.text)
			So(b, ShouldEqual, check.b)
			So(a, ShouldEqual, check.a)
			So(ok, ShouldEqual, check.ok)
		}

	})

}

func BenchmarkScan(b *testing.B) {
	for i := 0; i < 1000; i++ {
		end := rand.Intn(gScanTestingParagraphLen)
		src := gScanTestingParagraph[:end]
		_, _, _ = Scan(src, "}}")
	}
}

const (
	gScanTestingParagraph = `
"quoted {{text}}" escaped \}} and actual }}
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`
	gScanTestingParagraphLen = len(gScanTestingParagraph) - 1
)
