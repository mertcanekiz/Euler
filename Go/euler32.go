// Project Euler Problem 32
// http://projecteuler.net/problem=32

package main

import (
	"fmt"
	"time"
)

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

func Concat(a, b int) int {
	c := b

	for c > 0 {
		a *= 10
		c /= 10
	}

	return a + b
}

func Contains(slice []int, elem int) bool {
	for _, e := range slice {
		if elem == e {
			return true
		}
	}

	return false
}

func Euler() {
	products := []int{}
	sum := 0
	prod, current := 0, 0

	for m := 2; m < 100; m++ {
		nbegin := 0
		nend := 10000 / m + 1
		if m > 9 {
			nbegin = 123
		} else {
			nbegin = 1234
		}

		for n := nbegin; n < nend; n++ {
			prod = m * n
			current = Concat(Concat(prod, n), m)
			if current > 100000000 && current < 1000000000 && IsPandigital(current){
				if !Contains(products, prod) {
					products = append(products, prod)
				}
			}
		}
	}

	for _, p := range products {
		sum += p
	}

	fmt.Println(sum)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
