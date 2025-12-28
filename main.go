package main

import "fmt"

const NMAX = 40

/* ITERATIF */
func iteratif(nilai [NMAX]int, n int) float64 {
	total := 0

	for i := 0; i < n; i++ {
		total = total + nilai[i]
	}

	return float64(total) / float64(n)
}

/* REKURSIF*/
func rekursif(nilai [NMAX]int, n int) int {
	if n == 0 {
		return 0
	}

	return nilai[n-1] + rekursif(nilai, n-1)
}

func main() {
	var nilai [NMAX]int
	n := 0

	fmt.Print("Masukkan jumlah mahasiswa (maks 40): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan nilai mahasiswa ke-", i+1, ": ")
		fmt.Scan(&nilai[i])
	}

	fmt.Println("\n=== ITERATIF ===")
	fmt.Println("Rata-rata nilai:", iteratif(nilai, n))

	fmt.Println("\n=== REKURSIF ===")
	fmt.Println("Rata-rata nilai:", float64(rekursif(nilai, n))/float64(n))
}
