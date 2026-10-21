# <h1 align="center">Laporan Praktikum Modul 06 - Perulangan While Loop</h1>

<p align="center">Akhmad Nizar - 109092600018</p>

## Dasar Teori

### A. Pengenalan While Loop di GO

Perulangan while mengeksekusi blok kode selama ekspresi boolean bernilai benar. Perulangan ini berhenti segera setelah ekspresi bernilai salah. Ekspresi boolean dievaluasi sebelum setiap iterasi.

Alih-alih menggunakan kata kunci 'while', ia menggunakan kata kunci 'for' hanya dengan ekspresi Boolean.
Tidak ada loop do-while di Go.

Sintaksis

    for booleanExpression {
    //do something
    }

Catatan

Karena ekspresi boolean dievaluasi sebelum setiap iterasi, blok kode HANYA akan dieksekusi jika booleanExpression bernilai benar.

booleanExpression menghasilkan output berupa true atau false. Output ini dibuat menggunakan operator perbandingan (==, >, =, statement).

Contoh

    x := 0
    for x < 10 {
    	fmt.Println(x) //prints x until x < 10 evaluates to false
    	x++
    }

## Guided

### 1. terkecil.go

```go
package main

import "fmt"

func main() {
	var batas int
	fmt.Scan(&batas)

	n := 1
	for n*n*n <= batas {
		n++
	}

	fmt.Println(n)
}

```

##### Output

![Screenshot Output guided](praktikum\06-perulangan-while-loop-dan-repeat\guided\guided-1\output.png)

#### Deskripsi

Program di atas ialah bahasa Go yang menggunakan perulangan while (for kondisi) untuk mencari bilangan bulat terkecil n sehingga n³ (n pangkat 3) lebih besar dari 12000. Program harus menghitung nilai n mulai dari 1 ke atas hingga menemukan bilangan terkecil n yang memenuhi syarat tersebut.

### 2. terbesar.go

```go
package main

import "fmt"

func main() {
	var batas int
	fmt.Scan(&batas)

	n := 1
	for n*n < batas {
		n++
	}

	fmt.Println(n - 1)
}


```

##### Output

![Screenshot Output guided](guided\guided-2\output.png)

#### Deskripsi

Program dalam bahasa Go yang menggunakan perulangan while (for kondisi) untuk menemukan bilangan bulat terbesar n sedemikian rupa sehingga n² kurang dari batas yang diinputkan. Program harus mencari bilangan n yang memenuhi kondisi tersebut dan menampilkan nilai n terbesar yang memenuhi syarat.

### 3. vokal.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	kalimat, _ := reader.ReadString('\n')

	i := 0
	jumlahVokal := 0

	for i < len(kalimat) {
		ch := strings.ToLower(string(kalimat[i]))
		if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
			jumlahVokal++
		}
		i++
	}

	fmt.Println(jumlahVokal)
}

```

##### Output

![Screenshot Output Unguided](guided\guided-3\output.png)

#### Deskripsi

Program dalam bahasa Go yang meminta pengguna memasukkan sebuah kata atau kalimat. Program akan menghitung berapa banyak huruf vokal (a, i, u, e, o) yang terdapat dalam kata atau kalimat tersebut, baik huruf kecil maupun huruf kapital. Program harus menggunakan perulangan while (for kondisi) dengan indeks untuk memeriksa setiap karakter dan berhenti setelah semua karakter dalam string diperiksa.

## Unguided

### 1. atm.go

```go
package main

import "fmt"

func main() {
	var password string
	var kesempatan, saldo, pilihan, tarik int = 0, 10000000, 0, 0

	for password != "1234" && kesempatan < 3 {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		if password != "1234" {
			kesempatan++
			fmt.Println("Password salah. Silakan coba lagi.")
		}
	}

	if password != "1234" {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}
	fmt.Println("Login berhasil!")

	for pilihan != 3 {
		fmt.Println("Menu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		} else if pilihan == 2 {
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
			fmt.Scan(&tarik)

			if tarik > 0 && tarik <= saldo {
				saldo -= tarik
				fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\n", tarik)
				fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		} else if pilihan == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

```

##### Output

![Screenshot Output Unguided](unguided\soal-1\output.png)

#### Deskripsi

Program di atas adalah program atm yang dimana kita disuruh inputkan password "1234" untuk melanjutkan ke dalam program, jika salah password sebanyak 3 kali maka akan muncul output "akun diblokir". Program ini menggunakan perulangan while.

## Kesimpulan

Praktikum ini bertujuan untuk memahami perulangan while serta implementasinya dalam coding.

## Referensi

1. _syntaxdb.com_.
