// Project Euler Problem 23
// http://projecteuler.net/problem=23

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

func FactorSum(number int, primelist []int) int {
	n := number
	sum := 1
	p := primelist[0]
	i, j := 0, 0

	for p*p <= n && n > 1 && i < len(primelist) {
		p = primelist[i]
		i++
		if n%p == 0 {
			j = p * p
			n /= p
			for n%p == 0 {
				j = j * p
				n /= p
			}
			sum *= (j - 1) / (p - 1)
		}
	}

	if n > 1 {
		sum *= n + 1
	}

	return sum - number
}

func Contains(slice []int, elem int) bool {
	for _, a := range slice {
		if a == elem {
			return true
		}
	}

	return false
}

func Euler() {
	const limit int = 28123
	primelist := Sieve(int(math.Sqrt(float64(limit))) + 1)
	abundants := []int{}
	possible := make([]bool, limit+1)
	sum := 0

	for i := 2; i <= limit; i++ {
		if FactorSum(i, primelist) > i {
			abundants = append(abundants, i)
		}
	}

	for i := 0; i < limit+1; i++ {
		possible[i] = false
	}

	for _, i := range abundants {
		for _, j := range abundants {
			if i+j <= limit {
				possible[i+j] = true
			} else {
				break
			}
		}
	}

	for i := 1; i <= limit; i++ {
		if !possible[i] {
			sum += i
		}
	}

	fmt.Println(sum)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
