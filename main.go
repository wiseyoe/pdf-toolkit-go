package main

import "fmt"

func main() {
	var pilihan int

	for {
		menu()
		fmt.Scan(&pilihan)

		for pilihan < 1 || pilihan > 4 {
			fmt.Println("Tolong masukan angka sesuai opsi")
			fmt.Scan(&pilihan)
		}

		switch pilihan {
		case 1:
			gabungPDF()
		case 2:
			pdfKeGambar()
		case 3:
			gambarKePDF()
		}

		if pilihan == 4 {
			break
		}
	}
}
