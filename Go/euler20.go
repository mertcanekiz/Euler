// Project Euler Problem 20
// http://projecteuler.net/problem=20

package main

import (
	"fmt"
	"math/big"
	"time"
)

func euler() {
	n := big.NewInt(1)
	n.MulRange(1, 100)
	numstr := n.String()
	sum := 0

	for _, char := range numstr {
		sum += int(char - '0')
	}

	fmt.Println(sum)
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
