package strings

import (
	"strings"
)

// Carve finds the `start` and `end` markers in `src` and carves out the
// "before carve", "middle of start/end range" and "after carve" segments. If
// the `start` and `end` range is not `found` then the `before` will contain
// the entire `src` input
//
// Carve is based on two calls to [strings.Cut], one for the `start` and again
// for the `end`
func Carve(src, start, end string) (before, middle, after string, found bool) {
	var b0, a0, b1, a1 string
	if b0, a0, found = strings.Cut(src, start); found {
		if b1, a1, found = strings.Cut(a0, end); found {
			before = b0
			middle = b1
			after = a1
			return
		}
	}
	before = src
	return
}

// ScanCarve is like Carve except that ScanCarve uses Scan instead of [strings.Cut]
// for finding the `end` separator
func ScanCarve(src, start, end string) (before, middle, after string, found bool) {
	var b0, a0, b1, a1 string
	if b0, a0, found = strings.Cut(src, start); found {
		if b1, a1, found = Scan(a0, end); found {
			before = b0
			middle = b1
			after = a1
			return
		}
	}
	before = src
	return
}

// ScanBothCarve is like ScanCarve except that ScanBothCarve uses Scan to find
// both the `start` and `end` separators
func ScanBothCarve(src, start, end string) (before, middle, after string, found bool) {
	var b0, a0, b1, a1 string
	if b0, a0, found = Scan(src, start); found {
		if b1, a1, found = Scan(a0, end); found {
			before = b0
			middle = b1
			after = a1
			return
		}
	}
	before = src
	return
}
