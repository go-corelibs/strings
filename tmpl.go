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
