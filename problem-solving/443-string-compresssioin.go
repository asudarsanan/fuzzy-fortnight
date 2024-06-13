package main

import (
	"fmt"
	"strings"
)

func stringCompression(chars []string) int {
	if len(chars) == 1 {
		return 1
	}
	// sort.Strings(chars)
	var output []string
	j := 0
	for _, char := range chars {
		count := strings.Count(strings.Join(chars, ""), char)
		if len(output) == 0 || (count < 10 && output[j-2] != char) {
			output = append(output, char)
			output = append(output, fmt.Sprint(count))
			j += 2
		} else if len(output) == 0 || (count >= 10 && output[j-2] != char) {
			output = append(output, char)
			output = append(output, fmt.Sprint(count%10))
			j += 2
		}
	}
	// fmt.Print(output)
	return len(output)

}

func sC(chars []string) int {
	// var s string
	pos := 0
	i := 0
	for i < len(chars) {
		letter := chars[i]
		fmt.Print(letter)
		count := 0
		for i < len(chars) && chars[i] == letter {
			count++
			i++

		}
		chars[pos] = letter
		pos++
		if count > 1 {
			if count >= 10 {
				chars[pos] = fmt.Sprint(count / 10)

				pos++
				chars[pos] = fmt.Sprint(count % 10)
				pos++
			}
			chars[pos] = fmt.Sprint(count)
			pos++

		}
	}
	fmt.Print(chars)

	return pos
}
func main() {
	fmt.Print(sC([]string{"a", "b", "b", "b", "b", "a", "b", "b", "b", "b", "b", "b", "b"}))

}
