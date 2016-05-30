// Project Euler Problem 18
// http://projecteuler.net/problem=18

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func euler() {
	text := strings.Split(ReadFile("euler18.txt"), "\n")

	length := len(text)
	triangle := make([]int, length*length)

	for i := 0; i < length*length; i++ {
		triangle[i] = 0
	}

	for row, line := range text {
		numbers := strings.Split(line[:len(line)-1], " ")
		for col, num := range numbers {
			currentnum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			triangle[col+row*length] = currentnum
		}
	}

	for y := length - 2; y >= 0; y-- {
		for x := 0; x < length-1; x++ {
			triangle[x+y*length] += int(math.Max(float64(triangle[x+(y+1)*length]), float64(triangle[x+1+(y+1)*length])))
		}
	}

	fmt.Println(triangle[0])
}

func main() {
	t1 := time.Now()
	euler()
	fmt.Println("Elapsed time:", time.Since(t1))
}
