package main

import (
	"testing"
	"unicode/utf8"
)

// func TestReverse(t *testing.T) {
// 	testcases := []struct {
// 		in, want string
// 	}{
// 		{"Hello, world", "dlrow ,olleH"},
// 		{" ", " "},
// 		{"!12345", "54321!"},
// 	}
// 	for _, tc := range testcases {
// 		rev := Reverse(tc.in)
// 		if rev != tc.want {
// 			t.Errorf("Reverse: %q, want %q", rev, tc.want)
// 		}
// 	}
// }


// Fuzzing has a few limitations as well. 
// In your unit test, you could predict the expected output of the Reverse function, and verify that the actual output met those expectations.

// When fuzzing, you can’t predict the expected output, since you don’t have control over the inputs.

// The two properties being checked in this fuzz test are:
// 1) Reversing a string twice preserves the original value
// 2) The reversed string preserves its state as valid UTF-8.
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	// Doesn't go rune by rune
	// f.Fuzz(func(t *testing.T, orig string) {
	// 	rev := Reverse(orig)
	// 	doubleRev := Reverse(rev)
	// 	if orig != doubleRev {
	// 		t.Errorf("Before: %q, after: %q", orig, doubleRev)
	// 	}
	// 	if utf8.ValidString(orig) && !utf8.ValidString(rev) {
	// 		t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
	// 	}
	// })

	f.Fuzz(func(t *testing.T, orig string) {
    rev := Reverse(orig)
    doubleRev := Reverse(rev)
    t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
    if orig != doubleRev {
        t.Errorf("Before: %q, after: %q", orig, doubleRev)
    }
    if utf8.ValidString(orig) && !utf8.ValidString(rev) {
        t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
    }
})
}


