package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	mobil := n / 7
	if n%7 != 0 {
		mobil++
	}

	fmt.Println("Jumlah mobil yang diperlukan: ", mobil)
}
