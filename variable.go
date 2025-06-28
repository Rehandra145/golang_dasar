package main

import "fmt"

func main() {

	// ini static typing
	var name string

	name = "Rehandra"
	fmt.Println(name)

	name = "Rehandra Pratama"
	fmt.Println(name)

	//dynamic typing bisa kek gini
	var nama = "Rehandra"
	fmt.Println(nama)

	nama = "Rehandra Pratama"
	fmt.Println(nama)

	//sebenernya kata kunci var waktu inisialisasi variabel itu  tidak wajib, bisa aja langsung kek gini :
	nama2 := "Rehandra" // tapi ini hanya untuk inisialisasi variabel baru, untuk mengganti nilai tetap kayak biasa
	fmt.Println(nama2)

	//ada juga yang namanya multiple variable
	var (
		firstName = "Rehandra"
		lastName  = "Pratama"
	)

	fmt.Println(firstName, lastName)
}
