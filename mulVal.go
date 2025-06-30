package main

import "fmt"

func getFullName() (string, string) {
	return "Rehan", "Agam"
}
func main() {
	name, alamat := getFullName()
	fmt.Println("Nama saya", name, "dari", alamat)

	//multiple value wajib ditagkap semua valuenya
	//kalau ingin dihiraukan, maka tinggal ganti jadi '_' saja

	nama, _ := getFullName()
	fmt.Println("Nama saya", nama)
}
