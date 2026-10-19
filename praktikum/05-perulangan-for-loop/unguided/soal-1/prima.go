package main

import "fmt"

func main() {
	var n int
	var faktor int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			faktor++
		}
	}

	fmt.Println()

	if faktor == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
