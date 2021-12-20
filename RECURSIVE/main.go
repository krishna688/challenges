/*
sum of first 10 number without using loop
*/

package main

import "fmt"

func main() {
	fmt.Println(add(10))
}

func add(n int) int {
	if n == 1 {
		return 1
	}
	return n + add(n-1)
}
