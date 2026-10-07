package main

import "fmt"

func main() {
	var tahun int
	var bulan string
	var hari int

	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	fmt.Print("Masukkan bulan: ")
	fmt.Scan(&bulan)

	if bulan == "Jan" || bulan == "Mar" || bulan == "Mei" ||
		bulan == "Jul" || bulan == "Agu" || bulan == "Okt" || bulan == "Des" {
		hari = 31
	} else if bulan == "Apr" || bulan == "Jun" || bulan == "Sep" || bulan == "Nov" {
		hari = 30
	} else if bulan == "Feb" {
		if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
			hari = 29
		} else {
			hari = 28
		}
	} else {
		fmt.Println(`"` + bulan + `" tidak valid`)
		return
	}

	fmt.Println(hari)
}
