// Project Euler Problem 33
// http://projecteuler.net/problem=33

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

func Euler() {
	// Denominator and nominator products
	dp := 1
	np := 1

	for i := 1; i < 10; i++ {
		for d := 1; d < i; d++ {
			for n := 1; n < d; n++ {
				if (n*10+i)*d == n*(i*10+d) {
					dp *= d
					np *= n
				}
			}
		}
	}

	dp /= gcd(np, dp)

	fmt.Println(dp)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
