package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for bil := 1; bil <= n; bil++ {
		fmt.Print(bil, " ")
	}
}
