package main

import "fmt"

func generate(numRows int) [][]int {
	ans := [][]int{
		{1},
	}

	for i := 1; i < numRows; i++ {
		new := []int{1}
		for j := 1; j < i; j++ {
			n := ans[i-1][j-1] + ans[i-1][j]
			new = append(new, n)
		}
		new = append(new, 1)
		ans = append(ans, new)
	}

	return ans
}

func main() {
	res := generate(5)
	fmt.Println(res)
}
