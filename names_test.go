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

func TestNames(t *testing.T) {
	Convey("FirstName, LastName", t, func() {
		So(FirstName("First M Last"), ShouldEqual, "First")
		So(LastName("First Middle Last"), ShouldEqual, "Last")
	})

	Convey("ParseDomainName", t, func() {
		tld, name, subs := ParseDomainName("ja.go-enjin.org")
		So(tld, ShouldEqual, "org")
		So(name, ShouldEqual, "go-enjin")
		So(subs, ShouldEqual, []string{"ja"})

		tld, name, subs = ParseDomainName(".go-enjin.org")
		So(tld, ShouldEqual, "org")
		So(name, ShouldEqual, "go-enjin")
		So(subs, ShouldEqual, []string{""})
	})

	Convey("NameFromEmail", t, func() {
		So(NameFromEmail("name@addr.ess"), ShouldEqual, "Name @Addr")
	})
}
