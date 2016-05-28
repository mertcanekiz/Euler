// Project Euler Problem 3
// http://projecteuler.net/problem=3

package main

import (
	"fmt"
	"time"
)

func largest_prime_factor(n int) int {
	number := n
	largest := 0
	counter := 2

	for counter*counter <= number {
		if number%counter == 0 {
			number /= counter
			largest = counter
		} else {
			counter++
		}
	}

	if number > largest {
		largest = number
	}

	return largest
}

func euler() {
	number := 600851475143
	fmt.Println(largest_prime_factor(number))
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
