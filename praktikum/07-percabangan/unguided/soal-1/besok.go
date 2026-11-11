package main

import "fmt"

func main() {
	var tanggal, bulan, tahun int
	var jumlahHari int
	fmt.Scan(&tanggal, &bulan, &tahun)

	if bulan < 1 || bulan > 12 {
		fmt.Println("Tanggal tidak valid")
		return
	}

	if bulan == 2 {
		jumlahHari = 28

		if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
			jumlahHari = 29
		}
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		jumlahHari = 30
	} else {
		jumlahHari = 31
	}

	if tanggal < 1 || tanggal > jumlahHari {
		fmt.Println("Tanggal tidak valid")
		return
	}

	fmt.Println("Valid")
	tanggal++

	if tanggal > jumlahHari {
		tanggal = 1
		bulan++

		if bulan > 12 {
			bulan = 1
			tahun++
		}
	}

	fmt.Printf("%d/%d/%d\n", tanggal, bulan, tahun)
}
