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
)

type QuotePair struct {
	Start rune
	End   rune
}

var (
	FancyQuotes = []QuotePair{
		{Start: '“', End: '”'},
		{Start: '‘', End: '’'},
		{Start: '‹', End: '›'},
		{Start: '«', End: '»'},
		{Start: '»', End: '«'},
		{Start: '„', End: '“'},
		{Start: '„', End: '”'},
		{Start: '「', End: '」'},
		{Start: '『', End: '』'},
	}
)

// BookendRunes returns the first and last runes of the string, with `ok`
// being true when there is more than one rune in the input string
func BookendRunes(input string) (start, end rune, ok bool) {
	for idx, r := range input {
		if idx == 0 {
			start = r
		} else if ok = idx > 0; ok {
			end = r
		}
	}
	return
}

// GetFancyQuote returns the FancyQuote pair matching the given rune
func GetFancyQuote[V uint8 | rune](r V) (quote QuotePair, ok bool) {
	v := rune(r)
	for _, p := range FancyQuotes {
		if ok = v == p.Start || v == p.End; ok {
			quote = p
			return
		}
	}
	return
}

// IsQuote returns true if all `runes` given are one of double-quote ("),
// single-quote (') or backtick (`)
func IsQuote[V uint8 | rune](runes ...V) (quote bool) {
	found := 0
	for _, r := range runes {
		switch r {
		case '\'', '`', '"':
			// valid quote detected, trim string
			found += 1
		}
	}
	quote = len(runes) == found
	return
}

// IsFancyQuote returns true if all `runes` given match IsFancyQuote
func IsFancyQuote[V uint8 | rune](runes ...V) (quote bool) {
	found := 0
	for _, r := range runes {
		if _, ok := GetFancyQuote(rune(r)); ok {
			found += 1
		}
	}
	quote = len(runes) == found
	return
}

// IsAnyQuote returns true if all `runes` given match IsQuote or IsFancyQuote
func IsAnyQuote[V uint8 | rune](runes ...V) (quote bool) {
	found := 0
	count := len(runes)
	for _, r := range runes {
		v := rune(r)
		if IsQuote(v) || IsFancyQuote(v) {
			if found += 1; found == count {
				break
			}
		}
	}
	quote = count == found
	return
}

// IsQuoted returns true if the first and last characters in the input are the same and are one of the three main quote
// types: single ('), double (") and literal (`)
func IsQuoted(maybeQuoted string) (quoted bool) {
	if total := len(maybeQuoted); total >= 2 {
		// there's enough length for quotes to be possible
		if start, end, ok := BookendRunes(maybeQuoted); ok && start == end {
			// the first and last characters are the same
			if quoted = IsQuote(start); quoted {
				return
			}
		}
	}
	return
}

// IsQuotedFancy finds the starting and ending runes, returning true if they
// match any of the following pairs of quotations:
func IsQuotedFancy(maybeQuoted string) (start, end rune, quoted bool) {
	if s, e, ok := BookendRunes(maybeQuoted); ok {
		if quoted = IsFancyQuote(s, e); quoted {
			start, end = s, e
		}
	}
	return
}

// TrimQuotes returns the string with the first and last characters trimmed from the string if the string IsQuoted and
// returns the unmodified input string otherwise
func TrimQuotes(maybeQuoted string) (unquoted string) {
	if start, end, ok := BookendRunes(maybeQuoted); ok {
		if start == end && IsQuote(start) {
			unquoted = maybeQuoted[1 : len(maybeQuoted)-1]
			return
		} else if IsFancyQuote(start, end) {
			unquoted = strings.TrimPrefix(maybeQuoted, string(start))
			unquoted = strings.TrimSuffix(unquoted, string(end))
			return
		}
	}
	unquoted = maybeQuoted
	return
}
