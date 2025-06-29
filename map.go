package main

import "fmt"

func main() {
	//tipe data map, hash map lah ya
	//key => value
	//panjangnya boleh bebas sebanyak banyaknya, asalkan key nya unik
	//kalau sama key nya, maka value yang terakhir yang akan disimpan

	//ini cara pertama
	var orang map[string]string = map[string]string{}
	orang["name"] = "Rehan"
	orang["address"] = "Agam"

	person := map[string]string{
		"name":    "Rehan",
		"address": "Agam",
	}

	fmt.Println(orang)

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//function yang dapat digunakan di map
	//1. len(map)
	//untuk mengetahui panjang dari map
	fmt.Println("Panjang dari map person : ", len(person))
	//2. map[key]
	//untuk mengambil value dari key yang ada di map
	fmt.Println("value dari key name di map person : ", person["name"])
	//3. map[key] = value
	//untuk mengubah data di dalam map dengan key
	fmt.Println("Sebelum diubah : ", person["name"])
	person["name"] = "handera"
	fmt.Println("Setelah diubah : ", person["name"])
	//4. make(map[TypeKey]TypeValue)
	//untuk membuat map baru
	book := make(map[string]string) //book := map[string]string{} => kek gini juga boleh
	book["title"] = "Belajar Golang"
	book["author"] = "Rehandra"
	book["wrong"] = "Ups"
	fmt.Println(book)
	fmt.Println(book["title"])
	fmt.Println(book["author"])
	fmt.Println(book["wrong"])
	//5. delete(map, key)
	//menghapus data didalam map sesuai dengan key
	delete(book, "wrong")
	fmt.Println(book)
}
