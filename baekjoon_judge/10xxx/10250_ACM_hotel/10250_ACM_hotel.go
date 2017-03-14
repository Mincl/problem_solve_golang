package main

import "fmt"

func main() {
	var T, H, W, N int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&H, &W, &N)
		floor := N % H
		num := N / H
		if floor == 0 {
			floor = H
		} else {
			num += 1
		}
		fmt.Printf("%d%02d\n", floor, num)
	}
}
