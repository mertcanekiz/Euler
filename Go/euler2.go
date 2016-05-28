// Project Euler Problem 2
// http://projecteuler.net/problem=2

package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		temp := a + b
		a = b
		b = temp
	}

	return a
}

func euler() {
	counter := 0
	sum := 0
	limit := 4000000

	for true {
		counter++
		current := fib(counter)
		if current > limit {
			break
		}
		if current%2 == 0 {
			sum += current
		}
	}

	fmt.Println(sum)
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
