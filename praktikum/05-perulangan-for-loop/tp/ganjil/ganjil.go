package main

import "fmt"

func main() {
	var n, jum int
	fmt.Scan(&n)

	for i := 1; i <= n; i += 2 {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(i)
		jum += i
	}
	fmt.Println()
	fmt.Println(jum)
}
