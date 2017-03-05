package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var T int
	var A, B int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &T)
	for i := 0; i < T; i++ {
		scanner.Scan()
		parsed := strings.Split(scanner.Text(), ",")
		fmt.Sscan(parsed[0], &A)
		fmt.Sscan(parsed[1], &B)
		fmt.Printf("%d\n", A+B)
	}
}
