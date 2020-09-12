package main

import "fmt"

func main() {

	// Kondisional if else if dan else
	// var point = 2

	// if point == 10 {
	// 	fmt.Println("Selamat! Anda berhasil")
	// } else if point > 5 {
	// 	fmt.Println("Selamat Datang!")
	// } else if point == 4 {
	// 	fmt.Println("Selamat Lagi!")
	// } else {
	// 	fmt.Println("Bagus")
	// }

	// Variabel Temporary pada if else
	// 	var point = 8840.0

	// 	if percent := point / 100; percent >= 100 {
	// 		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	// 	} else if percent >= 70 {
	// 		fmt.Printf("%.1f%s good!\n", percent, "%")
	// 	} else {
	// 		fmt.Printf("%.1f%s not bad!\n", percent, "%")
	// 	}

	// Switch Case
	// var point = 11

	// switch point {
	// case 1:
	// 	fmt.Println("Satu")
	// case 2:
	// 	fmt.Println("Dua")
	// case 3:
	// 	fmt.Println("Tiga")
	// default:
	// 	fmt.Println("Default")
	// }

	// Switch Case banyak kondisi
	// switch point {
	// case 1, 2, 3, 4:
	// 	fmt.Println("Bagus")
	// case 10:
	// 	fmt.Println("Sip")
	// default:
	// 	// fmt.Println("Semangat")
	// 	{
	// 		fmt.Println("Semangat")
	// 		if point == 11 {
	// 			fmt.Println("Combine Switch and If")
	// 		}
	// 		fmt.Println("Mantap")
	// 	}
	// }

	// Switch Case dengan Style If else
	// switch {
	// case point == 2:
	// 	fmt.Println("Keren")
	// case point == 10 || point == 11:
	// 	fmt.Println("Mantap")
	// default:
	// 	fmt.Println("Siapp")
	// }

	// fallthrough adalah memaksa case di go untuk membaca case setelahnya walaupun itu benar atau salah
	// fallthrough cuma bisa digunkaan di swtich
	// switch {
	// case point == 2:
	// 	fmt.Println("Bagus")
	// case (point > 10) && (point < 12):
	// 	fmt.Println("Selamat")
	// 	fallthrough
	// case point < 2:
	// 	fmt.Println("Learn more")
	// default:
	// 	{
	// 		fmt.Println("Belajar Lagi")
	// 		fmt.Println("Tetap Semangat")
	// 	}
	// }

	// Seleksi Bersarang
	var point = 8

	if point > 5 {

		switch {
		case point == 5:
			fmt.Println("Lima")
		case point == 8:
			fmt.Println("Delapan")
			fallthrough
		default:
			{
				if point == 10 {
					fmt.Println("Selamat Anda Berhasil!")
				} else {
					fmt.Println("Don't be sad")
				}
			}

		}
	} else {
		switch point {
		case 1, 2, 3, 4, 5:
			fmt.Println("Angka dibawah 5")
		case 6, 7, 8, 9:
			fmt.Println("Match number 10")
		default:
			{
				fmt.Println("Masuk Default")
				fmt.Println("Kamu Pasti Bisa!")
			}
		}
	}

}
