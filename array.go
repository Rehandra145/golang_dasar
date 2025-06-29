package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Rehan"
	names[1] = "handera"
	names[2] = "Rehandra"

	fmt.Println("Nama saya adalah: ", names[0])
	fmt.Println("Nama instagram saya adalah: ", names[1])
	fmt.Println("Nama lengkap saya adalah: ", names[2])

	//bentuk lain deklarasi array
	var numbers = [5]int{
		1, 2, 3, 4,
	}

	fmt.Println("Angka : ", numbers)
	fmt.Println("angka default : ", numbers[4]) //akan return 0 karena tidak ada nilai di index ke 4

	//operasi untuk array
	//1. len : untuk cek panjang array
	fmt.Println("Panjang array numbers: ", len(numbers))

	//bentuk lain array
	var values = [...]int{
		90,
		80,
		100,
	}
	//kalau menggunakan bentuk ini, panjan array adalah otomatis sesuai jumlah data
	//di golang, tidak ada fungsi atau method yang digunakan untuk menghapus data di dalam array

	fmt.Println("Panjang array values: ", len(values))
}
