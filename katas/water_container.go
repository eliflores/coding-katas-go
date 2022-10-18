package katas

import (
	"math"
)

// Container With Most Water Kata
// From: https://leetcode.com/problems/container-with-most-water/

// maxArea returns the maximum amount of water a container can store.
// You are given an integer array height of length n.
// There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
func maxArea(height []int) int {
	pointLeft := 0
	pointRight := len(height) - 1
	h := int(math.Min(float64(height[pointLeft]), float64(height[pointRight])))
	w := int(math.Abs(float64(pointLeft - pointRight)))
	area := h * w
	highestArea := area

	for i := len(height); (pointRight - pointLeft) > 0; i-- {
		if height[pointLeft] < height[pointRight] {
			pointLeft++
		} else {
			pointRight--
		}

		h = int(math.Min(float64(height[pointLeft]), float64(height[pointRight])))
		w = int(math.Abs(float64(pointLeft - pointRight)))
		area = h * w
		if area > highestArea {
			highestArea = area
		}
	}

	return highestArea
}
