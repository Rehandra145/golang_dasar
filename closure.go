package main

import "fmt"

//closure adalah kemampuan sebuah func berinteraksi dengan data data di sekitarnya dalam scope yang sama
//Harap gunakan fitur closure dengan bijak karna jika dipakai terlalu banyak akan membuat bingung kode program

func main() {
	counter := 0

	increment := func() {
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)

}
