package main

import "fmt"

func main() {
	var x int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)

	var sepuluhRibu int = x / 10000
	sisa := x % 10000

	var limaRibuan int = sisa / 5000
	sisa = x % 5000

	var seribu int = sisa / 1000

	fmt.Println("Hasil: ", sepuluhRibu, limaRibuan, seribu)

}
