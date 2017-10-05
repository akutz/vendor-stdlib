package sort

import (
	"unicode"
	"unicode/utf8"
)

// StringEqualFoldSlice attaches the methods of Interface to []string,
// sorting in increasing order using a case insensitive comparison.
type StringEqualFoldSlice []string

func (p StringEqualFoldSlice) Len() int { return len(p) }
func (p StringEqualFoldSlice) Less(i, j int) bool {
	return EqualFoldCompare(p[i], p[j]) < 0
}
func (p StringEqualFoldSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p StringEqualFoldSlice) Sort() { Sort(p) }

// StringsEqFold sorts a slice of strings in increasing order using
// a case insensitive comparison.
func StringsEqFold(a []string) { Sort(StringEqualFoldSlice(a)) }

// StringsEqFoldAreSorted tests whether a slice of strings is sorted in
// increasing order using a case insensitive comparison.
func StringsEqFoldAreSorted(a []string) bool {
	return IsSorted(StringEqualFoldSlice(a))
}

// EqualFoldCompare reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding. A return value of 0 means s==t,
// <0 means s<t, and >0 means s>t.
func EqualFoldCompare(s, t string) int {
	for s != "" && t != "" {

		// Extract first rune from each string.
		var sr, tr rune
		if s[0] < utf8.RuneSelf {
			sr, s = rune(s[0]), s[1:]
		} else {
			r, size := utf8.DecodeRuneInString(s)
			sr, s = r, s[size:]
		}
		if t[0] < utf8.RuneSelf {
			tr, t = rune(t[0]), t[1:]
		} else {
			r, size := utf8.DecodeRuneInString(t)
			tr, t = r, t[size:]
		}

		// If they match, keep going; if not, return false.

		// Easy case.
		if tr == sr {
			continue
		}

		// Make sr < tr to simplify what follows.
		result := 1
		if tr < sr {
			result = -result
			tr, sr = sr, tr
		}
		// Fast check for ASCII.
		if tr < utf8.RuneSelf && 'A' <= sr && sr <= 'Z' {
			// ASCII, and sr is upper case.  tr must be lower case.
			srr := sr + 'a' - 'A'
			if tr == srr {
				continue
			}
			if tr < srr {
				return result
			}
			if tr > srr {
				return -result
			}
		}

		// General case. SimpleFold(x) returns the next equivalent rune > x
		// or wraps around to smaller values.
		r := unicode.SimpleFold(sr)
		for r != sr && r < tr {
			r = unicode.SimpleFold(r)
		}
		if r == tr {
			continue
		}
		if tr < r {
			return result
		}
		if tr > r {
			return -result
		}
	}

	// One string is empty. Are both?
	if s == "" && t != "" {
		return -1
	}
	if s != "" && t == "" {
		return 1
	}
	return 0
}
