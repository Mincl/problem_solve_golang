package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "sort"
)

type Axis struct {
    x int
    y int
}

type ByX []Axis
type ByY []Axis

func (bx ByX) Len() int { return len(bx) }
func (bx ByX) Less(i, j int) bool {
    if bx[i].x == bx[j].x {
        return bx[i].y < bx[j].y
    }
    return bx[i].x < bx[j].x
}
func (bx ByX) Swap(i, j int) { bx[i], bx[j] = bx[j], bx[i] }
func (by ByY) Len() int      { return len(by) }
func (by ByY) Less(i, j int) bool {
    if by[i].y == by[j].y {
        return by[i].x < by[j].x
    }
    return by[i].y < by[j].y
}
func (by ByY) Swap(i, j int) { by[i], by[j] = by[j], by[i] }

func distance(a Axis, b Axis) int {
    return int(math.Pow(float64(a.x-b.x), 2)) + int(math.Pow(float64(a.y-b.y), 2))
}

func divideConquer(axis []Axis, N, left, right int) int {
    min := 1 << 31
    if left > right {
        return min
    }
    middle := (left + right) / 2
    if 0 <= middle-1 {
        min = int(math.Min(float64(min), float64(distance(axis[middle], axis[middle-1]))))
    }
    if N > middle+1 {
        min = int(math.Min(float64(min), float64(distance(axis[middle], axis[middle+1]))))
    }
    if 0 <= middle-1 && N > middle+1 {
        min = int(math.Min(float64(min), float64(distance(axis[middle-1], axis[middle+1]))))
    }
    min = int(math.Min(float64(min), float64(divideConquer(axis, N, left, middle-1))))
    min = int(math.Min(float64(min), float64(divideConquer(axis, N, middle+1, right))))
    return min
}

func main() {
    var N int
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    fmt.Sscan(scanner.Text(), &N)
    axis := make([]Axis, N)
    for i := 0; i < N; i++ {
        scanner.Scan()
        fmt.Sscan(scanner.Text(), &axis[i].x, &axis[i].y)
    }
    sort.Sort(ByX(axis))
    minX := divideConquer(axis, N, 0, N-1)
    sort.Sort(ByY(axis))
    minY := divideConquer(axis, N, 0, N-1)
    fmt.Printf("%d\n", int(math.Min(float64(minX), float64(minY))))
}
