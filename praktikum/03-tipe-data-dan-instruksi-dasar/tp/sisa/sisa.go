package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scan(&y)

	fmt.Println("Sisa pembagian:", x%y)
}
