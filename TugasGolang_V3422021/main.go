package main

import (
	"fmt"
	"os"
)

// fungsi tambah untuk operasi penjumlahan
func tambah(bilangan1, bilangan2 int) int {
	return bilangan1 + bilangan2
}

// fungsi kurang untuk operasi pengurangan
func kurang(bilangan1, bilangan2 int) int {
	return bilangan1 - bilangan2
}

// fungsi kali untuk operasi perkalian
func kali(bilangan1, bilangan2 int) int {
	return bilangan1 * bilangan2
}

// fungsi bagi untuk operasi pembagian (bonus)
func bagi(bilangan1, bilangan2 int) (int, error) {
	if bilangan2 == 0 {
		return 0, fmt.Errorf("Pembagian dengan nol tidak diizinkan")
	}
	return bilangan1 / bilangan2, nil
}

func main() {
	// Deklarasi dan inisialisasi variabel
	var bilangan1, bilangan2 int

	// Input pengguna untuk bilangan1
	fmt.Print("Masukkan nilai bilangan1: ")
	_, err := fmt.Scan(&bilangan1)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Input pengguna untuk bilangan2
	fmt.Print("Masukkan nilai bilangan2: ")
	_, err = fmt.Scan(&bilangan2)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Memanggil fungsi-fungsi dan menampilkan hasilnya
	fmt.Printf("Hasil penjumlahan: %d\n", tambah(bilangan1, bilangan2))
	fmt.Printf("Hasil pengurangan: %d\n", kurang(bilangan1, bilangan2))
	fmt.Printf("Hasil perkalian: %d\n", kali(bilangan1, bilangan2))

	// Memanggil fungsi bagi dan menangani kesalahan pembagian dengan nol
	result, err := bagi(bilangan1, bilangan2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hasil pembagian: %d\n", result)
	}
}
