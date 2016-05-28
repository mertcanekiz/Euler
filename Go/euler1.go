// Project Euler Problem 1
// http://projecteuler.net/problem=1

package main

import (
	"fmt"
	"time"
)

func euler() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
