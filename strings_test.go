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

func TestStrings(t *testing.T) {
	Convey("ToKebabs, ToLowers", t, func() {
		So(ToKebabs("OneTwo", "ManyMore"), ShouldEqual, []string{"one-two", "many-more"})
		So(ToLowers("One", "Two"), ShouldEqual, []string{"one", "two"})
	})

	Convey("GetBasicMime", t, func() {
		So(GetBasicMime("text/plain"), ShouldEqual, "text/plain")
		So(GetBasicMime("text/plain; charset=utf8"), ShouldEqual, "text/plain")
	})

	Convey("QuoteJsonValue", t, func() {
		So(QuoteJsonValue(""), ShouldEqual, `""`)
		So(QuoteJsonValue("True"), ShouldEqual, `true`)
		So(QuoteJsonValue("this"), ShouldEqual, `"this"`)
		So(QuoteJsonValue("10"), ShouldEqual, `10`)
		So(QuoteJsonValue("1,000"), ShouldEqual, `"1,000"`)
		So(QuoteJsonValue("1.10000100"), ShouldEqual, `1.100001`)
	})

	Convey("IsTrue, IsFalse", t, func() {
		So(IsTrue(""), ShouldEqual, false)
		So(IsTrue("0"), ShouldEqual, false)
		So(IsTrue("1"), ShouldEqual, true)
		So(IsTrue("0.0"), ShouldEqual, false)
		So(IsTrue("0.1"), ShouldEqual, true)
		for _, b := range []string{"true", "yes", "on", "1", "t", "y"} {
			So(IsTrue(b), ShouldEqual, true)
		}
		So(IsFalse(""), ShouldEqual, true)
		So(IsFalse("0"), ShouldEqual, true)
		So(IsFalse("1"), ShouldEqual, false)
		So(IsFalse("0.0"), ShouldEqual, true)
		So(IsFalse("0.1"), ShouldEqual, false)
		for _, b := range []string{"false", "no", "off", "0", "f", "n"} {
			So(IsFalse(b), ShouldEqual, true)
		}
		So(IsFalse("a"), ShouldEqual, false)
	})

	Convey("UniqueFromSpaceSep", t, func() {
		So(UniqueFromSpaceSep("one  two", []string{"a", "one"}), ShouldEqual, []string{
			"a", "one", "two",
		})
	})

	Convey("Empty", t, func() {
		So(Empty(""), ShouldEqual, true)
		So(Empty("  "), ShouldEqual, true)
		So(Empty(" ! "), ShouldEqual, false)
	})

	Convey("AppendWithSpace", t, func() {
		So(AppendWithSpace("", "a"), ShouldEqual, "a")
		So(AppendWithSpace("a", ""), ShouldEqual, "a")
		So(AppendWithSpace("a", "b"), ShouldEqual, "a b")
		So(AppendWithSpace("a ", "b"), ShouldEqual, "a b")
		So(AppendWithSpace("a", ". b"), ShouldEqual, "a. b")
	})

	Convey("TrimPrefixes", t, func() {
		So(TrimPrefixes("/one/two", "nope", "one"), ShouldEqual, "two")
		So(TrimPrefixes("/one/two"), ShouldEqual, "/one/two")
		So(TrimPrefixes("/one", "one"), ShouldEqual, "")
	})

}
