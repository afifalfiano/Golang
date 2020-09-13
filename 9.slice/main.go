package main

import "fmt"

func main() {

	// tipe slice dalam array
	// slice adalah tipe data reference
	// var fruits = []int{1, 2, 3, 3} // biasnaya slice itu alokasi elemen atau dataunya dikosongkan
	// var newFruits = fruits[0:2]    // slice melakukan refernsi dari var fruits dan mengambil nilai dari 0 sampai sebelum 2
	// const test = "Yap"

	// fmt.Println(fruits)
	// fmt.Println(newFruits)

	// mengunakan fungsi len untuk menghitung banyak elemen dalam satu array
	// var fruits = []int{1, 2, 3, 4, 5}

	// fmt.Println(len(fruits))

	// Fungsi Cap
	// Diguakan untuk menghitung libur atau kapasitas maksimum slice.
	// var fruits = []string{"Halo", "satu", "dua"}

	// var newFruits = fruits[0:2]

	// fmt.Println("fruits", fruits)
	// fmt.Println("new Fruits", newFruits)
	// fmt.Println("cap new fruits", cap(newFruits))

	// Fungsi Append
	// digunakan untuk menambahkan elemen paling akhir pada slice.
	// var fruits = []string{"Afif", "Alfiano", "Hermasyah"}
	// var newFruits = fruits[0:2]
	// var appendFruitts = append(newFruits, "Muhammad")

	// fmt.Println(fruits)
	// fmt.Println(newFruits)
	// fmt.Println(appendFruitts)

	// Fungsi Copy
	// Diguakan untuk mengcopy refernsi dari satu slice ke slice lainnya
	// var dst = make([]string, 3)
	// src := []string{"Jeruk", "Mangga", "Apel", "Melon"}
	// n := copy(dst, src)

	// fmt.Println(n)
	// fmt.Println(dst)
	// fmt.Println(src)

	// atau contoh lain
	var buah1 = []string{"Jeruk", "Apel", "Semangka"}
	var buah2 = []string{"Melon", "Pepaya", "Alpukat", "Stroberi"}
	var n = copy(buah2, buah1)
	// format copy(dst, src)
	fmt.Println(buah1)
	fmt.Println(buah2)
	fmt.Println(n)

}
