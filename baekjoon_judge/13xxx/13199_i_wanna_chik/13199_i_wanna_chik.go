package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	var P, M, F, C int
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &P, &M, &F, &C)

		eatCnt := M / P
		couponCnt := eatCnt * C
		eatCntNoCoupon := eatCnt + couponCnt/F

		eatCntHasCoupon := eatCnt
		// fmt.Println(eatCntHasCoupon)
		for couponCnt/F > 0 {
			eatCntHasCoupon += couponCnt / F
			couponCnt = couponCnt/F*C + couponCnt%F
			// fmt.Println(eatCntHasCoupon, couponCnt)
		}

		fmt.Fprintf(writer, "%d\n", eatCntHasCoupon-eatCntNoCoupon)
		writer.Flush()
	}
}
