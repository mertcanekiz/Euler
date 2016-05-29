// Project Euler Problem 16
// http://projecteuler.net/problem=16

package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func euler() {
	result := big.NewInt(1)
	two := big.NewInt(2)

	for i := 0; i < 1000; i++ {
		result.Mul(result, two)
	}

	numstr := result.String()
	sum := 0

	for _, char := range numstr {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		sum += digit
	}

	fmt.Println(sum)
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
