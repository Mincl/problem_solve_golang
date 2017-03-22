package main

import (
	"fmt"
)

func check(cnt []int) bool {
	for i := 0; i < 10; i++ {
		if cnt[i] > 1 {
			return false
		}
	}
	return true
}

func main() {
	var result int
	fmt.Scan(&result)
	cnt := make([]int, 10)
	for hello := 10000; hello <= 99999; hello++ {
		if hello/100%10 != hello/10%10 {
			continue
		}
		for i := 0; i < 10; i++ {
			cnt[i] = 0
		}
		cnt[hello/10000]++
		cnt[hello/1000%10]++
		cnt[hello/100%10]++
		cnt[hello%10]++
		if !check(cnt) {
			continue
		}
		world := result - hello
		if world/100000 >= 1 || world < 10000 ||
			world < 0 || world/1000%10 != hello%10 ||
			world/10%10 != hello/10%10 {
			continue
		}
		cnt[world/10000]++
		cnt[world/100%10]++
		cnt[world%10]++
		if !check(cnt) {
			continue
		}
		fmt.Printf("%7d\n", hello)
		fmt.Printf("+%6d\n", world)
		fmt.Printf("-------\n")
		fmt.Printf("%7d", result)
		return
	}
	fmt.Printf("No Answer")
}
