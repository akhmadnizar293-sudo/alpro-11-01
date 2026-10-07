package main

import "fmt"

func main() {
	var penghasilan, pajak float64

	fmt.Print("Masukkan penghasilan: ")
	fmt.Scan(&penghasilan)

	if penghasilan <= 50 {
		pajak = 5.0 / 100 * penghasilan
	} else if penghasilan <= 100 {
		pajak = (5.0 / 100 * 50) +
			(10.0 / 100 * (penghasilan - 50))
	} else if penghasilan <= 200 {
		pajak = (5.0 / 100 * 50) +
			(10.0 / 100 * 50) +
			(15.0 / 100 * (penghasilan - 100))
	} else {
		pajak = (5.0 / 100 * 50) +
			(10.0 / 100 * 50) +
			(15.0 / 100 * 100) +
			(20.0 / 100 * (penghasilan - 200))
	}

	fmt.Printf("Pajak yang harus dibayar: %.1f\n", pajak)
}
