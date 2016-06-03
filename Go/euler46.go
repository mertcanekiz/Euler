// Project Euler Problem 46
// http://projecteuler.net/problem=46

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

func IsPrime(n int, primelist []int) bool {
	for _, p := range primelist {
		if p == n {
			return true
		}
	}

	return false
}

func Euler() {
	i := 33
	primes := Sieve(10000)
	squares := []int{}

	for n := 1; n <= 100; n++ {
		squares = append(squares, n*n)
	}

	for true {
		i += 2
		if IsPrime(i, primes) {
			continue
		}

		found := false
		for _, sq := range squares {
			if IsPrime(i - 2*sq, primes) {
				found = true
				break
			}
		}

		if !found {
			fmt.Println(i)
			return
		}
	}
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
