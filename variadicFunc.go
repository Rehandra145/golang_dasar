package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
func main() {
	fmt.Println(sumAll(5, 5, 5, 5, 5))

	//jika sudah ada slice
	//maka bisa langsung dioper ke fungsi, tapi harus menggunakan ...
	//jika tidak menggunakan ... maka akan error
	//karena slice itu bukan variadic function
	//jadi harus dioper ke fungsi dengan ... agar bisa diiterasi
	//jadi ... itu menandakan bahwa kita mengoper slice ke fungsi
	//jadi variadic function itu bisa menerima banyak parameter
	//jadi kita bisa mengoper slice ke variadic function
	numbers := []int{10, 10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
