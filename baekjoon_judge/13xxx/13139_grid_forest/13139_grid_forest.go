package main

import "fmt"

type Queue struct {
	data []Axis
}

type Direction struct {
	x int
	y int
}

type Axis struct {
	xy  Direction
	cnt int
}

func (q *Queue) Push(a Axis) {
	q.data = append(q.data, a)
}

func (q *Queue) Pop() Axis {
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *Queue) Size() int { return len(q.data) }

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nearestPrimes(n int) (int, int, int, int) {
	i := n + 1
	for !isPrime(i) {
		i++
	}
	j := n - 1
	for !isPrime(j) {
		j--
	}
	k := i + 1
	l := j - 1
	for !isPrime(k) {
		k++
	}
	for !isPrime(l) {
		l--
	}
	return l, j, i, k
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if y > x {
		x, y = y, x
	}

	dir := [4]Direction{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	q := new(Queue)
	history := make(map[Direction]bool)
	q.Push(Axis{Direction{x, y}, 0})
	history[Direction{x, y}] = true
	minXPrime, lxPrime, rxPrime, maxXPrime := nearestPrimes(x)
	minYPrime, _, _, maxYPrime := nearestPrimes(y)
	cnt := 1 << 31
	b := false
	for q.Size() > 0 {
		cur := q.Pop()
		if lxPrime == cur.xy.x || rxPrime == cur.xy.x {
			// fmt.Printf("(%d, %d) [%d]\n", cur.xy.x, cur.xy.y, cur.cnt+cur.xy.x+cur.xy.y)
			if cnt > cur.cnt+cur.xy.x+cur.xy.y {
				cnt = cur.cnt + cur.xy.x + cur.xy.y
				b = true
			}
		}
		for i := 0; i < 4; i++ {
			d := Direction{cur.xy.x + dir[i].x, cur.xy.y + dir[i].y}
			if _, ok := history[d]; !ok && 0 < d.x && 0 < d.y && gcd(d.x, d.y) == 1 &&
				minXPrime <= d.x && d.x <= maxXPrime && minYPrime <= d.y && d.y <= maxYPrime {
				q.Push(Axis{d, cur.cnt + 1})
				history[d] = true
			}
		}
	}
	if b {
		fmt.Printf("%d\n", cnt)
	} else {
		fmt.Printf("-1\n")
	}
}
