package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) // ini harusnya number overflow dan akan balik ke nilai -32768

	fmt.Println("Nilai 32 = ", nilai32)
	fmt.Println("Nilai 64 = ", nilai64)
	fmt.Println("Nilai 16 = ", nilai16)

	//string
	var name = "Rehandra"
	var e uint8 = name[0]     //ini akan return byte, bukan string literal
	var eToString = string(e) //convert byte ke string

	fmt.Println("Nama = ", name)
	fmt.Println("Huruf byte = ", e)
	fmt.Println("Huruf byte ke string = ", eToString)
}
