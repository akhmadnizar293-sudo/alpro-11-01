package main

import "fmt"

func main() {
	x := 11
	y := 5
	fmt.Println(x == 10 || y == 5)
	fmt.Println(x == 10 && y == 5) // false

}
