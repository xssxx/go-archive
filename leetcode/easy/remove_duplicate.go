package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	slow := 1

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

func main() {
	sl := []int{0, 0, 1, 1, 1, 2, 2}
	l := removeDuplicates(sl)
	fmt.Println(l)
	fmt.Println(sl)
}
