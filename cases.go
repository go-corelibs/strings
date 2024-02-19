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
	"unicode"

	"github.com/iancoleman/strcase"
)

// PathToSnake trims any leading and trailing slashes and converts the
// string to a snake_cased__directory_separated format where the
// directory separator is two underscores
func PathToSnake(path string) (snake string) {
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	path = strings.ReplaceAll(path, "/", "--")
	snake = strcase.ToSnake(path)
	return
}

// ToTitleWords title-cases all words in the given text
func ToTitleWords(text string) (capitalized string) {
	// inspired by: https://stackoverflow.com/a/70284562
	var within bool
	characters := make([]rune, len(text))
	for idx, character := range []rune(text) {
		if isLetter := unicode.IsLetter(character); isLetter && !within {
			characters[idx] = unicode.ToTitle(character)
			within = true
		} else if isLetter {
			characters[idx] = character
		} else if within = unicode.IsNumber(character); within {
			characters[idx] = character
		} else {
			characters[idx] = character
			within = false
		}
	}
	capitalized = string(characters)
	return
}

// ToSpaced is a wrapper around strcase.ToDelimited with a space delimiter
func ToSpaced(text string) (spaced string) {
	spaced = strcase.ToDelimited(text, ' ')
	return
}

// ToSpacedTitle is a wrapper around ToSpaced and ToTitleWords
func ToSpacedTitle(text string) (spacedTitle string) {
	spaced := ToSpaced(text)
	spacedTitle = ToTitleWords(spaced)
	return
}

// ToSpacedCamel is a wrapper around ToSpaced and strcase.ToCamel for each
// word. The difference between ToSpacedTitle and ToSpacedCamel is in what
// is considered a capital letter. ToSpacedTitle uses unicode.ToTitle to
// figure that out while ToSpacedCamel uses strcase.ToCamel for that
func ToSpacedCamel(text string) (spacedCamel string) {
	spaced := ToSpaced(text)
	fields := strings.Fields(spaced)
	for idx := 0; idx < len(fields); idx++ {
		fields[idx] = strcase.ToCamel(fields[idx])
	}
	spacedCamel = strings.Join(fields, " ")
	return
}

// ToDeepKey converts go template variables like `.ThisThing.Variable` to a
// map key used in like `.this-thing.variable`
func ToDeepKey(text string) (deepKey string) {
	parts := strings.Split(strings.TrimPrefix(text, "."), ".")
	for _, part := range parts {
		deepKey += "." + strcase.ToKebab(part)
	}
	return
}

// ToDeepVar is the opposite of ToDeepKey, translating `.this-thing.variable`
// to `.ThisThing.Variable` format
func ToDeepVar(text string) (deepVar string) {
	parts := strings.Split(strings.TrimPrefix(text, "."), ".")
	for _, part := range parts {
		deepVar += "." + strcase.ToCamel(part)
	}
	return
}
