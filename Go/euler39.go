// Project Euler Problem 39
// http://projecteuler.net/problem=39

package main

import (
	"fmt"
	"time"
)
func NumPythagoreanTriples(limit int) int {
	result := 0

	for a := 1; a <= limit / 3; a++ {
		for b := a; b <= limit / 2; b++ {
			c := limit - a - b
			if a*a + b*b == c*c {
				result++
			}
		}
	}

	return result
}

func Euler() {
	result := 0
	maxSolutions := 0

	for i := 1000; i > 0; i-- {
		if i < maxSolutions {
			break
		}

		current := NumPythagoreanTriples(i)
		if current > maxSolutions {
			maxSolutions = current
			result = i
		}
	}

	fmt.Println(result)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
