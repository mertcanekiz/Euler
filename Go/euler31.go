// Project Euler Problem 31
// http://projecteuler.net/problem=31

package main

import (
	"fmt"
	"time"
)

func BruteForce() {
	limit := 200
	result := 0

	for a := limit; a >= 0; a -= 200 {
		for b := a; b >= 0; b -= 100 {
			for c := b; c >= 0; c -= 50 {
				for d := c; d >= 0; d -= 20 {
					for e := d; e >= 0; e -= 10 {
						for f := e; f >= 0; f -= 5 {
							for g := f; g >= 0; g -= 2 {
								result++
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(result)
}

// Taken from www.mathblog.dk
func Dynamic() {
	limit := 200
	coinSizes := []int{ 1, 2, 5, 10, 20, 50, 100, 200 }
	ways := make([]int, limit+1)
	ways[0] = 1

	for i := 0; i < len(coinSizes); i++ {
		for j := coinSizes[i]; j <= limit; j++ {
			ways[j] += ways[j - coinSizes[i]]
		}
	}

	fmt.Println(ways[len(ways)-1])
}

func main() {
	t1 := time.Now()
	BruteForce()
	fmt.Println("Elapsed time:", time.Since(t1))

	t1 = time.Now()
	Dynamic()
	fmt.Println("Elapsed time:", time.Since(t1))
}
