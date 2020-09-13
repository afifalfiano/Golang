package main

import "fmt"

func main() {
	// var names [4]string

	// names[0] = "Nol"
	// names[1] = "Satu"
	// names[2] = "Dua"
	// names[3] = "Tiga"

	// fmt.Println(names[0], names[1], names[2], names[3])

	// Inisialisasi Nilai Awal Array
	// var fruits = [4]string{"Satu", "Dua", "Tiga", "Empat"} // Gaya Horizontal
	// var fruits = [4]string{
	// 	"Test",
	// 	"Satu",
	// 	"Dua",
	// 	"Tiga",
	// } // gaya vertical

	// fmt.Println("Jumlah elemen \t\t", len(fruits))
	// fmt.Println("isi semua elemen \t\t", fruits)

	//  Inisialisasi array tanpa jumlah array
	// var numbers = [...]int{1, 2, 3, 4, 1, 2, 3, 23}

	// fmt.Println("data Array", numbers)
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	// Array multidimensi
	// var numbers1 = [2][3]int{{1, 2, 3}, {3, 2, 1}}
	// var numbers2 = [3][3]int{[3]int{2, 5, 7}, [3]int{5, 6, 9}, [3]int{2, 3, 1}}

	// fmt.Println("numbers1 \t\t", numbers1)
	// fmt.Println("numbers2 \t\t", numbers2)

	// Perulangan dengan array
	// fruits := [4]string{"Halo", "Nama", "Saya", "Afif"}
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("element %d : %s\n", i, fruits[i])
	// }

	// Perulangan menggunakan for range
	// var fruits = [4]string{"Satu", "Dua", "Tiga", "Empat"}

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// Menggunakan undescore untuk menampung variabel yang tdk digunakan dalam for range
	// var fruits = [4]string{"Satu", "Dua", "Tiga", "Empat"}

	// for _, fruit := range fruits {
	// 	// for fruit := range fruits {
	// 	fmt.Printf("elemen : %s\n", fruit)
	// 	// fmt.Println("elemen", fruit)
	// }

	// Alokasi ELemen array menggunakan keyword make
	var fruits = make([]string, 2)

	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits[0], fruits[1])

}
