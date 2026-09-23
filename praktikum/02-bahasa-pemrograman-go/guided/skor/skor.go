package main

import "fmt"

func main() {
	var namaSiswa string
	var skorMtk uint8
	var skorBhsInggris uint8

	//membaca input

	fmt.Print("Masukkan nama siswa		: ")
	fmt.Scanln(&namaSiswa)
	fmt.Print("Masukkan skor Matematika	: ")
	fmt.Scanln(&skorMtk)
	fmt.Print("Masukkan skor Bahasa Inggris	: ")
	fmt.Scanln(&skorBhsInggris)

	//menghitung total skor, rata-rata, dan menampilkan output

	fmt.Println("================ DATA ==================")

	fmt.Println("Nama siswa		:", namaSiswa)
	fmt.Println("Total skor		:", skorMtk+skorBhsInggris)
	fmt.Println("Rata-rata		:", float64(skorMtk+skorBhsInggris)/2)
}
