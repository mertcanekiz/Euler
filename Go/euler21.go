// Project Euler Problem 21
// http://projecteuler.net/problem=21

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

func Euler() {
	sum := 0
	factors_i := 0
	limit := 10000
	primelimit := int(math.Sqrt(float64(limit)) + 1)
	primelist := Sieve(primelimit)
	cache := make(map[int]int)

	for i := 2; i <= limit; i++ {
		factors_i = FactorSum(i, primelist)
		if factors_i > i {
			cache[i] = factors_i
		} else if factors_i < i {
			val, ok := cache[factors_i]
			if ok {
				if val == i {
					sum += i + factors_i
				}
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
