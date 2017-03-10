package main

import (
	"bytes"
	"fmt"
)

func main() {
	var K int
	var res bytes.Buffer
	i := 0
	fmt.Scanf("%d\n", &K)
	for i < K {
		res.WriteString("1")
		i++
	}
	for i > 1 {
		res.WriteString("0")
		i--
	}
	fmt.Printf("%s", res.String())
}
