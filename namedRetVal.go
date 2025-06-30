package main

import "fmt"

func getCompleteName() (name, alamat, pacar string) {
	name = "Rehan"
	alamat = "Agam"
	pacar = "Syarifah 😋"
	return name, alamat, pacar
}

func main() {
	nama, alamat, _ := getCompleteName()
	fmt.Println("Nama saya", nama, "dari", alamat)
}
