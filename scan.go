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

// Scan is a text scanner which looks for unquoted and unescaped `sep`
func Scan(src, sep string) (before, after string, found bool) {

	s := struct {
		quote  rune
		quoted bool
	}{}

	runes := []rune(src)
	total, size := len(runes), len(sep)

	for idx := 0; idx < total; idx++ {

		r := runes[idx]

		if r == '\\' {
			// next character is escaped, skip
			idx += 1
			continue
		} else if IsQuote(r) {
			// this character is a double, single or backtick quotation detected
			if s.quoted {
				// scanning within a quoted string
				if s.quote == r {
					// this character is the ending quotation
					s.quote = 0
					s.quoted = false
				}
				// nothing to do with quoted contents
				continue
			}
			// this character is a starting quotation
			s.quote = r
			s.quoted = true
			continue
		} else if s.quoted {
			// nothing to do with quoted contents
			continue
		} else if remainder := total - idx; size > remainder {
			// early out, not enough characters for sep matching
			break
		} else if found = src[idx:idx+size] == sep; found {
			// sep match found, Scan complete
			before = src[:idx]
			after = src[idx+size:]
			return
		}

	}

	return src, "", false
}
