package main

import "fmt"

func main() {
	var celsius float64
	//membaca input

	fmt.Print("Masukkan celsius	: ")
	fmt.Scanln(&celsius)

	//menghitung total skor, rata-rata, dan menampilkan output

	fmt.Println("================ OUTPUT ==================")

	fmt.Println("Suhu dalam Reamur	:", celsius*4/5)
	fmt.Println("Suhu dalam Fahrenheit	:", celsius*9/5+32)
	fmt.Println("Suhu dalam Kelvin	:", celsius+273.15)

	fmt.Println("================ NAH INI ==================")
}
