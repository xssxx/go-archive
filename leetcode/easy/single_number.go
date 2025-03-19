package main

import "fmt"

// func count(nums []int) map[int]int {
// 	counts := make(map[int]int)

// 	for _, num := range nums {
// 		counts[num]++
// 	}

// 	return counts
// }

// func singleNumber(nums []int) int {
// 	counted := count(nums)

// 	for num, count := range counted {
// 		if count == 1 {
// 			return num
// 		}
// 	}

// 	return -1
// }

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}

	return res
}

func main() {
	nums := []int{4, 1, 2, 1, 2}

	result := singleNumber(nums)
	fmt.Println(result) // 4
}
