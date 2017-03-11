package main

import (
	"fmt"
	"math"
)

func main() {
	hamburger := make([]int, 3)
	drink := make([]int, 2)
	for i := 0; i < 3; i++ {
		fmt.Scan(&hamburger[i])
	}
	for i := 0; i < 2; i++ {
		fmt.Scan(&drink[i])
	}

	min_h := 1 << 15
	min_d := 1 << 15
	for i := 0; i < 3; i++ {
		min_h = int(math.Min(float64(min_h), float64(hamburger[i])))
	}
	for i := 0; i < 2; i++ {
		min_d = int(math.Min(float64(min_d), float64(drink[i])))
	}
	fmt.Print(min_h + min_d - 50)
}
