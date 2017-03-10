package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Deck struct {
	data []int
}

func (d *Deck) PushFront(data int) {
	d.data = append([]int{data}, d.data...)
}

func (d *Deck) PushBack(data int) {
	d.data = append(d.data, data)
}

func (d *Deck) PopFront() int {
	data := d.Front()
	if data != -1 {
		d.data = d.data[1:]
	}
	return data
}

func (d *Deck) PopBack() int {
	data := d.Back()
	if data != -1 {
		d.data = d.data[:len(d.data)-1]
	}
	return data
}

func (d *Deck) Size() int {
	return len(d.data)
}

func (d *Deck) IsEmpty() int {
	if len(d.data) == 0 {
		return 1
	}
	return 0
}

func (d *Deck) Front() int {
	if len(d.data) == 0 {
		return -1
	}
	return d.data[0]
}

func (d *Deck) Back() int {
	if len(d.data) == 0 {
		return -1
	}
	return d.data[len(d.data)-1]
}

func main() {
	var N, d int
	scanner := bufio.NewScanner(os.Stdin)

	deck := new(Deck)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		parsed := strings.Split(scanner.Text(), " ")
		switch parsed[0] {
		case "push_front":
			fmt.Sscan(parsed[1], &d)
			deck.PushFront(d)
		case "push_back":
			fmt.Sscan(parsed[1], &d)
			deck.PushBack(d)
		case "pop_front":
			fmt.Printf("%d\n", deck.PopFront())
		case "pop_back":
			fmt.Printf("%d\n", deck.PopBack())
		case "size":
			fmt.Printf("%d\n", deck.Size())
		case "empty":
			fmt.Printf("%d\n", deck.IsEmpty())
		case "front":
			fmt.Printf("%d\n", deck.Front())
		case "back":
			fmt.Printf("%d\n", deck.Back())
		}
	}
}
