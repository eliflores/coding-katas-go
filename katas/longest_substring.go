package katas

// Longest Substring Without Repeating Characters
// From: https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Given a string s, find the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	letters := make(map[uint8]int16, len(s))
	n := len(s)
	left := 0
	max := 0

	for right := 0; right < n; right++ {
		letterRight := s[right]
		letters[letterRight] = letters[letterRight] + 1

		for letters[letterRight] > 1 {
			letterLeft := s[left]
			letters[letterLeft] = letters[letterLeft] - 1
			left++
		}

		if right-left+1 > max {
			max = right - left + 1
		}
	}
	return max
}
