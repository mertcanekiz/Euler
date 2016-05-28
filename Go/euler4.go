// Project Euler Problem 4
// http://projecteuler.net/problem=4

package main

import (
	"fmt"
	"time"
)

func is_palindrome(n int) bool {
	reversed := 0
	k := n

	for k > 0 {
		reversed = 10*reversed + k%10
		k /= 10
	}

	return n == reversed
}

func euler() {
	result := 0

	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			current := i * j
			if current > result {
				if is_palindrome(current) {
					result = current
				}
			}
		}
	}

	fmt.Println(result)
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
