package main

import "fmt"

func main() {
    var str string
    fmt.Scan(&str)
    for i := 0; i < len(str); i++ {
        if str[i] >= 'a' && str[i] <= 'z' {
            fmt.Printf("%c", str[i]-'a'+'A')
        } else if str[i] >= 'A' && str[i] <= 'Z' {
            fmt.Printf("%c", str[i]-'A'+'a')
        } else {
            fmt.Printf("%c", str[i])
        }
    }
}
