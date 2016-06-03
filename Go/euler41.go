// Project Euler Problem 41
// http://projecteuler.net/problem=41

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


func IsPandigital(n int) bool {
	digits, tmp, i := 0, 0, 0

	for n > 0 {
		tmp = digits
		digits = digits | 1 << uint((n%10) -1)
		if tmp == digits {
			return false
		}

		i++
		n /= 10
	}

	return digits == (1 << uint(i)) - 1
}

var primes []int = Sieve(7654321)
func Euler() {
	for i := len(primes)-1; i > 0; i-- {
		if IsPandigital(primes[i]) {
			fmt.Println(primes[i])
			return
		}
	}
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}

