// Project Euler Problem 36
// http://projecteuler.net/problem=36

package main

import (
	"fmt"
	"time"
)

func IsPandigital(n int) bool {
	digits, tmp, i := 0, 0, 0

	for n > 0 {
		tmp = digits
		digits = digits | 1 << uint((n%10) -1)
		if tmp == digits {
			return false
		}

		i++
		n /= 10
	}

	return digits == (1 << uint(i)) - 1
}

func Concat(a, b int) int {
	c := b

	for c > 0 {
		a *= 10
		c /= 10
	}

	return a + b
}

func Euler() {
	result := 0

	for i := 9876; i >= 9123; i-- {
		result = Concat(i, 2*i)
		if IsPandigital(result) {
			break
		}
	}

	fmt.Println(result)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
