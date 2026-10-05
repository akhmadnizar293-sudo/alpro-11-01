package main

import "fmt"

func main() {
	var x, y, z int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)

	fmt.Print("Masukkan nilai y: ")
	fmt.Scanln(&y)

	fmt.Print("Masukkan nilai z: ")
	fmt.Scanln(&z)

	temp := x
	x = z
	z = y
	y = temp

	fmt.Println("Nilai x setelah ditukar: ", x)
	fmt.Println("Nilai y setelah ditukar: ", y)
	fmt.Println("Nilai z setelah ditukar: ", z)
}
