package main

import "fmt"

func main() {
	//slice adalah potongan data dari array
	//dslice mirip dengan array, yang membedakkan adalah slice tidak memiliki panjang tetap
	//hal ini berarti kita bisa menambah atau mengurangi data di dalam slice
	//slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di array

	//slice memiliki 3 bagian yang penting dan harus diketahui
	//1. pointer : adalah alamat dari data pertama di dalam slice
	//2. length : adalah jumlah data yang ada di dalam slice
	//3. capacity : adalah jumlah data yang bisa diakses di dalam slice, dimulai dari pointer sampai akhir array

	//tabel :
	//array[low:high] => membuat slice dari array yang dimulai dari index low sampai index sebelum high (high-1)
	//array[low:] => membuat slice dari array yang dimulai dari index low sampai akhir array
	//array[:high] => membuat slice dari array yang dimulai dari index 0 sampai index sebelum high (high - 1)
	//array[:] => membuat slice dari array yang dimulai dari index 0 sampai akhir array

	//dari sini bisa disimpulkan bahwa slice adalah potongan sebagian atau seluruh data dari array
	//sesuai dengan yang sudah dijelaskan di atas

	names := [...]string{
		"Rehan",
		"Handera",
		"Rehandra",
		"Abit",
		"Akul",
		"Neca",
	}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	//bentuk manual deklarasi slice
	var slice5 []string = names[:]
	fmt.Println(slice5)

	//function di slice
	//1. len(slice) : untuk cek panjang slice
	fmt.Println("Panjang slice ke 4 : ", len(slice4))
	//2. cap(slice) : untuk cek kapasitas slice
	fmt.Println("Panjang kapasitas ke 4 : ", cap(slice4))
	//3. append(slice, data) : mmebuat slice baru dengan menambah data ke posisi terakhir slice
	//jika kapasistas slice sudah penuh, maka akan membuat array baru
	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}
	daySlice1 := days[5:]

	//mengubah data di dalam slice, akan mengubah data di array juga
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	//menambahkan data ke array yang sudah penuh
	daySlice2 := append(daySlice1, "Hari Baru") //disini akan membuat array baru yang tidak bisa diakses
	//dia akan membuat days baru dengan kapasitas yang sesuai dengan jumlah data di append di slice
	//kita tidak bisa akes array baru yang ada di buat oleh daySlice2
	daySlice2[0] = "Ups"   //mengubah data di dalam daySlice2 tidak akan mengubah data di daySlice1 ataupun array days karena dia sudah memakai array baru yang diciptakan oleh daySlice2
	fmt.Println(daySlice2) // akan menampilkan "Ups" karena daySlice2 adalah slice baru yang mengakses data dari daySlice1

	//4. make([]typeData, length, capacity) : untuk membuat slice baru dari array baru
	newSlice := make([]string, 2, 5) //membuat slice baru dengan panjang 2 dan kapasitas 5
	newSlice[0] = "Rehan"
	newSlice[1] = "Handera"
	//newSlice[2] = "Rehandra" //akan panic karena panjang slice hanya 2 harusnya pakai append
	newSlice2 := append(newSlice, "Rehandra")
	fmt.Println(newSlice)
	fmt.Println("Panjang newSlice: ", len(newSlice))   //akan menampilkan 2
	fmt.Println("Kapasitas newSlice: ", cap(newSlice)) //akan menampilkan
	fmt.Println(newSlice2)                             //akan menampilkan [Rehan Handera Rehandra]
	newSlice2[0] = "Abit"
	fmt.Println(newSlice2) //akan menampilkan [Rehan Handera Rehandra]
	fmt.Println(newSlice)
	//5. copy(destination, source) : untuk menyalin data dari source ke destination
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println("From Slice: ", fromSlice)
	fmt.Println("To Slice: ", toSlice) //akan menampilkan [Senin Selasa Rabu Kamis Jumat Sabtu Minggu]

	//hati hati dalam membuat array dan membuat slice, karena pada perbedaannya itu mungkin tidak terlalu nampak
	//array
	iniArray := [...]int{
		1, 2, 3, 4, 5,
	}
	//slice
	iniSlice := []int{
		1, 2, 3, 4, 5,
	}

	//perbedaannya adalah di operator [...] dan [] pada deklarasi

	fmt.Println("Ini Array: ", iniArray) // akan menampilkan [1 2 3 4 5]
	fmt.Println("Ini Slice: ", iniSlice) // akan menampilkan [1 2 3 4 5]
}
