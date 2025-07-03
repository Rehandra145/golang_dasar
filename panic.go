package main

import "fmt"

/*
panic function adalah function yang bisa dipanggil untuk menghentikan program,
panic function biasanya dipanggil ketika terjadi panic saat program berjalan
saat panic function dipanggil, program akan terhenti, namun defer ttpp dieksekusi
*/

func endApp() {
	fmt.Println("end app")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ups")
	}

}
func main() {
	runApp(false)
}
