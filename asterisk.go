package main

import "fmt"

type Address struct {
	City, Province, Region string
}

func main() {
	var address1 Address = Address{"Padang", "Sumatra Barat", "Indonesia"}
	var address2 *Address = &address1 //merujuk ke alamat dimana addres satu disimpan

	address2.City = "Bukittinggi"

	//sedangkan yang ini berarti dia merubah value di alamat yang dirujuk
	*address2 = Address{"Payakumbuah", "Sumatra Barat", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
