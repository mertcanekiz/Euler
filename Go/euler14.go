package main

import (
	"fmt"
	"time"
)

var results map[int]int = make(map[int]int)

func collatz_length(n int) int {
	next := n
	count := 0

	for next != 1 {
		if next%2 == 0 {
			next /= 2
		} else {
			next = next*3 + 1
		}

		count++

		result, ok := results[next]
		if ok {
			results[n] = result + count
			return result + count
		}
	}

	results[n] = count + 1
	return count + 1
}

func euler() {
	maxChain := 0
	result := 0
	limit := 1000000

	for i := 2; i < limit; i++ {
		current := collatz_length(i)
		if current > maxChain {
			maxChain = current
			result = i
		}
	}

	fmt.Println(result)
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
