package main

import "fmt"

func main() {
	//konstan itu sesuai namanya, value nya hanya bisa di assign sekali ketika deklarasi
	//konstan tidak bisa diubah tapi jika ttidak dipakai tidak masalah ketika di compile
	//kalau diubah akan error

	const name string = "Rehandra" //static typing bisa
	const name2 = "Rehandra"       //dynamic typing juga bisa

	//sama seperti variable, konstan juga bisa multiple declaration
	const (
		firstName = "Rehandra"
		lastName  = "Pratama"
	)

	fmt.Println(firstName, lastName)
}
