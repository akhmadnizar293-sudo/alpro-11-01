package main

import "fmt"

func main() {
	var celcius float64
	var kelvin float64

	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scanln(&celcius)

	kelvin = celcius + 273
	fmt.Println("Suhu dalam Kelvin: ", kelvin)
}
