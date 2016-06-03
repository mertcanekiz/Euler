// Project Euler Problem 40
// http://projecteuler.net/problem=40

package main

import (
	"fmt"
	"strconv"
	"time"
)

func Euler() {
	dn := 1
	i := 1
	length := 0
	prod := 1
	limit := 1000000

	for dn <= limit {
		number := strconv.Itoa(i)
		i++
		l := len(number)
		if length + l >= dn {
			prod *= int(number[dn-length-1]-'0')
			dn *= 10
		}
		length += l
	}

	fmt.Println(prod)
}

func main() {
	t1 := time.Now()
	Euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
