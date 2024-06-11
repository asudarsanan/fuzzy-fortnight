package main

import (
	"fmt"
)

func multiplyList(nums []int) []int {
	answers := make([]int, len(nums))

	for j := 0; j <= len(nums)-1; j++ {
		answer := 1
		for i := 0; i <= len(nums)-1; i++ {
			if i != j {
				answer *= nums[i]
			}
			if i == j {
				continue
			}
		}
		fmt.Print(answer)
		answers[j] = answer
	}
	return answers
}

// func main() {
// 	fmt.Print(multiplyList([]int{1, 2, 3, 4}))
// 	// 24,12,8,6
// }
