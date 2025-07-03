package main

import "fmt"

/* * * defer function adalah function yang bisa dijadwalkan untuk dieksekusi setelah sebuah function
selesai di eksekusi * * */

/* * * * defer akan selalu dieksekusi walapun terjadi eror pada function yang dieksekusi * * * */

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	//pemanggilan defer diawal untuk menghindari eror di function ini
	defer logging()

	fmt.Println("Run application")
}
func main() {
	runApplication()
}
