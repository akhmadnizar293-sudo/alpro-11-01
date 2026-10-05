# <h1 align="center">Tugas Pendahuluan Modul 003 - Tipe Data dan Instruksi Dasar</h1>

<p align="center">Laporan Tugas Pendahuluan - 109092600018</p>

### 1. sisa.go

```go
package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scan(&y)
	fmt.Println(x % y)
}

```

##### Output

![Screenshot Output Unguided](tp\sisa\output.png)

#### Deskripsi

Program di atas dia membaca inputan lalu mencari hasil dari sisa pembagian dari variable x dan y.

### 2. bool.go

```go
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

```

##### Output

![Screenshot Output Unguided](tp\bool\output.png)

#### Deskripsi

Program di atas membaca inputan untuk menentukan true atau false, 1 = true dan 2 = false

### 3. konversi.go

```go
package main

import "fmt"

func main() {
	var mil float64
	var km float64

	fmt.Print("Masukkan mil	: ")
	fmt.Scanln(&mil)

	km = mil * 1.6
	fmt.Printf("Jarak dalam Kilometer	: %.1f km", km)
}
```

##### Output

![Screenshot Output Unguided](tp\konversi\output.png)

#### Deskripsi

Program di atas membaca inputan untuk mengkonversikan mil menjadi kilometer dengan membaca inputan mil yang kita inputkan.

## Kesimpulan

Kesimpulannya adalah mempelajari tentang variable dan operator, juga bagaimana cara mengetikkannya ke dalam sebuah kode pemrograman dengan bahasa Go.
