package main

import "fmt"

func plusOne(digits []int) []int {
	carry := 1
	ptr := len(digits) - 1

	for ptr >= 0 {
		digits[ptr] += carry
		if digits[ptr] == 10 {
			digits[ptr] = 0
			carry = 1
		} else {
			carry = 0
		}
		ptr--
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	digits := []int{9}
	digits = plusOne(digits)
	fmt.Print("[")
	for i, v := range digits {
		if i < len(digits)-1 {
			fmt.Printf("%d,", v)
		} else {
			fmt.Printf("%d", v)
		}
	}
	fmt.Print("]")
}
