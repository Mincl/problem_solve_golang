package main

import "fmt"
import "math"

func main() {
	var L, A, B, C, D int
	fmt.Scanf("%d\n", &L)
	fmt.Scanf("%d\n", &A)
	fmt.Scanf("%d\n", &B)
	fmt.Scanf("%d\n", &C)
	fmt.Scanf("%d", &D)

	language_date := math.Ceil(float64(A) / float64(C))
	math_date := math.Ceil(float64(B) / float64(D))

	work_date := math.Max(language_date, math_date)

	fmt.Printf("%d", int(math.Max(float64(L)-work_date, 0)))
}
