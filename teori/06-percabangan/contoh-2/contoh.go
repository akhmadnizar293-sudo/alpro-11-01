package main

import "fmt"

func main() {
	var nilai int
	fmt.Scan(&nilai)

	if nilai >= 85 && nilai <= 100 {
		fmt.Println("A")
	} else if nilai >= 70 && nilai <= 84 {
		fmt.Println("B")
	} else if nilai >= 50 && nilai <= 69 {
		fmt.Println("C")
	} else if nilai <= 50 && nilai >= 0 {
		fmt.Println("E")
	} else {
		fmt.Println("tidak valid")
	}
}
