package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	bil := 1
	for {
		fmt.Print(bil, " ")

		if bil == n {
			break
		}
		bil++
	}
}
