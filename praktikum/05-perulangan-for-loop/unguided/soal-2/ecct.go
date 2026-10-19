package main

import "fmt"

func main() {
	var m int
	var jumlah int

	fmt.Scan(&m)

	for i := 1; i <= m; i++ {
		var p, v, ga, fs, ci, tc int

		fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

		total := p + v + ga + fs + ci + tc

		if float64(total)/6 >= 3.33 {
			fmt.Println("Ya")
			jumlah++
		} else {
			fmt.Println("Tidak")
		}
	}

	fmt.Println(jumlah)
}
