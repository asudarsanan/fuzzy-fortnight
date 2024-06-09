package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	splittedString := strings.Split(s, " ")
	orginalString := make([]string, len(splittedString))
	copy(orginalString, splittedString)
	for i := 0; i < len(splittedString); i++ {
		splittedString[i] = orginalString[len(splittedString)-1-i]
	}
	fmt.Print(splittedString)
	return strings.Join(splittedString, " ")
}
func main() {
	fmt.Print(reverseWords("Life Is Beauiful"))

}
