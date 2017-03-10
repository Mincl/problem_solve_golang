package main

import "fmt"

func main() {
	var month, day int
	special := 2*32 + 18
	fmt.Scan(&month, &day)
	today := month*32 + day
	if today > special {
		fmt.Print("After")
	} else if today < special {
		fmt.Print("Before")
	} else {
		fmt.Print("Special")
	}
}
