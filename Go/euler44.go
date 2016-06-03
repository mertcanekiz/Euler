// Project Euler Problem 44
// http://projecteuler.net/problem=44

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
	k := 1
	for true {
		k++
		for j := k-1; j > 0; j-- {
			pk := k*(3*k-1)/2
			pj := j*(3*j-1)/2
			D := pk - pj
			S := pk + pj
			if IsPentagonal(D) && IsPentagonal(S) {
				fmt.Println(D)
				return
			}
		}
	}
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}

