package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	splittedString := strings.Split(s, " ")
	var output []string
	for i := len(splittedString) - 1; i >= 0; i-- {
		if len(splittedString[i]) != 0 {
			output = append(output, splittedString[i])
		}
	}
	return strings.Join(output, " ")
}
func main() {
	fmt.Print(reverseWords(" Life Is    Beauiful "))

}
