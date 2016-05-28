// Project Euler Problem 7
// http://projecteuler.net/problem=7

package main

import (
	"fmt"
	"time"
)

func is_prime(n int) bool {
	if n == 1 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	if n == 5 || n == 7 {
		return true
	}

	if n%3 == 0 {
		return false
	}

	f := 5
	// Every prime number greater than 3 can be written as 6k+1 or 6k-1
	for f*f <= n {
		if n%f == 0 {
			return false
		}
		if n%(f+2) == 0 {
			return false
		}
		f += 6
	}

	return true
}

func euler() {
	numprimes := 1
	counter := 1
	limit := 10001
	for numprimes < limit {
		counter += 2
		if is_prime(counter) {
			numprimes++
		}
	}
	fmt.Println(counter)
}

// Another approach that uses a sieve and has much better execution speed
// than mine, found this on the internet but I'm keeping it since it performs better.
func euler2() {
	const max int = 104745
	var arr [max]int
	counter := 0
	for i := 0; i < max; i++ {
		arr[i] = 1
	}
	for i := 2; i < max; i++ {
		if arr[i] == 1 {
			counter++
			if counter == 10001 {
				fmt.Println(i)
				return
			}
			for j := 2 * i; j < max; j += i {
				arr[j] = 0
			}
		}
	}
}

func main() {
	t1 := time.Now()
	euler2()
	fmt.Println("Elapsed time:", time.Since(t1))
}
