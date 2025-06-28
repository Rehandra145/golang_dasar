package main

import "fmt"

func main() {
	var nilaiAkhir = 80
	var nilaiAbsensi = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusNilaiAbsensi = nilaiAbsensi > 80
	var lulus = lulusNilaiAkhir && lulusNilaiAbsensi
	fmt.Println("Lulus : ", lulus)
}
