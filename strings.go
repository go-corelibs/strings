// Copyright (c) 2023  The Go-Curses Authors
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
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"

	"github.com/go-corelibs/slices"
)

// ToKebabs converts all the given strings to kebab-case
func ToKebabs(inputs ...string) (out []string) {
	for _, i := range inputs {
		out = append(out, strcase.ToKebab(i))
	}
	return
}

// ToLowers converts all the given strings to lower case
func ToLowers(in ...string) (out []string) {
	for _, i := range in {
		out = append(out, strings.ToLower(i))
	}
	return
}

// GetBasicMime removes the semicolon and any trailing bits from the given
// mime string
func GetBasicMime(mime string) (basic string) {
	basic, _, _ = strings.Cut(mime, ";")
	basic = strings.TrimSpace(basic)
	return
}

// QuoteJsonValue is intended to be used when marshalling json content and is
// applied to only the json values as strings
func QuoteJsonValue(value string) (out string) {
	if value = strings.TrimSpace(value); value == "" {
		return `""` // empty
	} else if lower := strings.ToLower(value); lower == "true" || lower == "false" {
		return lower // bool
	} else if i, ee := strconv.Atoi(value); ee == nil {
		return strconv.Itoa(i) // clean int
	} else if _, eee := strconv.ParseFloat(value, 64); eee == nil {
		return strings.TrimRight(value, "0")
	}
	// not any of the other types, just return quoted
	out = fmt.Sprintf(`%q`, value)
	return
}

// IsTrue returns true if the text given is a truthy word or any positive number
//
// Truthy words:
//
//	"true", "t", "yes", "y" and "on"
func IsTrue(text string) bool {
	switch strings.ToLower(text) {
	case "true", "yes", "on", "1", "t", "y":
		return true
	}
	if v, err := strconv.Atoi(text); err == nil {
		return v > 0
	} else if f, err := strconv.ParseFloat(text, 64); err == nil {
		return f > 0.0
	}
	return false
}

// IsFalse returns true if the text given is a (case-insensitive) falsey word
// or a number that is less than or equal to zero.
//
// Falsey words is:
//
//	"false", "f", "no", "n" and "off"
func IsFalse(text string) bool {
	switch strings.ToLower(text) {
	case "false", "no", "off", "0", "f", "n", "":
		return true
	}
	if v, err := strconv.Atoi(text); err == nil {
		return v <= 0
	} else if f, err := strconv.ParseFloat(text, 64); err == nil {
		return f <= 0.0
	}
	return false
}

// IsQuoted returns true if the first and last characters in the input are the same and are one of the three main quote
// types: single ('), double (") and literal (`)
func IsQuoted(maybeQuoted string) (quoted bool) {
	if total := len(maybeQuoted); total > 2 {
		// there's enough length for quotes to be possible
		if last := total - 1; maybeQuoted[0] == maybeQuoted[last] {
			// the first and last characters are the same
			switch maybeQuoted[0] {
			case '\'', '`', '"':
				// valid quote detected, trim string
				quoted = true
				return
			}
		}
	}
	return
}

// TrimQuotes returns the string with the first and last characters trimmed from the string if the string IsQuoted and
// returns the unmodified input string otherwise
func TrimQuotes(maybeQuoted string) (unquoted string) {
	if IsQuoted(maybeQuoted) {
		unquoted = maybeQuoted[1 : len(maybeQuoted)-1]
		return
	}
	unquoted = maybeQuoted
	return
}

// UniqueFromSpaceSep splits the given value on spaces and only appends it to
// the original slice if not already present, returning the updated results
func UniqueFromSpaceSep(value string, original []string) (updated []string) {
	updated = original[:]
	lookup := make(map[string]struct{})
	for _, v := range updated {
		lookup[v] = struct{}{}
	}
	for _, part := range strings.Split(value, " ") {
		if part == "" {
			continue
		} else if _, present := lookup[part]; !present {
			updated = append(updated, part)
		}
	}
	return
}

// Empty returns true if the string has a length of zero or is all spaces
func Empty(value string) (empty bool) {
	empty = len(strings.TrimSpace(value)) == 0
	return
}

// PruneTmplTags removes all go template statements, things wrapped in doubled
// curly braces like: `{{ stuff }}`
func PruneTmplTags(value string) (clean string) {
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
						// this is the top of the stack
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
		} else {
			// non-statement text is clean
			clean += string(r)
		}

	}
	return
}

// AppendWithSpace appends add to src with a space, ensuring that only one space
// is separating the two given strings. If the add string starts with
// punctuation, no space is used.
//
// AppendWithSpace is intended to be used within the Go-Enjin page format
// features that need to join text with html constructs. The `.njn` block
// format for example joins strings and json objects describing html elements
// into the rendered output.
//
// Enjin Block Example:
//
//	{
//	   "type": "p",
//	   "text": [
//	     "This sentence ends with a",
//	     { "type": "a", "href": "https://go-enjin.org", "text": "link" },
//	     "."
//	   ]
//	}
//
// Is rendered without a space between the "link" and the "." but does have
// a space after the first sentence text.
func AppendWithSpace(src, add string) (combined string) {
	combined = src
	if add == "" {
		return
	} else if src == "" {
		return add
	} else if last := len(src) - 1; last >= 0 {
		switch {
		case unicode.IsPunct(rune(add[0])):
		case unicode.IsSpace(rune(add[0])):
		case unicode.IsSpace(rune(src[last])):
		default:
			combined += " "
		}
	}
	combined += add
	return
}

// TrimPrefixes trims the first path prefix matching value, used to prune known
// things from arbitrary path strings which may or may not be prefixed with any
// of the prefixes given
func TrimPrefixes(value string, prefixes ...string) (trimmed string) {
	trimmed = value
	for _, prefix := range prefixes {
		if len(trimmed) > 0 && trimmed[0] == '/' {
			trimmed = trimmed[1:]
		}
		if check := strings.TrimPrefix(trimmed, prefix); trimmed != check {
			// stop at the first trim
			trimmed = check
			if trimmed != "" && trimmed[0] == '/' {
				trimmed = trimmed[1:]
			}
			return
		}
	}
	return
}
