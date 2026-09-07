# PDF Toolkit (Go)

CLI offline untuk:
- Menggabungkan sampai 25 file PDF sekaligus
- Convert PDF -> JPG/PNG
- Convert JPG/PNG -> PDF

Semua pemilihan file pakai popup Windows Explorer (tidak perlu ketik path manual).



## Struktur

- `main.go` — loop menu utama + dispatch
- `menu.go` — tampilan pilihan menu
- `tipe.go` — konstanta & tipe data (array fixed, bukan slice)
- `filepicker.go` — semua popup file explorer
- `merge.go` — logic gabung PDF
- `konversi.go` — logic convert PDF<->gambar

