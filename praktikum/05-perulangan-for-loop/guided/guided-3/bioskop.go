package main

import "fmt"

func main() {
	var t int
	fmt.Print("Masukkan jumlah transaksi: ")
	fmt.Scan(&t)

	totalPendapatan := 0

	for i := 1; i <= t; i++ {
		var film, hari string
		var tiket, harga, total int

		fmt.Print("Masukkan film: ")
		fmt.Scan(&film)
		fmt.Print("Masukkan hari: ")
		fmt.Scan(&hari)
		fmt.Print("Masukkan jumlah tiket: ")
		fmt.Scan(&tiket)

		switch film {
		case "Reguler":
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				harga = 40000
			} else {
				harga = 50000
			}
		case "3D":
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				harga = 60000
			} else {
				harga = 70000
			}
		case "IMAX":
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				harga = 80000
			} else {
				harga = 90000
			}
		}

		total = harga * tiket

		if tiket >= 5 {
			total = total - total*10/100
		}

		fmt.Println("Total: ", total)
		totalPendapatan += total
	}

	fmt.Println("Total Pendapatan: ", totalPendapatan)
}
