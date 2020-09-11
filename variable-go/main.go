package main

import "fmt"

func main() {
	// Deklarasi dengan manifest typing menggunakan var
	// var firstName string = "Afif"

	// var lastName string = "Alfiano"
	// reassign variabel lastName
	// lastName = "Hermasyah"
	// fmt.Printf("Halo %s %s!\n", firstName, lastName)

	// type inference tanpa menggunakan variabel jadi dia meyesuaikan isinya
	// nama := "Afif"
	// semester := 3

	// fmt.Println("Haloo! nama saya", nama, "Semester", semester)

	// deklarasi multi variabel
	// var satu, dua, tiga = "halo", 3, false
	// fmt.Println(satu, dua, tiga)

	// Nilai underscore artinya masuk kedalam keranjang dan variabel tidak pernah terpakai
	// _ = "belajar Golang"

	// deklarasi variabel menggunakan pointer New
	namaLengkap := new(string)

	fmt.Println(namaLengkap)
	fmt.Println(*namaLengkap)
}
