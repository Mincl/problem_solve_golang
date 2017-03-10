package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	for i := N; i >= 1; i-- {
		fmt.Fprintf(writer, "%d\n", i)
	}
	writer.Flush()
}
