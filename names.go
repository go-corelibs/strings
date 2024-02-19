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
	"strings"

	fullname "github.com/amonsat/fullname_parser"
	"github.com/iancoleman/strcase"
	"github.com/weppos/publicsuffix-go/publicsuffix"
)

// FirstName is a wrapper around fullname_parser.ParseFullName
func FirstName(fullName string) (firstName string) {
	// TODO: fork fullname_parser, seems unmaintained
	parsed := fullname.ParseFullname(fullName)
	firstName = parsed.First
	return
}

// LastName is a wrapper around fullname_parser.ParseFullName
func LastName(fullName string) (lastName string) {
	parsed := fullname.ParseFullname(fullName)
	lastName = parsed.Last
	return
}

// ParseDomainName returns the given name split into is component
// parts, in reverse order
func ParseDomainName(input string) (tld, name string, subdomains []string) {
	if parsed, err := publicsuffix.Parse(input); err == nil {
		tld, name = parsed.TLD, parsed.SLD
		if parsed.TRD != "" {
			subdomains = SplitSortReversed(parsed.TRD, ".")
		}
		return
	}
	// fallback to always return something
	list := SplitSortReversed(input, ".")
	if count := len(list); count > 0 {
		if tld = list[0]; count > 1 {
			if name = list[1]; count > 2 {
				subdomains = list[2:]
			}
		}
	}
	return
}

// NameFromEmail returns a user's default name based on just their
// email address, intended to be used as an interesting placeholder
// on a text input field for the user to supply something better
func NameFromEmail(email string) (name string) {
	// split the interesting parts
	before, after, _ := strings.Cut(email, "@")
	// make the name and check the after
	if name = ToSpacedCamel(before); after != "" {
		// suffix the name with a parsed domain
		_, domain, _ := ParseDomainName(after)
		name += " @" + strcase.ToCamel(domain)
	}
	return
}
