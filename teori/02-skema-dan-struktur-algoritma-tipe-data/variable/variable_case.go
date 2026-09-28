package main

import "fmt"

func main() {

	//ini variable, variable itu buat nyimpen data/nilai. Kata pak Yudha

	//cara pertama
	var name string
	var age int

	name = "Akhmad Usluk Nizar"
	age = 20

	fmt.Println("Nama saya adalah", name)
	fmt.Println("Umur saya adalah", age, "tahun")

	//cara kedua
	var nameBelakang string = "Nizar"
	fmt.Println("Nama belakang saya adalah", nameBelakang)

	//cara ketiga
	middleName := "Usluk"
	fmt.Println("Nama tengah saya adalah", middleName)

	//cara keempat
	var (
		firstName string = "Akhmad"
		lastName  string = "Nizar"
	)
	fmt.Println("Nama depan saya adalah", firstName)
	fmt.Println("Nama belakang saya adalah", lastName)
}
