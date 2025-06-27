package main

import (
	"fmt"
	"math"
)

func maxRainWaterBruteForce(input []int) int {
	result := 0
	for i, h := range input {
		maxLeft, maxRight := 0, 0
		for j := i - 1; j > -1; j-- {
			maxLeft = max(maxLeft, input[j])
		}

		for j := i + 1; j < len(input); j++ {
			maxRight = max(maxRight, input[j])
		}

		wh := min(maxLeft, maxRight) - h
		if wh > 0 {
			result += wh
		}
	}

	return result

}

func maxRainWaterCachedHeights(input []int) int {
	n := len(input)
	left, right := make([]int, n), make([]int, n)
	leftMax, rightMax := 0, 0
	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		left[i] = leftMax
		leftMax = max(leftMax, input[i])

		right[j] = rightMax
		rightMax = max(rightMax, input[j])
	}
	result := 0
	for i, h := range input {
		wh := min(left[i], right[i]) - h
		if wh > 0 {
			result += wh
		}
	}
	return result
}

func maxRainWaterTwoPointer(input []int) int {
	// This function uses the two-pointer approach to calculate the maximum rainwater trapped.
	// The idea is to maintain two pointers, one starting from the left and the other from the right.
	// At each step, compare the heights at the two pointers and move the pointer with the smaller height inward.
	// This ensures that the trapped water is calculated efficiently without precomputing left and right maximums.

	n := len(input)
	leftMax := math.MinInt
	left := 0
	rightMax := math.MinInt
	right := n - 1
	result := 0
	for left < right {
		leftHeight := input[left]
		rightHeight := input[right]
		if leftHeight < rightHeight {
			if leftHeight < leftMax {
				result += leftMax - leftHeight
			} else {
				leftMax = leftHeight
			}
			left++
		} else {
			if rightHeight < rightMax {
				result += rightMax - rightHeight
			} else {
				rightMax = rightHeight
			}
			right--
		}
	}

	return result
}

func maxRainWater(input []int) int {
	return maxRainWaterTwoPointer(input)
}

type TestCase struct {
	Input    []int
	Expected int
}

func main() {
	testcases := []TestCase{
		{Input: []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}, Expected: 8},
		{Input: []int{5, 0, 3, 0, 0, 0, 2, 3, 4, 2, 1}, Expected: 20},
		{Input: []int{}, Expected: 0},
		{Input: []int{3}, Expected: 0},
		{Input: []int{3, 4, 3}, Expected: 0},
	}

	for _, testcase := range testcases {
		result := maxRainWater(testcase.Input)

		fmt.Println("Input:", testcase.Input, "Expected MaxRainWater:", testcase.Expected, "Result:", result)
		if result != testcase.Expected {
			fmt.Println("Test Case failed")
			fmt.Println()
		} else {
			fmt.Println("Test Case passed")
			fmt.Println()
		}
	}
}
