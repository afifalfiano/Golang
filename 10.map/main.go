package main

import "fmt"

func main() {
	// Map berprisnpikan key value,
	// Modelnnya seperti slice hanya saja mengaksesnya berdasarkan key seperti pada object tpi dalam array
	// map bersifat overwrite jadi jika ada key yg sama dengan value berbeda maka yg terbaru akan overwrite value lama
	// var chicken map[string]int
	// chicken = map[string]int{}
	// chicken["januari"] = 50
	// chicken["februari"] = 100
	// chicken["februari"] = 200

	// fmt.Println("januari ", chicken["januari"])
	// fmt.Println("februari ", chicken["februari"])

	// Inisailisasi langsung dari map
	// var bio = map[string]int{}
	// var bio = make(map[string]int)
	// var bio = *new(map[string]int) // karena new ini menggunakanp pointer maka untuk mengambil nilai gunakan asterisk
	// var chicken = map[string]string{"nama": "Afif", "kerja": "developer"}
	// chicken["rumah"] = "pelemsari"
	// bio["test"] = 2

	// fmt.Println("Nama Saya: ", chicken["nama"])
	// fmt.Println(chicken)

	// Iterasi item map menggunakan for - range

	// var bulan = map[string]int{
	// 	"januari":  1,
	// 	"februari": 2,
	// 	"maret":    3,
	// 	"april":    4,
	// 	"mei":      5,
	// }

	// for key, value := range bulan {
	// 	fmt.Println(key, "  \t:", value)

	// }

	// Delete item pada map
	// var chicken = map[string]int{"januari": 1, "februari": 2}

	// fmt.Println(len(chicken))
	// fmt.Println(chicken)

	// delete(chicken, "januari")

	// fmt.Println(len(chicken))
	// fmt.Println(chicken)

	// Mendeteksi keberadaan item dengan key tertentu
	// var chicken = map[string]int{"jaunari": 1, "februari": 2, "mei": 3}
	// var value, isExist = chicken["mei"]

	// if isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Item is not exist")
	// }

	// Kombinasi menggunakan map dan slice
	// var dataMahasiswa = []map[string]string{
	// 	map[string]string{"nama": "Muhammad Afif", "jurusan": "Informatika"},

	// 	map[string]string{"nama": "Muhammad Alfiano", "jurusan": "Informatika"},

	// 	map[string]string{"nama": "Muhammad Hermasyah", "jurusan": "Informatika"},
	// }

	// for _, data := range dataMahasiswa {
	// 	fmt.Println(data["nama"], data["jurusan"])
	// }

	// Cara lebih singkat dalam konbinasi slice dan map
	var dataMahasiswa = []map[string]string{
		{"nama": "Muhammad Afif Alfiano", "jurusan": "informatika"},
		{"nama": "Muhammad Afif Alfiano", "jurusan": "informatika"},
		{"kuliah": "STMIK AKAKOM", "qutotes": "Belajar"},
	}

	for _, data := range dataMahasiswa {
		fmt.Println(data["nama"], data["kuliah"])
	}

}
