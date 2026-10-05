package main

import "fmt"

func main() {
	var c, r float64

	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scanln(&c)

	r = (4.0 / 5.0) * c
	fmt.Println("Suhu dalam Fahrenheit: ", r)
}
