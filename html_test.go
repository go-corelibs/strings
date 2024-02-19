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

func TestHtml(t *testing.T) {
	Convey("EscapeHtmlAttribute", t, func() {
		So(EscapeHtmlAttribute(``), ShouldEqual, ``)
		So(EscapeHtmlAttribute(`'this and that'`), ShouldEqual, `this and that`)
		So(EscapeHtmlAttribute(`"this & that"`), ShouldEqual, `this & that`)
		So(EscapeHtmlAttribute(`this "that"`), ShouldEqual, `this &quot;that&quot;`)
	})
}
