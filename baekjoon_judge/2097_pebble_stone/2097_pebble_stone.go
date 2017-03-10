package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	if N <= 2 {
		fmt.Print("4")
	} else {
		k := 2
		for k*(k+1) < N {
			k++
		}
		if N <= k*k {
			fmt.Print((k*2 - 2) * 2)
		} else {
			fmt.Print((k*2 - 1) * 2)
		}
	}
}

/*
k(k-1) < n <= k^2 : (k*2-2)*2
k^2 < n <= k(k+1) : (k*2-1)*2
*/
