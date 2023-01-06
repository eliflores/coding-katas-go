package katas

// Longest Substring Without Repeating Characters
// From: https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Given a string s, find the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	max := 0
	n := len(s)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if uniqueLetters(s, i, j) {
				if j-i+1 > max {
					max = j - i + 1
				}
			}
		}
	}
	return max
}

func uniqueLetters(s string, start int, end int) bool {
	var letters = make(map[uint8]bool, len(s))
	for i := start; i <= end; i++ {
		letter := s[i]
		if _, exists := letters[letter]; exists {
			return false
		} else {
			letters[letter] = true
		}
	}
	return true
}
