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
	"unicode"
)

// PruneSpaces returns the input string without any space characters
// as defined by unicode.IsSpace
func PruneSpaces(input string) (pruned string) {
	for _, r := range []rune(input) {
		if !unicode.IsSpace(r) {
			pruned += string(r)
		}
	}
	return
}

// CollapseSpaces returns the input string with all consecutive space
// characters collapsed to just one space character (tabs are not considered
// spaces)
func CollapseSpaces(input string) (collapsed string) {
	var skip bool
	for _, r := range input {

		if r == ' ' {
			// found a space
			if skip {
				// already added, keep skipping spaces
				continue
			}
			// collapse any further spaces
			skip = true
		} else if skip {
			// stop collapsing spaces
			skip = false
		}

		collapsed += string(r)
	}
	return
}
