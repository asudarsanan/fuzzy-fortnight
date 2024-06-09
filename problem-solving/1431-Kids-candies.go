package main

import (
	"fmt"
	"sort"
)

func candyDistribution(candies []int, extraCandies int) []bool {
	sortedCandies := make([]int, len(candies))
	copy(sortedCandies, candies)
	var outputList []bool
	sort.Ints(sortedCandies)
	gc := sortedCandies[len(sortedCandies)-1]
	fmt.Printf("Sorted Candies %d, %d,%d", sortedCandies, gc, candies)
	for _, candy := range candies {
		if candy+extraCandies >= gc {
			outputList = append(outputList, true)
		} else {
			outputList = append(outputList, false)
		}
	}

	return outputList
}

// func main() {
// 	fmt.Print(candyDistribution([]int{2, 3, 5, 1, 3}, 3))
// }
