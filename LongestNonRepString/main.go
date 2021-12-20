package main

import (
	"fmt"
	"strings"
)

func main() {

	valStr := "abcbae"
	arr := longSubStr(valStr)
	fmt.Println(arr)
	maxLeng(arr)
}

func maxLeng(strArr []string) {
	var longSubStr []string
	maxLeng := 0
	for _, v := range strArr {

		if len(v) >= maxLeng {
			maxLeng = len(v)
			longSubStr = append(longSubStr, v)
		}
	}
	fmt.Println("long substring", longSubStr)
}

func longSubStr(valStr string) []string {
	var strArr []string
	for i := 0; i < len(valStr); i++ {

		tempStr := valStr[i:]
		fmt.Println("tempStr", tempStr)
		isUnique := true
		for j := 0; j < len(tempStr); j++ {
			if strings.Count(tempStr, string(tempStr[j])) >= 2 {
				isUnique = false
				break
			}
		}
		if isUnique {
			strArr = append(strArr, tempStr)
		}
	}
	return strArr
}
