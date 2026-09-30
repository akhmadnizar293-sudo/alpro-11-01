package main

import "fmt"

func main() {
	var bool1, bool2 bool

	fmt.Print("Masukkan boolean pertama (true/false): ")
	fmt.Scan(&bool1)
	fmt.Print("Masukkan boolean kedua (true/false): ")
	fmt.Scan(&bool2)

	fmt.Println(bool1)
	fmt.Println(bool2)
}
