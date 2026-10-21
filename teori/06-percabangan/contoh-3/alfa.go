package main

import "fmt"

func main() {
	var kar int
	fmt.Scan("%c", &kar)
	if kar >= 'a' && kar <= 'z' || kar >= 'A' && kar <= 'Z' {
		fmt.Println("alphabet")
	}
	{
	}
}
