package main

import "fmt"

func getDiv(N int) int {
	for j := 9; j >= 2; j-- {
		if N%j == 0 {
			return j
		}
	}
	return -1
}

func main() {
	var T, N int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		st := N
		cnt := 0
		for st > 9 {
			d := getDiv(st)
			if d > 9 || d == -1 {
				break
			} else {
				st /= d
				cnt++
			}
		}
		if st <= 9 {
			fmt.Printf("%d\n", cnt+1)
		} else {
			fmt.Printf("-1\n")
		}
	}
}
