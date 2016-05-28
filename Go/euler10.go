// Project Euler Problem 10
// http://projecteuler.net/problem=10

package main

import (
	"fmt"
	"math"
	"time"
)

func sieve(upperlimit int) []int {
	sievelimit := (upperlimit - 1) / 2
	uppersqrt := int(math.Sqrt(float64(upperlimit))-1) / 2
	primebits := make([]bool, sievelimit+1)

	for i := 0; i < sievelimit+1; i++ {
		primebits[i] = true
	}

	for i := 1; i <= uppersqrt; i++ {
		if primebits[i] {
			for j := i * 2 * (i + 1); j <= sievelimit; j += 2*i + 1 {
				primebits[j] = false
			}
		}
	}

	numbers := []int{}
	numbers = append(numbers, 2)

	for i := 1; i <= sievelimit; i++ {
		if primebits[i] {
			numbers = append(numbers, 2*i+1)
		}
	}

	return numbers
}

func euler() {
	limit := 2000000
	primes := sieve(limit)
	sum := 0

	for _, prime := range primes {
		sum += prime
	}

	fmt.Println(sum)
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
