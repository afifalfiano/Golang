package main

import "fmt"

func main() {
	// for

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Angka", i+1)
	// }

	// for hanya dengan argumen
	// var i = 0
	// for i < 5 {
	// 	fmt.Println("Angka", i)

	// 	i++
	// }

	// for tanpa argumen dan kondisi
	// var i = 0
	// for {
	// 	fmt.Println("Angka", i)

	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// menggunakna break dan continue
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }

	// Perulangan bersarang

	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}

	// 	fmt.Println()
	// }

	// Label berguna ketka mengguakan break untuk diluar scope
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {

			if i == 5 {
				break outerLoop
			}

			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

}
