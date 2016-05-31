// Project Euler Problem 30
// http://projecteuler.net/problem=30

package main

import (
	"fmt"
	"time"
)

func SumDigitFifthPower(n int) int {
	number := n
	sum := 0

	for number > 0 {
		d := number % 10
		number /= 10
		sum += d*d*d*d*d
	}

	return sum
}

func Euler() {
	limit := 355000
	result := 0

	for i := 2; i < limit; i++ {
		if SumDigitFifthPower(i) == i {
			result += i
		}
	}

	fmt.Println(result)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
