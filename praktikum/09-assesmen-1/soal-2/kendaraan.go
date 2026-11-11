package main

import "fmt"

func main() {
	var tipeKendaraan string
	var m, k, t int

	for {
		fmt.Scan(&tipeKendaraan)
		if tipeKendaraan == "M" {
			m++
		} else if tipeKendaraan == "K" {
			k++
		} else if tipeKendaraan == "T" {
			t++
		} else {
			break
		}
	}

	fmt.Println("Tipe M =", m)
	fmt.Println("Tipe K =", k)
	fmt.Println("Tipe T =", t)
}
