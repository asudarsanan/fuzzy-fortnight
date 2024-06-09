package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	vowels := "AEIOU"
	var stringPicks []string
	var volwelsPos []int
	splitterString := strings.Split(s, "")
	for i, char := range splitterString {
		if strings.Contains(vowels, strings.ToUpper(char)) {
			stringPicks = append(stringPicks, char)
			volwelsPos = append(volwelsPos, i)
		}
	}
	fmt.Print(volwelsPos)
	fmt.Print(stringPicks)

	for i, pos := range volwelsPos {
		splitterString[pos] = stringPicks[len(stringPicks)-1-i]
	}
	fmt.Print(splitterString)
	return strings.Join(splitterString, "")
}

// func main() {
// 	fmt.Print(reverseVowels("leetCODE"))
// }
