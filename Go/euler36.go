// Project Euler Problem 36
// http://projecteuler.net/problem=36

package main

import (
	"fmt"
	"time"
)

func IsPalindrome(number, base int) bool {
	reversed := 0
	n := number

	for n > 0 {
		reversed = base*reversed + n%base
		n /= base
	}

	return number == reversed
}

func Euler() {
	limit := 1000000
	sum := 0

	for i := 1; i < limit; i++ {
		if IsPalindrome(i, 10) && IsPalindrome(i, 2) {
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
