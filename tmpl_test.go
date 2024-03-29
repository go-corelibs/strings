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

func TestTmpl(t *testing.T) {
	Convey("TrimTmplVar", t, func() {
		So(TrimTmplVar(""), ShouldEqual, "")
		So(TrimTmplVar("."), ShouldEqual, "")
		So(TrimTmplVar("$"), ShouldEqual, "")
		So(TrimTmplVar("Name"), ShouldEqual, "Name")
		So(TrimTmplVar(".Name"), ShouldEqual, "Name")
		So(TrimTmplVar("$Name"), ShouldEqual, "Name")
		So(TrimTmplVar("$.Name"), ShouldEqual, "Name")
	})

	Convey("PruneTmpTags", t, func() {
		So(PruneTmplActions(`{{ this }}`), ShouldEqual, "")
		So(PruneTmplActions(`{ this }`), ShouldEqual, "{ this }")
		So(PruneTmplActions(`{{ if }}stuff{{ else }}moar stuff{{ end }}.`), ShouldEqual, "stuff moar stuff.")
		So(PruneTmplActions(`This is multi-line text.
This line has a single-line {{ "statement" -}}.
This line has a multi-line {{
  "statement"
-}}.
This line has a statement with valid curly braces within: {{
	"{{ curly braces! }}"
-}}.
`), ShouldEqual, `This is multi-line text.
This line has a single-line .
This line has a multi-line .
This line has a statement with valid curly braces within: .
`)
	})

}
