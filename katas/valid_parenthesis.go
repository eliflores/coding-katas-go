package katas

// Valid Parenthesis
// From: https://leetcode.com/problems/valid-parentheses/

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is length.
// An input string is length if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
func isValid(s string) bool {
	stack := make([]uint8, 0, len(s))
	var last uint8
	var pairs = map[uint8]uint8{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			n := len(stack) - 1
			last = stack[n]
			if s[i] == pairs[last] {
				stack = stack[:n]
			} else {
				stack = append(stack, s[i])
			}
		}
	}

	return len(stack) == 0
}
