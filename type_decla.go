package main

import "fmt"

func main() {
	//type declaration adalah pengaliasan nama sebuah tipe data
	//tujuannya agar lebih mudah dalam membaca kode
	//kata kuncinya adalah type diikuti dengan nama tipe data
	type NoKTP string

	var NoKtpRehan NoKTP = "123"

	fmt.Println("No KTP Rehan = ", NoKtpRehan)

	//tipe data yang di aliaskan juga bisa digunakan untuk konversi
	var contoh = "2222"
	var NoKtpContoh NoKTP = NoKTP(contoh) //disini variabel contoh yang bertipe string di konversi menjadi tipe data NoKTP yang dialiaskan dari tipe data string
	fmt.Println("No KTP Contoh = ", NoKtpContoh)
}
