package main

import "fmt"

func main() {
	var n, lebar int
	sempit := 900

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&lebar)
		if lebar < sempit {
			sempit = lebar
		}
	}
}
