package main

import (
	"fmt"
)

func platable(plots []int, newflowers int) bool {
	for i, plot := range plots {
		if len(plots) == 1 && plot == 0 && newflowers > 0 {
			newflowers = newflowers - 1
		}
		if len(plots) == 2 {
			if plot == 1 {
				break
			} else if plots[0] == 0 && plots[1] == 0 && newflowers > 0 {
				plots[0] = 1
				newflowers = newflowers - 1
			}
		}
		if plot == 0 && i < len(plots)-2 && newflowers > 0 {
			if i == 0 && plots[i+1] == 0 {
				plots[0] = 1
				newflowers = newflowers - 1
			}
			if i > 1 && plots[i-1] == 0 && plots[i+1] == 0 && newflowers > 0 {
				plots[i] = 1
				newflowers = newflowers - 1
			}

		}
		if len(plots) >= 3 && i == len(plots)-1 && plots[len(plots)-2] == 0 && plot == 0 && newflowers > 0 {
			plots[i] = 1
			newflowers = newflowers - 1
		}

	}
	fmt.Print(plots)
	if newflowers == 0 {

		return true
	}
	return false
}

func optimisedPlatable(plots []int, newflowers int) bool {

	count := 0
	for i, plot := range plots {
		if plot == 0 {
			availableLeftPlotExists := (i == 0) || plots[i-1] == 0
			availableRightPlotExists := (i == len(plots)-1) || plots[i+1] == 0

			if availableLeftPlotExists && availableRightPlotExists {
				count = count + 1
			}
		}
	}
	return count >= newflowers
}

// func main() {
// 	fmt.Print(optimisedPlatable([]int{0}, 1))
// }
