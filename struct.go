package main

import "fmt"

/*
	struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data
	lainnya dalam satu kesatuan.
	Struct biasanya representasi data dalam proram aplikasi yang dibuat
	Data di struct disimpan di dalam field
	Sederhananya struct adalah kumpulan dari field
*/

type Customer struct {
	Name, Address string
	Age           int
}

/*
	Membuat Data Struct
	Struct adalah template data atau prototype data
	Struct tidak bisa langsung digunakan
	Namun kita bisa membuat data/object dari struct yang telah kita buat
*/

//ini adalah method, dia nempel di obek struct, method adalah func
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {

	// ini cara manual
	var Rehan Customer
	Rehan.Name = "Rehandra"
	Rehan.Address = "Indonesia"
	Rehan.Age = 21
	fmt.Println(Rehan)
	fmt.Println(Rehan.Name)
	fmt.Println(Rehan.Address)
	fmt.Println(Rehan.Age)

	//pemanggilan kek gini
	Rehan.sayHello("Dea")

	//struct literal

	Abit := Customer{
		Name:    "Abit",
		Address: "Merak",
		Age:     19,
	}

	fmt.Println(Abit)
	fmt.Println(Abit.Name)
	fmt.Println(Abit.Address)
	fmt.Println(Abit.Age)

	Budi := Customer{"Budi", "Amerika", 25}

	fmt.Println(Budi)
	fmt.Println(Budi.Name)
	fmt.Println(Budi.Address)
	fmt.Println(Budi.Age)
}
