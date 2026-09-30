package main

import "fmt"

func main() {
	var mil float64
	var km float64

	fmt.Print("Masukkan mil	: ")
	fmt.Scanln(&mil)

	km = mil * 1.6
	fmt.Printf("Jarak dalam Kilometer	: %.1f km", km)
}
