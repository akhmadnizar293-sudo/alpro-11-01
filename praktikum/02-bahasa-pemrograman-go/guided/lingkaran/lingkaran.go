package main

import "fmt"

func main() {
	var r float64
	var pi float64 = 3.14

	fmt.Print("Masukkan jari-jari	: ")
	fmt.Scanln(&r)

	fmt.Println("================ OUTPUT ==================")
	fmt.Println("Luas lingkaran		:", pi*(r)*(r))
	fmt.Println("================ NAH INI ==================")
}
