# PDF Toolkit (Go)

CLI offline untuk:
- Menggabungkan sampai 25 file PDF sekaligus
- Convert PDF -> JPG/PNG
- Convert JPG/PNG -> PDF

Semua pemilihan file pakai popup Windows Explorer (tidak perlu ketik path manual).

## Setup

1. Install Go dari https://go.dev/dl
2. Buka folder ini di VS Code, install extension "Go" (publisher `golang.go`)
3. Jalankan di terminal:

```
go mod tidy
```

Ini akan otomatis download dependency yang dibutuhkan:
- `github.com/pdfcpu/pdfcpu` (merge PDF, import gambar ke PDF)
- `github.com/sqweek/dialog` (popup file explorer)
- `github.com/klippa-app/go-pdfium` (render PDF ke gambar)

## Menjalankan

```
go run .
```

## Build jadi .exe (Windows)

```
go build -o pdf-toolkit.exe .
```

## Struktur

- `main.go` — loop menu utama + dispatch
- `menu.go` — tampilan pilihan menu
- `tipe.go` — konstanta & tipe data (array fixed, bukan slice)
- `filepicker.go` — semua popup file explorer
- `merge.go` — logic gabung PDF
- `konversi.go` — logic convert PDF<->gambar

## Catatan gaya kode

Project ini sengaja menghindari `:=` dan lebih memilih `append`/slice
kecuali di titik-titik yang memang diwajibkan oleh Go (mis. saat
memanggil fungsi library eksternal yang butuh `[]string`, atau hasil
fungsi multi-return yang variabelnya baru). Struktur data utama
(daftar file yang dipilih user) pakai array fixed-size (`nmax = 25`)
dengan counter manual, bukan slice yang tumbuh otomatis.
