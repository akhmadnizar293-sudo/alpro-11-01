package main

import (
	"fmt"
	"math"
)

func main() {
	var berat float64
	var zona int
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Println("Berat melebihi batas maksimum (30 kg)")
		return
	}

	var tarifPerKg int

	switch zona {
	case 1:
		tarifPerKg = 10000
	case 2:
		tarifPerKg = 15000
	case 3:
		tarifPerKg = 25000
	default:
		fmt.Println("Zona nguwawor")
		return
	}

	beratBulat := int(math.Ceil(berat))
	total := beratBulat * tarifPerKg

	if layanan == "Express" {
		total += total / 2
	}

	fmt.Println(beratBulat)
	fmt.Println(total)
}
