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
	"github.com/go-corelibs/slices"
)

// TrimTmplVar returns the name with leading `$` and `.` characters removed
func TrimTmplVar(name string) (trimmed string) {
	// prepare the trimmed output
	if trimmed = name[:]; len(name) > 0 {
		// while trimmed has a first rune and it is either dollar or dot
		for size := len(trimmed); size > 0 && (trimmed[0] == '$' || trimmed[0] == '.'); {
			// trim the first rune
			trimmed = trimmed[1:]
		}
	}
	return
}

// PruneTmplActions removes all go template action statements
//
// Note that this does not remove the content between if, else and end
// statements. For example:
//
//	`{{ if ... }}stuff{{ else }}moar stuff{{ end }}.`
//
// Turns into:
//
//		`stuff moar stuff.`
//	       ^
//	       |
//
// Also note that PruneTmplActions takes care to not concatenate non-space text
// when pruning surrounding actions though will allow punctuation, similarly to
// AppendWithSpace.
//
// See: https://pkg.go.dev/text/template#hdr-Actions
func PruneTmplActions(value string) (clean string) {
	var stack []string
	length := len(value)
	for idx, r := range value {

		var next uint8
		if idx < length-1 {
			next = value[idx+1]
		}

		// check if currently detecting things
		if current := len(stack) - 1; current > -1 {
			last := len(stack[current]) - 1 // always >= 0

			if r == '}' { // found closing curly brace

				if stack[current][last] == '}' {
					// statement is now closed
					stack, _ = slices.Pop(stack)
					if next > 0 && !IsSpaceOrPunct(next) {
						clean = AddLastSpace(clean)
					}
					continue
				}

			} else if r == '{' && next == '{' {

				// opening within the opening
				stack = append(stack, string(r))
				continue

			} else if last == 0 && stack[current][last] == '{' {
				// looking for another opening curly brace
				if r != '{' {
					// not actually a statement
					if current == 0 {
						// top of stack and not an action, keep as clean
						clean += stack[current] + string(r)
					}
					stack, _ = slices.Pop(stack)
					continue
				}
			}

			stack[current] += string(r)
			continue
		}

		// not within a statement
		if r == '{' {
			// push curly brace onto the detection stack
			stack = append(stack, string(r))
			continue
		}

		// non-statement text is clean
		clean += string(r)
	}
	return
}
