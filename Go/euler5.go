// Project Euler Problem 5
// http://projecteuler.net/problem=5

package main

import (
	"fmt"
	"time"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func euler() {
	result := 1

	for i := 2; i <= 20; i++ {
		result = lcm(i, result)
	}

	fmt.Println(result)
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
