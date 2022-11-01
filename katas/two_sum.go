package katas

// Two Sum
// From: https://leetcode.com/problems/two-sum/

// Given an array of integers nums and an integer target.
// twoSum returns indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func twoSum(nums []int, target int) (result []int) {
	var difference int
	numToIdx := make(map[int]int)

	for i, num := range nums {
		difference = target - num
		if idx, found := numToIdx[difference]; found {
			return []int{idx, i}
		}
		numToIdx[nums[i]] = i
	}

	return
}
