package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(d interface{}) {
	*h = append(*h, d.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	d := old[n-1]
	*h = old[0 : n-1]
	return d
}

func main() {
	var N int
	var d int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	h := &Heap{}
	heap.Init(h)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &d)

		if d == 0 {
			if h.Len() == 0 {
				fmt.Printf("0\n")
			} else {
				fmt.Printf("%d\n", heap.Pop(h))
			}
		} else {
			heap.Push(h, d)
		}
	}
}
