// Project Euler Problem 15
// http://projecteuler.net/problem=15

package main

import (
	"fmt"
	"time"
)

func choose(n, r int) int {
	result := 1

	for i := 1; i <= r; i++ {
		result *= (n - r + i)
		result /= i
	}

	return result
}

func euler() {
	result := choose(40, 20)
	fmt.Println(result)
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
