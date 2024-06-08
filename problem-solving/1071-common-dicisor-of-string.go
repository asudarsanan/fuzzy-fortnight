package main

import (
	"fmt"
	"strings"
)

func stringdivisor(str1 string, str2 string) string {
	subString1 := computeSubStrings(str1)
	subStrin2 := computeSubStrings(str2)

	if subString1 == subStrin2 && (subStrin2 != " " && subStrin2 != " ") {
		return subStrin2
	} else {
		return " "
	}
}

func computeSubStrings(str string) string {
	length := len(str)
	var subString string
	// splitedString := strings.Split(str,"")
	if length > 1 {
		i := 1
		for i <= length/2 {
			subString = str[:i]
			fmt.Printf("subString: %s\n", subString)
			fmt.Printf("Rest of the string; %s \n", str[i:i+len(subString)])
			if subString == str[i:i+len(subString)] {
				fmt.Printf("First Identifieds substing %s\n", subString)
				occurances := strings.Count(str, subString)
				fmt.Printf("Number of occurances %d\n", occurances)
				if length == occurances*(len(subString)) {
					return subString
				}

			}
			i++

		}
	} else {
		return str
	}
	return " "
}
func main() {
	// _ = computeSubStrings("ABABAB")
	output := stringdivisor("ABABABA", "ABABABABABABA")
	fmt.Print(output)

}
