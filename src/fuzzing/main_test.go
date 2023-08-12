package main

import (
	"testing"
	"unicode/utf8"
)

// func TestReverse(t *testing.T) {
// 	testcases := []struct {
// 		input string
// 		want  string
// 	}{
// 		{"Hello World", "dlroW olleH"},
// 		{"!12345", "54321!"},
// 		{"", ""},
// 		{"a", "a"},
// 	}
// 	for _, tc := range testcases {
// 		rev := Reverse(tc.input)
// 		if rev != tc.want {
// 			t.Errorf("Reverse(%q) == %q, want %q", tc.input, rev, tc.want)
// 		}
// 	}
// }

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hellow, World!", "12345", "", "a"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Reverse(%q) == %q, want %q, Doubule Reverse=%q", orig, rev, orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse(%q) == %q, which is invalid UTF-8", orig, rev)
		}
	})
}
