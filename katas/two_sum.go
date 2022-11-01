package katas

// Two Sum
// From: https://leetcode.com/problems/two-sum/

// Given an array of integers nums and an integer target.
// twoSum returns indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func twoSum(nums []int, target int) []int {
	// Based on the constraint that 2 <= nums.length
	if len(nums) < 2 {
		return []int{}
	}

	// Based on the constraint that there is always exactly one solution
	if len(nums) == 2 {
		return []int{0, 1}
	}

	var difference int
	numToIdx := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		difference = target - nums[i]
		if idx, found := numToIdx[difference]; found {
			return []int{idx, i}
		}
		numToIdx[nums[i]] = i
	}

	// It is assumed that there will always be solution, but if there isn't, we return an empty array.
	return []int{}
}
