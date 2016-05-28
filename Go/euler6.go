// Project Euler Problem 6
// http://projecteuler.net/problem=6

package main

import (
	"fmt"
	"time"
)

func euler() {
	n := 100
	result := n * (n + 1) * (n - 1) * (3*n + 2) / 12
	fmt.Println(result)
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
