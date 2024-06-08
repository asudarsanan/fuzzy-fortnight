package main

import (
	"fmt"
	"strings"
)

func gcd(str1 string, str2 string) string {
	baseString := str1
	longString := str2
	if len(str1) > len(str2) {
		baseString = str2
		longString = str1
	}
	i := len(baseString)
	for i > 0 {
		subString := baseString[:i]
		fmt.Printf("Substrings: %s \n", subString)
		long_occurance := strings.Count(longString, subString)
		short_occurance := strings.Count(baseString, subString)
		fmt.Printf("Number of occurance: %d, length: %d\n", long_occurance, len(longString))
		if len(longString) == long_occurance*len(subString) && len(baseString) == short_occurance*len(subString) {
			fmt.Printf("Long: %s, Sub: %s, Occ: %d\n", longString, subString, long_occurance)
			return subString
		}
		i--
	}
	return ""
}

// func main() {
// 	op := gcd("ABA", "ABABC")
// 	fmt.Print(op)
// }
