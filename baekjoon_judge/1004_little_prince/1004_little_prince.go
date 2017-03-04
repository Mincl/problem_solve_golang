package main

import (
  "fmt"
  "math"
)

func main() {
  var T, N int
  var st_x, st_y, ed_x, ed_y int
  var c_x, c_y, r int

  fmt.Scanf("%d\n", &T)
  for i := 0; i < T ; i++ {
    cnt := 0
    fmt.Scanf("%d %d %d %d\n", &st_x, &st_y, &ed_x, &ed_y)
    fmt.Scanf("%d\n", &N)
    for j := 0 ; j < N ; j++ {
      fmt.Scanf("%d %d %d\n", &c_x, &c_y, &r)

      st_dist := math.Sqrt(math.Pow(float64(c_x - st_x), 2) + math.Pow(float64(c_y - st_y), 2))
      ed_dist := math.Sqrt(math.Pow(float64(c_x - ed_x), 2) + math.Pow(float64(c_y - ed_y), 2))

      if (st_dist < float64(r) || ed_dist < float64(r)) && !(st_dist < float64(r) && ed_dist < float64(r)) {
        cnt++
      }
    }
    fmt.Printf("%d\n", cnt)
  }
}
