package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(digits []int) []int {
	if len(digits) < 2 {
		return digits
	}
	if len(digits) == 2 {
		if digits[0] > digits[1] {
			digits[0], digits[1] = digits[1], digits[0]
		}
		return digits
	}
	base := digits[0]
	digits = digits[1:]
	var left, right []int
	for _, elem := range digits {
		if elem <= base {
			left = append(left, elem)
		} else {
			right = append(right, elem)
		}
	}
	left = append(left, base)
	return append(quickSort(left), quickSort(right)...)
}

func main() {
	var digits []int

	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 20; i++ {
		digits = append(digits, int(rand.Int31()%100))
	}
	fmt.Println(quickSort(digits))
}
