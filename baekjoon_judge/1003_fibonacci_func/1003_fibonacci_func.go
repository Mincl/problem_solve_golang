package main

import "fmt"

func main() {
  var T, N int

  fmt.Scanf("%d\n", &T)
  for i := 0 ; i < T ; i++ {
    fmt.Scanf("%d\n", &N)
    one_cnt := make([]int, N+1)
    zero_cnt := make([]int, N+1)

    for j := 0 ; j <= N ; j++ {
      if j == 0 {
        zero_cnt[j]=1
      } else if j == 1 {
        one_cnt[j]=1
      } else {
        one_cnt[j] = one_cnt[j-1] + one_cnt[j-2]
        zero_cnt[j] = zero_cnt[j-1] + zero_cnt[j-2]
      }
    }

    fmt.Printf("%d %d\n", zero_cnt[N], one_cnt[N])
  }
}
