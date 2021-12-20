package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	var sumArr []int64

	for i := range arr {
		var tempArr []int32
		var sum int64
		for j, v := range arr {
			if i == j {
				continue
			}
			tempArr = append(tempArr, v)
		}

		for _, v := range tempArr {
			sum += int64(v)
		}
		sumArr = append(sumArr, sum)
	}
	var max, min = sumArr[0], sumArr[0]
	for _, v := range sumArr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Println(fmt.Sprintf("%d %d", min, max))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
