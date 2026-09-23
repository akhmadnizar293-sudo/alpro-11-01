package main

import "fmt"

func main() {
	var jari float64
	var pi float64 = 3.14
	//membaca input

	fmt.Print("Masukkan jari-jari	: ")
	fmt.Scanln(&jari)

	//menghitung total skor, rata-rata, dan menampilkan output

	fmt.Println("================ OUTPUT ==================")

	fmt.Println("Luas lingkaran		:", pi*float64(jari)*float64(jari))

	fmt.Println("================ NAH INI ==================")
}
