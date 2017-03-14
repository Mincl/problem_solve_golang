package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fibonacci := make([]int, N+1)
	fibonacci[1] = 1
	for i := 2; i <= N; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	fmt.Print(fibonacci[N])
}
