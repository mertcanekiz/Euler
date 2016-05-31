// Project Euler Problem 34
// http://projecteuler.net/problem=34

package main

import (
	"fmt"
	"time"
)

func Factorial(n int) int {
	// I only need 0-9 so I won't bother with the rest
	factorials := []int { 1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880 }
	return factorials[n]
}

func SumDigitFactorials(n int) int {
	number := n
	sum := 0
	for number > 0 {
		d := number % 10
		number /= 10
		sum += Factorial(d)
	}
	return sum
}

func Euler() {
	limit := 40585 // Post-solution limit, was initially (7 * 9!)
	sum := 0

	for i := 10; i <= limit; i++ {
		if i == SumDigitFactorials(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
