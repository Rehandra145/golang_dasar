package main

import "fmt"

func main() {

	//manual for loop
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke :", counter)
		counter++
	}

	fmt.Println("Selesai")

	//for loop dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke :", counter)
	}
	fmt.Println("Selesai")

	//for range
	//sebelum for range
	names := []string{
		"Rehan",
		"handera",
		"Rehandra",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println("Nama ke", i, "adalah", names[i])
	}

	//for range
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	//atau ga mau pakai index
	for _, name := range names {
		fmt.Println(name)
	}
}
