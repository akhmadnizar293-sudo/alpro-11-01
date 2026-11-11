package main

import "fmt"

func main() {
	var lembab, total, jem, kali int
	var rataRata float64

	rendah := 900
	tinggi := 0

	for {
		fmt.Scan(&lembab)

		if lembab < rendah {
			rendah = lembab
		}
		if lembab > tinggi {
			tinggi = lembab
		}
		if lembab == jem {
			break
		}
		kali++
		total += lembab
		jem = lembab
	}
	rataRata = float64(total) / float64(kali)
	fmt.Println(rendah, tinggi, rataRata)

}
