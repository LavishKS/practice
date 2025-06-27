package main

import "fmt"

func main() {
	input := []int{1, 3, 4, 2}
	var target int = 6
	var seen map[int]int = make(map[int]int)

	// Using a map to store the indices of the numbers
	for i, num := range input {
		var numberToFind = target - num
		if index, found := seen[numberToFind]; found {
			fmt.Printf("Indices found: %d and %d\n", index, i)
			return
		}
		seen[num] = i
	}
}
