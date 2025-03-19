package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}) // struct{} = 0 bytes memory usage in value

	for _, num := range nums {
		if _, exists := m[num]; exists {
			return true
		}

		m[num] = struct{}{}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}
