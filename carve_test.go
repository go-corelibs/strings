// Copyright (c) 2024  The Go-CoreLibs Authors
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

func TestCarve(t *testing.T) {

	Convey("Carve", t, func() {
		checks := []struct {
			src, start, end string
			b, m, a         string
			ok              bool
		}{
			{
				"one two more", " ", " ",
				"one", "two", "more",
				true,
			},
			{
				"one {{ two more }}", "{{", "}}",
				"one ", " two more ", "",
				true,
			},
			{
				"one { nope }", "{{", "}}",
				"one { nope }", "", "",
				false,
			},
		}

		for _, check := range checks {
			b, m, a, ok := Carve(check.src, check.start, check.end)
			So(b, ShouldEqual, check.b)
			So(m, ShouldEqual, check.m)
			So(a, ShouldEqual, check.a)
			So(ok, ShouldEqual, check.ok)
		}

	})

	Convey("ScanCarve", t, func() {
		checks := []struct {
			src, start, end string
			b, m, a         string
			ok              bool
		}{
			{
				"one two more", " ", " ",
				"one", "two", "more",
				true,
			},
			{
				"one {{ two more }}", "{{", "}}",
				"one ", " two more ", "",
				true,
			},
			{
				"one { nope }", "{{", "}}",
				"one { nope }", "", "",
				false,
			},
			{
				`two {{ "{{more}}" \}} }} after`, "{{", "}}",
				`two `, ` "{{more}}" \}} `, ` after`,
				true,
			},
		}

		for _, check := range checks {
			b, m, a, ok := ScanCarve(check.src, check.start, check.end)
			So(b, ShouldEqual, check.b)
			So(m, ShouldEqual, check.m)
			So(a, ShouldEqual, check.a)
			So(ok, ShouldEqual, check.ok)
		}

	})

	Convey("ScanBothCarve", t, func() {
		checks := []struct {
			src, start, end string
			b, m, a         string
			ok              bool
		}{
			{
				"one two more", " ", " ",
				"one", "two", "more",
				true,
			},
			{
				"one {{ two more }}", "{{", "}}",
				"one ", " two more ", "",
				true,
			},
			{
				"one { nope }", "{{", "}}",
				"one { nope }", "", "",
				false,
			},
			{
				`two {{ "{{more}}" \}} }} after`, "{{", "}}",
				`two `, ` "{{more}}" \}} `, ` after`,
				true,
			},
			{
				`two "{{more}}" {{ \}} }} after`, "{{", "}}",
				`two "{{more}}" `, ` \}} `, ` after`,
				true,
			},
			{
				`two "/* */" /* {{more}} */ after`, "/*", "*/",
				`two "/* */" `, ` {{more}} `, ` after`,
				true,
			},
		}

		for _, check := range checks {
			b, m, a, ok := ScanBothCarve(check.src, check.start, check.end)
			So(b, ShouldEqual, check.b)
			So(m, ShouldEqual, check.m)
			So(a, ShouldEqual, check.a)
			So(ok, ShouldEqual, check.ok)
		}

	})

}

func BenchmarkScanCarve(b *testing.B) {
	for i := 0; i < 1000; i++ {
		end := rand.Intn(gScanCarveTestingParagraphLen)
		src := gScanCarveTestingParagraph[:end]
		_, _, _, _ = ScanCarve(src, "{{", "}}")
	}
}

const (
	gScanCarveTestingParagraph = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
{{ "quoted {{text}}" escaped \}} and actual }}
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`
	gScanCarveTestingParagraphLen = len(gScanCarveTestingParagraph) - 1
)
