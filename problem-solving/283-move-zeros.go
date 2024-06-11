package main

import "fmt"

func moveZeros(nums []int) []int {
	j := 0
	n := len(nums)
	for i := 0; i <= n-1; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}

	}
	// fmt.Print(j)
	for j <= n-1 {
		nums[j] = 0
		j++
	}
	return nums
}

func main() {
	fmt.Print(moveZeros([]int{1, 0, 2, 0, 3}))
}
