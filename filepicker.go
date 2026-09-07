package main

import (
	"fmt"

	"github.com/ncruces/zenity"
)

// pilihSatuPDF menampilkan popup Windows Explorer untuk memilih 1 file PDF
func pilihSatuPDF() (string, bool) {
	var lokasi string
	var err error

	lokasi, err = zenity.SelectFile(
		zenity.FileFilter{Name: "PDF files", Patterns: []string{"*.pdf"}},
	)
	if err != nil {
		fmt.Println("Tidak ada file yang dipilih")
		return "", false
	}
	return lokasi, true
}

// pilihBanyakPDF menampilkan popup untuk memilih banyak file PDF sekaligus (maks nmax)
func pilihBanyakPDF() (daftarFile, int) {
	var hasil daftarFile
	var n int
	var pilihan []string
	var err error
	var i int

	pilihan, err = zenity.SelectFileMultiple(
		zenity.FileFilter{Name: "PDF files", Patterns: []string{"*.pdf"}},
	)
	if err != nil {
		fmt.Println("Tidak ada file yang dipilih")
		return hasil, 0
	}

	for i = 0; i < len(pilihan) && i < nmax; i++ {
		hasil[i] = pilihan[i]
		n++
	}
	if len(pilihan) > nmax {
		fmt.Println("Kamu memilih lebih dari 25 file, hanya 25 file pertama yang diproses")
	}
	return hasil, n
}

// pilihBanyakGambar menampilkan popup untuk memilih banyak file gambar sekaligus (maks nmax)
func pilihBanyakGambar() (daftarFile, int) {
	var hasil daftarFile
	var n int
	var pilihan []string
	var err error
	var i int

	pilihan, err = zenity.SelectFileMultiple(
		zenity.FileFilter{Name: "Image files", Patterns: []string{"*.jpg", "*.jpeg", "*.png"}},
	)
	if err != nil {
		fmt.Println("Tidak ada file yang dipilih")
		return hasil, 0
	}

	for i = 0; i < len(pilihan) && i < nmax; i++ {
		hasil[i] = pilihan[i]
		n++
	}
	if len(pilihan) > nmax {
		fmt.Println("Kamu memilih lebih dari 25 file, hanya 25 file pertama yang diproses")
	}
	return hasil, n
}

// pilihLokasiSimpan menampilkan popup untuk menentukan lokasi & nama file hasil
func pilihLokasiSimpan(namaDefault string, ekstensi string) (string, bool) {
	var lokasi string
	var err error

	lokasi, err = zenity.SelectFileSave(
		zenity.Filename(namaDefault+"."+ekstensi),
		zenity.FileFilter{Name: ekstensi + " files", Patterns: []string{"*." + ekstensi}},
	)
	if err != nil {
		fmt.Println("Proses penyimpanan dibatalkan")
		return "", false
	}
	return lokasi, true
}
