// Project Euler Problem 26
// http://projecteuler.net/problem=26

package main

import (
	"fmt"
	"time"
)

func Euler() {
	recurringLength := 0

	for i := 1000; i > 0; i-- {
		if recurringLength >= i {
			break
		}

		remainders := make([]int, i)
		value := 1
		position := 0

		for remainders[value] == 0 && value != 0 {
			remainders[value] = position
			value *= 10
			value %= i
			position++
		}

		if position-remainders[value] > recurringLength {
			recurringLength = position - remainders[value]
		}
	}

	fmt.Println(recurringLength + 1)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
