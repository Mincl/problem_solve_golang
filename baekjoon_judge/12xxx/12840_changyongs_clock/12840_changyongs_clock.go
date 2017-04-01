package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var time, h, m, s, N int
	var q, c int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &h)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &m)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &s)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	time = (h*60+m)*60 + s
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &q)
		if q != 3 {
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &c)
		}

		switch q {
		case 1:
			time += c % (24 * 60 * 60)
			time %= 24 * 60 * 60
		case 2:
			time -= c % (24 * 60 * 60)
			if time < 0 {
				time += 24 * 60 * 60
			}
		case 3:
			h = time / 3600
			m = time % 3600 / 60
			s = time % 3600 % 60
			fmt.Fprintf(writer, "%d %d %d\n", h, m, s)
			writer.Flush()
		}
	}
}
