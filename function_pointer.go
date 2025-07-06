package main

import "fmt"

/*
  disini kalau ingin mengubah suatu value lewat fungsi yang dilempar lewat paramater
	itu tidak bisa seperti biasa seperti di bahasa pemograman oop lainnya, karena oop itu sendiri
	tidak berlaku di go, jadi kalau ingin mengubah sebuah value lewat function yang diterima lewat parameter
	maka harus meng akses langsung ke memori
*/
type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	address := Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
