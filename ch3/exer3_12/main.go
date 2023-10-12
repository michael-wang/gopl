package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func anagrams(s, t string) bool {
	// Diff only in case are NOT anagrams, so we can do this in lower cases.
	s = strings.ToLower(s)
	t = strings.ToLower(t)

	scount := count(s)
	tcount := count(t)
	fmt.Printf("s: %s, scount: %v\n", s, scount)
	fmt.Printf("t: %s, tcount: %v\n", t, tcount)

	schars := maps.Keys(scount)
	tchars := maps.Keys(tcount)
	// Sort so we can compare with ordering ignored (like set).
	slices.Sort(schars)
	slices.Sort(tchars)
	if slices.Compare(schars, tchars) != 0 {
		// Diff letter set
		return false
	}

	for char, spos := range scount {
		tpos := tcount[char]
		if len(spos) != len(tpos) {
			// Diff number of same letter
			return false
		}
	}

	// Now s & t has same letter set, same number of each letter.
	// If we can find any letter with diff positions, it's anagram.
	for char, spos := range scount {
		tpos := tcount[char]
		if slices.Compare(spos, tpos) != 0 {
			return true
		}
	}

	// All letters has same positions, so s and t can only diff in cases, which is NOT anagram
	return false
}

//	If s = "foobar", schar = {
//		'f': [0],
//		'o': [1,2],
//		'b': [3],
//		'a': [4],
//		'r': [5],
//	}
//
// Notice: case ignored
func count(s string) map[byte][]int {
	chars := map[byte][]int{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if chars[c] == nil {
			chars[c] = []int{i}
		} else {
			chars[c] = append(chars[c], i)
		}
	}
	return chars
}
