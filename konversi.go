package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/responses"
	"github.com/klippa-app/go-pdfium/webassembly"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// gambarKePDF menggabungkan sampai nmax file gambar (jpg/png) menjadi satu file PDF
func gambarKePDF() {
	var gambar daftarFile
	var n int
	var i int
	var lokasiSimpan string
	var ok bool
	var daftarUntukLibrary []string
	var err error

	gambar, n = pilihBanyakGambar()
	if n == 0 {
		return
	}

	lokasiSimpan, ok = pilihLokasiSimpan("hasil-konversi", "pdf")
	if !ok {
		return
	}

	for i = 0; i < n; i++ {
		daftarUntukLibrary = append(daftarUntukLibrary, gambar[i])
	}

	err = api.ImportImagesFile(daftarUntukLibrary, lokasiSimpan, nil, nil)
	if err != nil {
		fmt.Println("Gagal mengonversi gambar ke PDF:", err)
		return
	}

	fmt.Println("Berhasil disimpan ke:", lokasiSimpan)
}

// pdfKeGambar merender setiap halaman PDF menjadi file gambar terpisah (jpg/png)
func pdfKeGambar() {
	var lokasi string
	var ok bool
	var format string
	var err error
	var pool pdfium.Pool
	var instance pdfium.Pdfium
	var isiFile []byte
	var dokumenPtr *responses.OpenDocument
	var jumlahHalamanPtr *responses.FPDF_GetPageCount
	var i int

	lokasi, ok = pilihSatuPDF()
	if !ok {
		return
	}

	fmt.Print("Simpan sebagai jpg atau png? ")
	fmt.Scan(&format)
	format = strings.ToLower(format)
	if format != "jpg" && format != "png" {
		fmt.Println("Format tidak dikenali")
		return
	}

	// baca seluruh isi file PDF jadi byte, karena go-pdfium butuh isi
	// filenya langsung (bukan sekadar lokasi/path-nya)
	isiFile, err = os.ReadFile(lokasi)
	if err != nil {
		fmt.Println("Gagal membaca file PDF:", err)
		return
	}

	pool, err = webassembly.Init(webassembly.Config{
		MinIdle:  1,
		MaxIdle:  1,
		MaxTotal: 1,
	})
	if err != nil {
		fmt.Println("Gagal menyiapkan mesin render:", err)
		return
	}
	defer pool.Close()

	instance, err = pool.GetInstance(time.Second * 30)
	if err != nil {
		fmt.Println("Gagal mendapatkan instance:", err)
		return
	}
	defer instance.Close()

	dokumenPtr, err = instance.OpenDocument(&requests.OpenDocument{
		File: &isiFile,
	})
	if err != nil {
		fmt.Println("Gagal membuka PDF:", err)
		return
	}
	defer instance.FPDF_CloseDocument(&requests.FPDF_CloseDocument{
		Document: dokumenPtr.Document,
	})

	jumlahHalamanPtr, err = instance.FPDF_GetPageCount(&requests.FPDF_GetPageCount{
		Document: dokumenPtr.Document,
	})
	if err != nil {
		fmt.Println("Gagal membaca jumlah halaman:", err)
		return
	}

	for i = 0; i < jumlahHalamanPtr.PageCount; i++ {
		var render *responses.RenderPageInDPI
		var errRender error
		var namaFile string
		var indexHalaman int

		indexHalaman = i
		render, errRender = instance.RenderPageInDPI(&requests.RenderPageInDPI{
			Page: requests.Page{
				ByIndex: &requests.PageByIndex{
					Document: dokumenPtr.Document,
					Index:    indexHalaman,
				},
			},
			DPI: 150,
		})
		if errRender != nil {
			fmt.Println("Gagal merender halaman", i+1, ":", errRender)
			continue
		}

		namaFile = lokasi + "-halaman" + strconv.Itoa(i+1) + "." + format
		simpanGambar(render.Result.Image, namaFile, format)
		render.Cleanup()
	}

	fmt.Println("Selesai. Total", jumlahHalamanPtr.PageCount, "halaman disimpan.")
}

// simpanGambar menyimpan satu gambar hasil render ke disk dalam format jpg/png
func simpanGambar(img image.Image, namaFile string, format string) {
	var f *os.File
	var err error

	f, err = os.Create(namaFile)
	if err != nil {
		fmt.Println("Gagal membuat file:", namaFile, err)
		return
	}
	defer f.Close()

	if format == "jpg" {
		jpeg.Encode(f, img, &jpeg.Options{Quality: 90})
	} else {
		png.Encode(f, img)
	}
	fmt.Println("Tersimpan:", namaFile)
}
