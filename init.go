package main

import (
	"belajar/database"

	/*
		ini namannya blank identifier, gunanya biar go tidak protes ketika kita memanggil sebuah package
		namun tidak memanggil fungsi fungsi yang ada di dalamnnya, tetapi seperti pada umumnya, jika
		package yang dipanggill mempunyai func init didalamnya, maka func init nya tetap akan terpanggil
		jika menggunakan blank identifier. namun kalau tidak menggunakan blank identifier maka func init
		juga tidak akan berjalan, malahan kadang di beberapa text editor jika memanggil sebuah package namun
		tidak pernah digunakan, ketika di save package yang dipanggil malah hilang.
	*/
	_ "belajar/internal"
	"fmt"
)

func main() {
	database := database.GetDatabase()
	fmt.Println(database)
}
