package main

import (
	"strings"
)

func merger(str1 string, str2 string) string {
	var mergedList []string
	split1 := strings.Split(str1, "")
	split2 := strings.Split(str2, "")

	lstr1 := len(str1)
	lstr2 := len(str2)
	todo := lstr1
	if lstr1 > lstr2 {
		todo = lstr2
	}
	i := 0
	for i < todo {
		mergedList = append(mergedList, split1[i])
		mergedList = append(mergedList, split2[i])
		i++
	}

	if lstr1 < lstr2 {
		slice := split2[todo:]
		mergedList = append(mergedList, slice...)
	} else if lstr1 > lstr2 {
		slice := split1[todo:]
		mergedList = append(mergedList, slice...)
	}
	// fmt.Print(mergedList)
	return strings.Join(mergedList, "")
}
func main() {
	output := merger("asssdraaIndian", "aihuasnnthankappan") //adbecfdef
	print(output)
}
