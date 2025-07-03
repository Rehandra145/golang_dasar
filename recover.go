package main

import "fmt"

/*
Recover adalah function yang bisa digunakan untuk menangkap data panic
Dengan recover proses panic akan terhent, sehingga program akan tetap berjalan.


Pemakaian recover yang benar adalah memanggilnya di defer function
*/

func endAppRec() {
	fmt.Println("end app")
	message := recover()
	fmt.Println("Terjadi kesalahan", message)
}

func runAppRec(error bool) {
	defer endAppRec()

	if error {
		panic("ups")
	}

}
func main() {
	runAppRec(false)
	fmt.Println("Rehandra")
}
