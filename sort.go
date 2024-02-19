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
	"sort"
	"strings"

	"github.com/maruel/natural"
)

// SplitSortReversed is a wrapper around strings.Split and reverses
// the order of the results
func SplitSortReversed(input, separator string) (split []string) {
	split = strings.Split(input, separator)
	sort.Slice(split, func(i, j int) (less bool) {
		less = i > j
		return
	})
	return
}

// SortedByLastName returns a natual-sorted list of the full
// names given
func SortedByLastName(fullNames []string) (sorted []string) {
	lookup := make(map[int]string)
	for idx, key := range fullNames {
		lookup[idx] = LastName(key)
		sorted = append(sorted, key)
	}
	sort.Slice(sorted, func(i, j int) (less bool) {
		less = natural.Less(lookup[i], lookup[j])
		return less
	})
	return
}

// SortByLength implements the sort.Interface, sorting the slice from longest
// to shortest and sorting equal length strings using natural.Less
type SortByLength []string

func (a SortByLength) Len() int {
	return len(a)
}

func (a SortByLength) Less(i, j int) bool {
	if len(a[i]) > len(a[j]) {
		return true
	}
	return natural.Less(a[i], a[j])
}

func (a SortByLength) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
