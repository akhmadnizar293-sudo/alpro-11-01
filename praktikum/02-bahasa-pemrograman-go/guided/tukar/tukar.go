package main

import "fmt"

func main() {
	var a, b uint8

	//membaca input

	fmt.Print("Masukkan nilai a		: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai b		: ")
	fmt.Scanln(&b)

	//menghitung total skor, rata-rata, dan menampilkan output

	fmt.Println("================ OUTPUT ==================")

	fmt.Println("Nilai a		:", b)
	fmt.Println("Nilai b		:", a)

	fmt.Println("================ NILAI A dan B ditukar ==================")
}
