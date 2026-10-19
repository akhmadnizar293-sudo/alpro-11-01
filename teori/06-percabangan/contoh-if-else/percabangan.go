package main

import "fmt"

func main() {
	var lamp_status string

	fmt.Print("Masukkan status lampu (merah/kuning/hijau): ")
	fmt.Scan(&lamp_status)

	if lamp_status == "merah" {
		fmt.Println("Berhenti")
	} else if lamp_status == "kuning" {
		fmt.Println("Hati-hati")
	} else if lamp_status == "hijau" {
		fmt.Println("Jalan")
	} else {
		fmt.Println("tidak valid")
	}

}
