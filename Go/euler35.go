// Project Euler Problem 35
// http://projecteuler.net/problem=35

// This is kind of brute forcing this problem but I couldn't think
// of a faster way. Still, it runs well within the 1 minute range
// so it's okay I guess ._.

package main

import (
	"fmt"
	"math"
	"time"
)

func Sieve(upperlimit int) []int {
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

func Rotate(n int) []int {
	number := n
	result := []int{}
	result = append(result, n)
	num_digits := int(math.Log10(float64(n)))

	for i := 0; i < num_digits; i++ {
		d := number % 10
		current := d*int(math.Pow(10.0, float64(num_digits))) + (number/10)
		result = append(result, current)
		number = current
	}

	return result
}

func IsPrime(n int, primelist []int) bool {
	for _, p := range primelist {
		if p == n {
			return true
		}
	}

	return false
}

func IsCircularPrime(n int, primelist []int) bool {
	rotations := Rotate(n)
	result := true
	for _, e := range rotations {
		if !IsPrime(e, primelist) {
			result = false
		}
	}

	return result
}

func Euler() {
	primes := Sieve(1000000)
	total := 0
	for _, p := range primes {
		if IsCircularPrime(p, primes) {
			total++
		}
	}
	fmt.Println(total)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
