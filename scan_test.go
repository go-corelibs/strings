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
			label     string
			src, text string
			b, a      string
			ok        bool
		}{
			{
				"single space sep",
				"one two more", " ",
				"one", "two more",
				true,
			},
			{
				"opening template curly braces",
				"one {{ two more }}", "{{",
				"one ", " two more }}",
				true,
			},
			{
				"not opening template curly braces",
				"one { nope }", "}}",
				"one { nope }", "",
				false,
			},
			{
				"escaped closing template curly braces",
				`"two {{more}}" \}} }} after`, "}}",
				`"two {{more}}" \}} `, ` after`,
				true,
			},
			{
				"escaped double-quotes, closing template curly braces",
				`"two \"{{more}}\"" \}} }} after`, "}}",
				`"two \"{{more}}\"" \}} `, ` after`,
				true,
			},
			{
				"escape escaped double-quotes, closing template curly braces",
				`"two \\\"{{more}}\\\"" \}} }} after`, "}}",
				`"two \\\"{{more}}\\\"" \}} `, ` after`,
				true,
			},
		}

		for _, check := range checks {
			Convey(check.label, func() {
				b, a, ok := Scan(check.src, check.text)
				So(b, ShouldEqual, check.b)
				So(a, ShouldEqual, check.a)
				So(ok, ShouldEqual, check.ok)
			})
		}

	})

	Convey("ScanQuote", t, func() {
		cases := []struct {
			input   string
			b, q, a string
			found   bool
		}{
			{
				`not a quote`,
				"not a quote", "", "", false,
			},
			{
				`"quoted"`,
				"", "quoted", "", true,
			},
			{
				`before "quoted"`,
				"before ", "quoted", "", true,
			},
			{
				`"quoted" after`,
				"", "quoted", " after", true,
			},
			{
				`before "quoted" after`,
				"before ", "quoted", " after", true,
			},
			{
				`before "quoted \"within\"" after`,
				"before ", `quoted "within"`, " after", true,
			},
			{
				`before "quoted 'within'" after`,
				"before ", `quoted 'within'`, " after", true,
			},
			{
				`before 'quoted "within"' after`,
				"before ", `quoted "within"`, " after", true,
			},
			{
				"before `quoted \"within\"` after",
				"before ", `quoted "within"`, " after", true,
			},
		}

		for _, check := range cases {
			before, quoted, after, found := ScanQuote(check.input)
			So(found, ShouldEqual, check.found)
			So(before, ShouldEqual, check.b)
			So(quoted, ShouldEqual, check.q)
			So(after, ShouldEqual, check.a)
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
