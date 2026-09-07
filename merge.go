package main

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// gabungPDF menggabungkan sampai nmax file PDF menjadi satu file
func gabungPDF() {
	var files daftarFile
	var n int
	var i int
	var lokasiSimpan string
	var ok bool
	var daftarUntukLibrary []string
	var err error

	files, n = pilihBanyakPDF()
	if n == 0 {
		return
	}

	fmt.Println("File yang akan digabung:")
	for i = 0; i < n; i++ {
		fmt.Println(i+1, "-", files[i])
	}

	lokasiSimpan, ok = pilihLokasiSimpan("hasil-gabungan", "pdf")
	if !ok {
		return
	}

	// pdfcpu butuh []string, jadi isi array manual kita dipindah ke slice
	// hanya di titik pemanggilan library ini saja, bukan bagian dari logic internal kita
	for i = 0; i < n; i++ {
		daftarUntukLibrary = append(daftarUntukLibrary, files[i])
	}

	err = api.MergeCreateFile(daftarUntukLibrary, lokasiSimpan, false, nil)
	if err != nil {
		fmt.Println("Gagal menggabungkan PDF:", err)
		return
	}

	fmt.Println("Berhasil digabung ke:", lokasiSimpan)
}
