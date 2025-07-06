package database

var connection string

/*
	function init adalah function yang pasti akan dieksekusi ketika sebuah package dipangil.
	function ini akses modifier nya adalah private, tetapi dia pasti akan dieksekusi.
*/

func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
