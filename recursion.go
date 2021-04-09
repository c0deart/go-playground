// https://gobyexample.com/recursion

package main

import "fmt"

func DoFactorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * DoFactorial(n-1)
}

func main1() {
	fmt.Printf("Doing fact for %d", 5)
	fmt.Printf("Result: %d ", DoFactorial(4))
}
