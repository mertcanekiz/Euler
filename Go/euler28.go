// Project Euler Problem 28
// http://projecteuler.net/problem=28

package main

import (
	"fmt"
	"time"
)

func SumCorners(n int) int {
	return (2*n-1)*(2*n-1) + (2*n-1)*(2*n-1) - 2*(n-1) + (2*n-1)*(2*n-1) - 2*2*(n-1) + (2*n-1)*(2*n-1) - 2*3*(n-1)
}

func Euler() {
	sum := 1

	for i := 2; i < 502; i++ {
		sum += SumCorners(i)
	}

	fmt.Println(sum)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
