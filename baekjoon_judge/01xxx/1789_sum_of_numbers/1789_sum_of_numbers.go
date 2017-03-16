package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	K := 1
	for K*(K+1)/2 <= N {
		K++
	}
	fmt.Printf("%d\n", K-1)
}
