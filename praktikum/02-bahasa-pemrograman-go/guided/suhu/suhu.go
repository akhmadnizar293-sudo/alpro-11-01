package main

import "fmt"

func main() {
	var celcsius float64
	//membaca input

	fmt.Print("Masukkan celsius	: ")
	fmt.Scanln(&celcsius)

	//menghitung total skor, rata-rata, dan menampilkan output

	fmt.Println("================ OUTPUT ==================")

	fmt.Println("Suhu dalam Reamur	:", celcsius*4/5)
	fmt.Println("Suhu dalam Fahrenheit	:", celcsius*9/5+32)
	fmt.Println("Suhu dalam Kelvin	:", celcsius+273.15)

	fmt.Println("================ NAH INI ==================")
}
