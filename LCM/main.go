package main

import (
	"fmt"
)

func main() {
	fmt.Println("working")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan() // use `for scanner.Scan()` to keep reading
	// line := scanner.Text()
	// fmt.Println("captured:", line)
	var k, l int
	fmt.Scanf("%d %d", &k, &l)
	fmt.Println(k, l)
	fmt.Println(Lcm(k, l))
}

func Lcm(a, b int) int {
	greater := 0
	if a >= b {
		greater = a
	} else {
		greater = b

	}
	fmt.Println("greater", greater)
	if greater%a == 0 && greater%b == 0 {
		return greater
	} else {
		found := false
		greater++
		for !found {
			greater++
			if greater%a == 0 && greater%b == 0 {
				found = true
			}
		}
	}
	return greater
}
