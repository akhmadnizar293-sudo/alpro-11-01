# <h1 align="center">Tugas Pendahuluan Modul 004 - Tipe Data dan Instruksi Dasar</h1>

<p align="center">Laporan Tugas Pendahuluan - 109092600018</p>

### 1. soal_satu.go

```go
package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	fmt.Println(intNum > 5 && intOther > 5 && sngNum > 0)
}
```

##### Output

![Screenshot Output Unguided](tp\soalSatu\output.png)

#### Deskripsi

Program di atas dia membaca inputan lalu mencari hasil true atau false dari:

    intNum > 5 && intOther > 5 && sngNum > 0

### 2. soal_dua.go

```go
package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}

	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}

	fmt.Println("Nilai akhir result:", result)
}
```

##### Output

![Screenshot Output Unguided](tp\soalDua\output.png)

#### Deskripsi

Program di atas menghitung nilai akhir result berdasarkan beberapa kondisi if, else if, dan operator logika.

### 3. soal_tiga.go

```go
package main

import "fmt"

func main() {
	var tahun int
	var bulan string
	var hari int

	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	fmt.Print("Masukkan bulan: ")
	fmt.Scan(&bulan)

	if bulan == "Jan" || bulan == "Mar" || bulan == "Mei" ||
		bulan == "Jul" || bulan == "Agu" || bulan == "Okt" || bulan == "Des" {
		hari = 31
	} else if bulan == "Apr" || bulan == "Jun" || bulan == "Sep" || bulan == "Nov" {
		hari = 30
	} else if bulan == "Feb" {
		if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
			hari = 29
		} else {
			hari = 28
		}
	} else {
		fmt.Println(`"` + bulan + `" tidak valid`)
		return
	}

	fmt.Println(hari)
}

```

##### Output

![Screenshot Output Unguided](tp\soalTiga\output.png)

#### Deskripsi

Program di atas membaca inputan untuk Menentukan Jumlah Hari dalam Sebulan Berdasarkan Tahun dan
Bulan.

### 3. soal_empat.go

```go
package main

import "fmt"

func main() {
	var hari int

	fmt.Print("Masukkan angka (1-7): ")
	fmt.Scan(&hari)

	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	}
}
```

##### Output

![Screenshot Output Unguided](tp\soalEmpat\output.png)

#### Deskripsi

Program di atas membaca inputan 1 - 7 yang nanti akan memilih hari dengan cara switch

## Kesimpulan

Kesimpulannya adalah mempelajari tentang operator dan switch case, juga bagaimana cara mengetikkannya ke dalam sebuah kode pemrograman dengan bahasa Go.
