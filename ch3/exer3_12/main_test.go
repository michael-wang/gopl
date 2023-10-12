package main

import "testing"

func TestComma(t *testing.T) {
	var cc = []struct {
		s   string
		t   string
		exp bool
	}{
		// Possitive cases: same letters with diff order.
		{"abc", "acb", true},
		{"abc", "bac", true},
		{"abc", "bca", true},
		{"abc", "cab", true},
		{"abc", "cba", true},
		// Diff order AND capitalization
		{"abc", "aCb", true},
		{"abc", "acB", true},
		{"abc", "ACB", true},
		// // Different words
		{"a", "b", false},
		{"abc", "def", false},
		// // Same words
		{"a", "a", false},
		{"abc", "abc", false},
		// // Same letters but with diff numbers
		{"anagram", "nnagram", false},
		{"anaagram", "anaggram", false},
		{"anagrama", "anagramm", false},
		// Diff only in capitalization
		{"aNaGram", "anagram", false},
		{"anagram", "anAgraAM", false},
	}

	for _, c := range cc {
		got := anagrams(c.s, c.t)
		if c.exp != got {
			t.Errorf("\ns:\t\t%s\nt:\t\t%s\nexpected:\t%t\ngot:\t\t%t\n", c.s, c.t, c.exp, got)
		}
	}
}
