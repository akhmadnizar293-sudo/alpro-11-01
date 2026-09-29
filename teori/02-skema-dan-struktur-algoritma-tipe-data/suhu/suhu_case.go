package main

import "fmt"

func main() {
	var umur int
	var suhu float64

	suhu = 36.5
	umur = 20

	fmt.Println("Suhu tubuh:", suhu)
	fmt.Println("Umur:", umur)
	fmt.Println("alamat memori dari suhu:", &suhu)
	fmt.Println("alamat memori dari umur:", &umur)
}
