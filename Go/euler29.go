// Project Euler Problem 29
// http://projecteuler.net/problem=29

package main

import (
	"fmt"
	"math/big"
	"time"
)

func Euler() {
	results := make(map[string]bool)

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			current := big.NewInt(0)
			current.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			results[current.String()] = true
		}
	}

	fmt.Println(len(results))
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
