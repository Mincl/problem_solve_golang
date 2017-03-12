package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y, w, h int
    fmt.Scan(&x, &y, &w, &h)
    min := x
    min = int(math.Min(float64(min), float64(y)))
    min = int(math.Min(float64(min), float64(w-x)))
    min = int(math.Min(float64(min), float64(h-y)))
    fmt.Println(min)
}
