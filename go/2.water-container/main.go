package main

import "fmt"

type TestCase struct {
	Input   []int
	MaxArea int
}

func maxAreaBruteForce(height []int) int {
	maxArea := 0
	for i, a := range height {
		width := 1 // Initialize width to 1 for the current height
		for j := i + width; j < len(height); j++ {
			b := height[j]
			newArea := width * min(a, b)
			if newArea > maxArea {
				maxArea = newArea
			}
			width++ // Increment width for the next iteration
		}
	}
	return maxArea
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		var newArea int
		if height[left] < height[right] {
			newArea = height[left] * (right - left)
			left++
		} else {
			newArea = height[right] * (right - left)
			right--
		}

		if newArea > maxArea {
			maxArea = newArea
		}

	}

	return maxArea
}

func main() {
	// Array of objects with fields `input` and `maxArea`
	testCases := []TestCase{
		{Input: []int{7, 1, 2, 3, 9}, MaxArea: 28},
		{Input: []int{6, 9, 3, 4, 5, 8}, MaxArea: 32},
		{Input: []int{4, 8, 1, 2, 3, 9}, MaxArea: 32},
		{Input: []int{1, 1}, MaxArea: 1},
		{Input: []int{}, MaxArea: 0},
		{Input: []int{1}, MaxArea: 0},
	}

	for _, testCase := range testCases {
		result := maxArea(testCase.Input)
		fmt.Println("Input:", testCase.Input, "Expected MaxArea:", testCase.MaxArea, "Result:", result)
		// Check if the result matches the expected MaxArea
		if result != testCase.MaxArea {
			fmt.Println("Test Case failed")
			fmt.Println()
		} else {
			fmt.Println("Test Case passed")
			fmt.Println()
		}
	}
}
