package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	root *TreeNode
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) Insert(val int) {
	if b.root == nil {
		b.root = &TreeNode{Val: val}
		return
	}

	curr := b.root
	for {
		if val < curr.Val {
			// Go left
			if curr.Left == nil {
				curr.Left = &TreeNode{Val: val}
				return
			}
			curr = curr.Left
		} else {
			// Go right
			if curr.Right == nil {
				curr.Right = &TreeNode{Val: val}
				return
			}
			curr = curr.Right
		}
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}

	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])

	return node
}

func levelOrder(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}

	var res []interface{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			res = append(res, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			res = append(res, nil)
		}
	}

	for len(res) > 0 && res[len(res)-1] == nil {
		res = res[:len(res)-1]
	}

	return res
}

func printResult(result []interface{}) {
	fmt.Print("[")
	for i, v := range result {
		if i < len(result)-1 {
			fmt.Printf("%v, ", v)
		} else {
			fmt.Printf("%v", v)
		}
	}
	fmt.Print("]\n")
}

func main() {
	ans := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	res := levelOrder(ans)
	printResult(res)
}
