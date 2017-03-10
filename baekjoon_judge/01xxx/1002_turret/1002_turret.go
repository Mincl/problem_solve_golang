package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	var x1, y1, r1 int
	var x2, y2, r2 int

	fmt.Scanf("%d\n", &T)

	for i := 0; i < T; i++ {
		fmt.Scanf("%d %d %d %d %d %d\n", &x1, &y1, &r1, &x2, &y2, &r2)
		if x1 == x2 && y1 == y2 {
			if r1 == r2 {
				fmt.Printf("-1\n")
			} else {
				fmt.Printf("0\n")
			}
		} else {
			dist := math.Sqrt(math.Pow(math.Abs(float64(x2-x1)), 2) + math.Pow(math.Abs(float64(y2-y1)), 2))
			if dist > float64(r1+r2) {
				fmt.Printf("0\n")
			} else if dist < float64(r1+r2) {
				max := math.Max(float64(r1), float64(r2))
				min := math.Min(float64(r1), float64(r2))
				if max > dist+min {
					fmt.Printf("0\n")
				} else if max == dist+min {
					fmt.Printf("1\n")
				} else {
					fmt.Printf("2\n")
				}
			} else {
				fmt.Printf("1\n")
			}
		}
	}
}
