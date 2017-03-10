package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var first, second, third string
	key := [10]string{
		"black", "brown", "red", "orange", "yellow",
		"green", "blue", "violet", "grey", "white",
	}
	val := make(map[string]int)
	mul := make(map[string]int)
	for i := 0; i < 10; i++ {
		val[key[i]] = i
		mul[key[i]] = int(math.Pow(10, float64(i)))
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &first)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &second)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &third)
	fmt.Print((val[first]*10 + val[second]) * mul[third])
}
