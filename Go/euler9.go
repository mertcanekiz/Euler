// Project Euler Problem 9
// http://projecteuler.net/problem=9

package main

import (
	"fmt"
	"time"
)

func euler() {
	target := 1000
	a_limit := target / 3
	b_limit := target / 2
	for a := 1; a < a_limit; a++ {
		for b := a; b < b_limit; b++ {
			c := 1000 - b - a
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
