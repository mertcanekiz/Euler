// Project Euler Problem 45
// http://projecteuler.net/problem=45

package main

import (
	"fmt"
	"math"
	"time"
)

func IsPentagonal(n int) bool {
	r := int(math.Sqrt(float64(24*n)+1)+1)/6
	return r*(3*r-1)/2 == n
}

func Euler() {
	i := 143
	for true {
		i++
		H := i*(2*i-1)
		if IsPentagonal(H) {
			fmt.Println(H)
			return
		}
	}
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
