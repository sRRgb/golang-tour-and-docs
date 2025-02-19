package types

import (
	"fmt"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSliceAppend(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func wordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, word := range strings.Fields(s) {
		m[word]++
	}

	return m
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4) // closure
}

func adder() func(int) int {
	sum := 0
	return func(x int) int { // closure
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int { // closure
		a, b = b, a+b
		return a
	}
}
