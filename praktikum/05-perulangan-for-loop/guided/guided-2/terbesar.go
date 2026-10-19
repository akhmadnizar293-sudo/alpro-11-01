package main

import "fmt"

func main() {
	var n, bil, terbesar int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print("Masukkan bilangan: ")
		fmt.Scan(&bil)

		if i == 1 || bil > terbesar {
			terbesar = bil
		}
	}

	fmt.Println("Keluaran: ", terbesar)
}
