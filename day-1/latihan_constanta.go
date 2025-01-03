package main

import "fmt"

func main() {
	const gajiPokok = 5000000
	const tunjangan = 1000000

	var insentif int = 20000
	hasil := gajiPokok + tunjangan
	hasil += insentif
	fmt.Println(hasil)
}
