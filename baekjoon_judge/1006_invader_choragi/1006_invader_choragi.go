package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanWords)
}

Bipartite matching

Dynamic Programming
