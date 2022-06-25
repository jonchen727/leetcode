package main

import "fmt"

func main() {
	nums := []int{-1, -2, -3, -4, -5}
	target := -8
	m := make(map[int]int)
	// loope over the array and store the sum of each element in the map
	for i := 0; i < len(nums); i++ {
		// find the number needed to sum to target
		goldennum := target - nums[i]
		fmt.Println("map:", m, "golden number:", goldennum)
		// if the number is in the map, return the indices
		if _, ok := m[goldennum]; ok {
			fmt.Println([]int{m[goldennum], i})
		}
		m[nums[i]] = i
	}
}
