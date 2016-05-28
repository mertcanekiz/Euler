package main

import (
	"fmt"
	"math"
	"time"
)

func num_divisors(n int, primelist []int) int {
	divisors := 1
	exponent := 1
	remainder := n

	for _, p := range primelist {
		if p*p > n {
			return divisors * 2
		}

		exponent = 1
		for remainder%p == 0 {
			exponent++
			remainder /= p
		}

		divisors *= exponent

		if remainder == 1 {
			return divisors
		}
	}

	return divisors
}

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
	counter := 1
	number := 0
	primes := sieve(10000)

	for num_divisors(number, primes) < 500 {
		number += counter
		counter++
	}

	fmt.Println(number)
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
