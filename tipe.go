package main

// batas maksimal file yang bisa diproses sekaligus
// (dipakai untuk merge PDF, dan untuk daftar gambar saat convert gambar->PDF)
const nmax int = 25

// daftarFile adalah array fixed-size, bukan slice
// jumlah slot yang benar-benar terisi dilacak manual pakai variabel counter (n)
type daftarFile [nmax]string
