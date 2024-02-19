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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuotes(t *testing.T) {

	Convey("TrimQuotes", t, func() {
		So(TrimQuotes(`nope`), ShouldEqual, `nope`)
		So(TrimQuotes(`'nope`), ShouldEqual, `'nope`)
		So(TrimQuotes(`nope"`), ShouldEqual, `nope"`)
		So(TrimQuotes("`nope"), ShouldEqual, "`nope")
		So(TrimQuotes(`'single'`), ShouldEqual, `single`)
		So(TrimQuotes(`"double"`), ShouldEqual, `double`)
		So(TrimQuotes("`backtick`"), ShouldEqual, `backtick`)

		Convey("fancy", func() {
			for _, pair := range FancyQuotes {
				So(TrimQuotes(string(pair.Start)+"fancy"+string(pair.End)), ShouldEqual, `fancy`)
			}
		})
	})

	Convey("IsQuoted", t, func() {
		So(IsQuoted(`“fancy”`), ShouldBeFalse)
		So(IsQuoted(`"normal"`), ShouldBeTrue)
	})

	Convey("IsQuotedFancy", t, func() {
		s, e, ok := IsQuotedFancy(`“fancy”`)
		So(ok, ShouldBeTrue)
		So(s, ShouldEqual, '“')
		So(e, ShouldEqual, '”')
		s, e, ok = IsQuotedFancy(`"fancy"`)
		So(ok, ShouldBeFalse)
		So(s, ShouldEqual, 0)
		So(e, ShouldEqual, 0)
	})

	Convey("IsAnyQuote", t, func() {
		So(IsAnyQuote('"', '!'), ShouldBeFalse)
		So(IsAnyQuote('"'), ShouldBeTrue)
		So(IsAnyQuote('\''), ShouldBeTrue)
		for _, p := range FancyQuotes {
			So(IsAnyQuote(p.Start, p.End), ShouldBeTrue)
		}
	})

}
