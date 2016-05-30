// Project Euler Problem 27
// http://projecteuler.net/problem=27

package main

import (
	"fmt"
	"math"
	"time"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	if n == 5 || n == 7 {
		return true
	}

	if n%3 == 0 {
		return false
	}

	f := 5
	for f*f <= n {
		if n%f == 0 {
			return false
		}
		if n%(f+2) == 0 {
			return false
		}
		f += 6
	}

	return true
}

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

func Euler() {
	maxLength := 0
	result := 0
	primelist := Sieve(1000)
	n := 0

	for a := -999; a < 1000; a++ {
		for _, b := range primelist {
			n = 0

			for IsPrime((a+n)*n + b) {
				n++
			}

			if n > maxLength {
				maxLength = n
				result = a * b
			}
		}
	}

	fmt.Println(result)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
