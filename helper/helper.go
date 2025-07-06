package helper

var Application = "golang"

/*
	di golang, nama function yang diawali dengan huruf besar = akses modifier publik, tetapi jika
	dimulai dengan huruf kecil, akses modifier = private
*/

func sayGoodBye(name string) string {
	return "Good bye " + name
}
func SayHello(name string) string {
	return "Hello " + name
}
