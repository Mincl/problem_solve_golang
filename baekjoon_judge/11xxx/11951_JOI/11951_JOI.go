package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Person struct {
	x   int64
	dir int64
}

// return: index
func binarySearch(array []int64, val int64, t bool) int {
	idx := -1
	diff := int64(1 << 60)
	left := 0
	right := len(array) - 1
	for left <= right {
		middle := (right + left) / 2
		if t {
			if array[middle] < val {
				left = middle + 1
			} else {
				if diff > array[middle]-val {
					diff = array[middle] - val
					idx = middle
				}
				right = middle - 1
			}
		} else {
			if array[middle] > val {
				right = middle - 1
			} else {
				if diff > val-array[middle] {
					diff = val - array[middle]
					idx = middle
				}
				left = middle + 1
			}
		}
	}
	return idx
}

func main() {
	var N, Q, h int
	var T int64
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N, &T, &Q)
	person := make([]Person, N)
	place := make([]int64, 0)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &person[i].x, &person[i].dir)
		if person[i].dir == 2 {
			person[i].dir = -1
		}
		// -> <-
		if i != 0 && person[i-1].dir == 1 && person[i].dir == -1 {
			place = append(place, (person[i].x+person[i-1].x)/2)
		}
	}
	for i := 0; i < Q; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &h)
		t := true
		if person[h-1].dir < 0 {
			t = false
		}
		pos := binarySearch(place, person[h-1].x, t)
		if pos == -1 || int64(math.Abs(float64(place[pos]-person[h-1].x))) > T {
			fmt.Printf("%d\n", person[h-1].x+T*person[h-1].dir)
		} else {
			fmt.Println(place[pos])
		}
	}
}
