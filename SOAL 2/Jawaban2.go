package main

import "fmt"


type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku []Buku

const nMax = 7919

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	*pustaka = make(DaftarBuku, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d  :\n", i+1)
		fmt.Print("Id : ")
		fmt.Scan(&(*pustaka)[i].ID)
		fmt.Print("Judul : ")
		fmt.Scan(&(*pustaka)[i].Judul)
		fmt.Print("Penulis : ")
		fmt.Scan(&(*pustaka)[i].Penulis)
		fmt.Print("Penerbit : ")
		fmt.Scan(&(*pustaka)[i].Penerbit)
		fmt.Print("Eksemplar : ")
		fmt.Scan(&(*pustaka)[i].Eksemplar)
		fmt.Print("Tahun : ")
		fmt.Scan(&(*pustaka)[i].Tahun)
		fmt.Print("Rating 1-10 : ")
		fmt.Scan(&(*pustaka)[i].Rating)
		fmt.Println("===========================")
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku.")
		return
	}
	maxRating := pustaka[0].Rating
	for i := 1; i < n; i++ {
		if pustaka[i].Rating > maxRating {
			maxRating = pustaka[i].Rating
		}
	}
	fmt.Println("Buku dengan rating tertinggi :")
	for i := 0; i < n; i++ {
		if pustaka[i].Rating == maxRating {
			fmt.Printf("Judul: %s,		Penulis: %s,		Penerbit: %s,		Rating: %d\n", 
				pustaka[i].Judul, pustaka[i].Penulis, pustaka[i].Penerbit, pustaka[i].Rating)
		}
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].Rating < key.Rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Buku dengan rating tertinggi :")
	for i := 0; i < n && i < 5; i++ {
		fmt.Printf("Judul: %s,		Penulis: %s,		Penerbit: %s,		Rating: %d\n", 
			pustaka[i].Judul, pustaka[i].Penulis, pustaka[i].Penerbit, pustaka[i].Rating)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	found := false
	fmt.Println("Buku ditemukan : ")
	for i := 0; i < n; i++ {
		if pustaka[i].Rating == r {
			fmt.Printf("Judul: %s,		Penulis: %s,		Penerbit: %s,		Tahun: %d,		Eksemplar: %d,		Rating: %d\n", 
				pustaka[i].Judul, pustaka[i].Penulis, pustaka[i].Penerbit, pustaka[i].Tahun, pustaka[i].Eksemplar, pustaka[i].Rating)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu.")
	}
}

func CetakDaftarBuku(pustaka DaftarBuku, n int) {
	fmt.Println("Daftar Buku :")
	for i := 0; i < n; i++ {
		fmt.Printf("Judul: %s,		Penulis: %s,		Penerbit: %s,		Tahun: %d,		Eksemplar: %d,		Rating: %d\n",
			pustaka[i].Judul, pustaka[i].Penulis, pustaka[i].Penerbit, pustaka[i].Tahun, pustaka[i].Eksemplar, pustaka[i].Rating)
	}
}

func main() {
	var pustaka DaftarBuku
	var n, r int

	fmt.Print("Masukkan jumlah buku : ")
	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)

	fmt.Println("\n=== Buku Terfavorit ===")
	CetakTerfavorit(pustaka, n)

	fmt.Println("\n=== Mengurutkan Buku Berdasarkan Rating ===")
	UrutBuku(&pustaka, n)
	CetakDaftarBuku(pustaka, n)

	fmt.Println("\n=== 5 Buku dengan Rating Tertinggi ===")
	Cetak5Terbaru(pustaka, n)

	fmt.Print("\nMasukkan rating buku yang ingin dicari : ")
	fmt.Scan(&r)
	CariBuku(pustaka, n, r)
}
