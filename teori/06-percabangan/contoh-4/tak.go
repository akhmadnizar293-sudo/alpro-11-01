package main

import "fmt"

func main() {
	var ikk float64
	fmt.Scan(&ikk)

	if ikk < 2.00 {
		fmt.Println("Poor")
	} else if ikk <= 2.00 && ikk <= 2.75 {
		fmt.Println("Fair")
	} else if ikk <= 2.756 && ikk <= 3.00 {
		fmt.Println("Satisfactory")
	} else if ikk <= 3.01 && ikk <= 3.50 {
		fmt.Println("Very Good")
	} else {
		fmt.Println("Excellent")
	}
}
