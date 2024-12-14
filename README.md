# MODUL TIGA BELAS
  ## SOAL 1
  Program di atas adalah sebuah implementasi untuk mengurutkan data menggunakan algoritma insertion sort dan memeriksa apakah data memiliki jarak yang tetap antar elemen setelah diurutkan.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan ( dalam hal ini 'fmt' ).
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Program menggunakan konstanta 'MAX' dengan nilai 100 untuk menentukan ukuran maksimum array.
      - Tipe data 'arrayInt' adalah alias untuk array dengan ukuran tetap MAX, yang berfungsi menyimpan bilangan bulat yang diinputkan oleh pengguna.
      - Fungsi 'insertionSort', Fungsi ini digunakan untuk mengurutkan array secara menaik dengan algoritma insertion sort.
      - Fungsi 'checkDataOrder', Fungsi ini mengevaluasi apakah elemen-elemen dalam array memiliki jarak tetap setelah pengurutan.
      
   ## Code Explanation
   ```go
	const MAX = 100
   ```
   Kode di atas adalah kode deklarasi konstanta yang menetapkan nilai tetap 100. Konstanta ini digunakan untuk menentukan batas maksimum ukuran array dalam program, memastikan bahwa jumlah elemen tidak melebihi kapasitas yang ditentukan. 
  
   ```go
	type arrayInt [MAX]int
   ```
   Kode di atas adalah kode deklarasi tipe data baru.Tipe ini mendefinisikan arrayInt sebagai array dengan ukuran tetap sebesar maxSize (yang sebelumnya telah dideklarasikan sebagai konstanta). Elemen-elemen dalam array ini bertipe int.

   ```go
	func insertionSort(T *arrayInt, n int) {
		var i, j int
		var temp int
	
		for i = 1; i < n; i++ {
			temp = (*T)[i]
			j = i - 1
	
			for j >= 0 && (*T)[j] > temp {
				(*T)[j+1] = (*T)[j]
				j = j - 1
			}
			(*T)[j+1] = temp
		}
	}
   ```
   Kode di atas adalah kode yang mengimplementasikan algoritma insertion sort. Algoritma ini digunakan untuk mengurutkan array secara menaik. Dimulai dengan elemen kedua dalam array, kode ini membandingkan elemen yang sedang diproses dengan elemen-elemen sebelumnya, lalu memindahkan elemen-elemen yang lebih besar ke kanan untuk memberi ruang bagi elemen yang sedang diproses. 
   
   ```go
	func checkDataOrder(T arrayInt, n int) string {
		for i := 1; i < n-1; i++ {
			if T[i+1]-T[i] != T[i]-T[i-1] {
				return "Data berjarak tidak tetap"
			}
		}
		return "Data berjarak tetap"
	}
   ```
   Kode di atas adalah kode yang memeriksa apakah elemen-elemen dalam array memiliki jarak yang tetap antar elemen setelah diurutkan. Fungsi checkDataOrder menerima array T dan jumlah elemen n sebagai parameter. Fungsi ini memulai perbandingan dari elemen kedua hingga elemen sebelum terakhir. 
   
   ```go
	var input arrayInt
	var n, x int
   ```
   Kode di atas adalah deklarasi variabel yang digunakan dalam fungsi main. Variabel input bertipe arrayInt, yang merupakan array dengan ukuran tetap yang dapat menampung bilangan bulat. Array ini akan digunakan untuk menyimpan data yang dimasukkan oleh pengguna. Sementara itu, variabel n digunakan untuk menghitung jumlah elemen yang telah dimasukkan secara valid ke dalam array input, dan variabel x digunakan untuk menyimpan setiap bilangan bulat yang dimasukkan oleh pengguna satu per satu. 

   ```go
	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")
	fmt.Print("Input: ")
   ```
   Kode di atas adalah bagian dari program yang menampilkan pesan kepada pengguna untuk memasukkan bilangan bulat. 
   
   ```go
	for {
		fmt.Scanln(&x, " ")

		if x < 0 { 
			break
		}

		if n < MAX {
			input[n] = x
			n++
		} else {
			fmt.Println("Maksimal data telah tercapai!")
			break
		}
	}
   ```
   Kode di atas adalah bagian dari program yang digunakan untuk menerima input bilangan bulat dari pengguna secara berulang hingga bilangan negatif dimasukkan atau batas maksimum data tercapai.

   ```go
	fmt.Println("Sebelum pengurutan:", input[:n])
   ```
   Kode di atas adalah perintah untuk menampilkan array input sebelum dilakukan pengurutan.

   ```go
	insertionSort(&input, n)
	fmt.Println("Setelah pengurutan:", input[:n])
   ```
   Kode di atas melakukan dua langkah penting untuk menampilkan perubahan array setelah dilakukan pengurutan. Pertama, fungsi insertionSort(&input, n) dipanggil untuk mengurutkan array input yang berisi data yang dimasukkan oleh pengguna. Setelah proses pengurutan selesai, perintah fmt.Println("Setelah pengurutan:", input[:n]) digunakan untuk menampilkan array yang telah terurut. Potongan array input[:n] menunjukkan hanya elemen yang valid.

   ```go
	status := checkDataOrder(input, n)
	fmt.Println(status)
   ```
   Kode di atas berfungsi untuk memeriksa apakah elemen-elemen dalam array yang telah diurutkan memiliki jarak yang tetap antar elemen. Pertama, fungsi checkDataOrder(input, n) dipanggil dengan parameter array input yang sudah diurutkan dan jumlah elemen valid (n). Fungsi ini akan memeriksa selisih antara elemen-elemen yang berurutan dalam array untuk menentukan apakah jaraknya konsisten. Hasil pemeriksaan ini kemudian disimpan dalam variabel status dan dicetak menggunakan fmt.Println(status), memberikan informasi kepada pengguna mengenai pola jarak antar elemen dalam array yang telah diurutkan.


   ## SOAL 2
  Program di atas adalah program manajemen daftar buku yang memungkinkan pengguna untuk memasukkan data buku, mengurutkan buku berdasarkan rating, dan mencari buku dengan rating tertentu. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan ( dalam hal ini 'fmt' ).
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Tipe data Buku digunakan untuk menyimpan informasi mengenai buku, seperti ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating.
      - DaftarBuku adalah tipe alias untuk slice yang menyimpan daftar objek Buku. Ini digunakan untuk menyimpan koleksi buku yang dimasukkan oleh pengguna.
      - Fungsi DaftarkanBuku, Fungsi ini digunakan untuk memasukkan data buku ke dalam daftar. Pengguna diminta untuk mengisi informasi buku seperti ID, judul, penulis, penerbit, jumlah eksemplar, tahun, dan rating.
      - Fungsi CetakTerfavorit, Fungsi ini mencetak buku dengan rating tertinggi dari daftar. Jika ada lebih dari satu buku dengan rating tertinggi, semua buku tersebut akan ditampilkan.
      - Fungsi UrutBuku, Fungsi ini mengurutkan daftar buku berdasarkan rating menggunakan algoritma insertion sort.
      - Fungsi Cetak5Terbaru, Fungsi ini menampilkan lima buku dengan rating tertinggi, memberikan daftar buku favorit berdasarkan penilaian.
      - Fungsi CariBuku, Fungsi ini memungkinkan pengguna untuk mencari buku berdasarkan rating tertentu. Jika ditemukan, detail buku akan ditampilkan.
      - Fungsi CetakDaftarBuku, Fungsi ini menampilkan seluruh daftar buku yang telah dimasukkan oleh pengguna beserta detailnya.
      
   ## Code Explanation
   ```go
	type Buku struct {
		ID        int
		Judul     string
		Penulis   string
		Penerbit  string
		Eksemplar int
		Tahun     int
		Rating    int
	}
   ```
   Kode di atas adalah kode yang mendefinisikan tipe data Buku. Tipe data Buku ini berisi tujuh atribut yang masing-masing memiliki fungsi untuk menyimpan informasi buku. 
  
   ```go
	type DaftarBuku []Buku
   ```
   Kode di atas adalah kode yang mendefinisikan tipe data DaftarBuku sebagai alias untuk tipe data slice yang berisi elemen-elemen bertipe Buku. 

   ```go
	const nMax = 7919
   ```
   Kode di atas adalah kode yang mendefinisikan sebuah konstanta nMax dengan nilai 7919. Konstanta ini digunakan untuk menetapkan batas maksimum jumlah buku yang dapat dimasukkan dalam program. 
   
   ```go
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
   ```
   Kode di atas adalah kode untuk mendefinisikan fungsi DaftarkanBuku yang berfungsi untuk mendaftarkan sejumlah buku ke dalam daftar pustaka. Fungsi ini menerima dua parameter: pustaka, yang merupakan pointer ke DaftarBuku, dan n, yang merupakan jumlah buku yang akan dimasukkan. Di dalam fungsi, pertama-tama, slice pustaka dialokasikan dengan ukuran n menggunakan make, yang memungkinkan penampung data buku sesuai dengan jumlah yang diminta.
   
   ```go
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
   ```
   Kode di atas adalah fungsi CetakTerfavorit yang mencetak buku dengan rating tertinggi dari daftar pustaka. Fungsi ini pertama memeriksa apakah daftar buku kosong. Jika tidak, ia mencari rating tertinggi dengan membandingkan setiap buku. Setelah menemukan rating tertinggi, fungsi mencetak semua buku yang memiliki rating tersebut, menampilkan informasi seperti judul, penulis, penerbit, dan rating. Jika lebih dari satu buku memiliki rating yang sama, semuanya akan dicetak.

   ```go
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
   ```
   Kode di atas adalah fungsi UrutBuku yang mengurutkan daftar buku berdasarkan rating menggunakan algoritma insertion sort. Fungsi ini membandingkan rating buku satu per satu, memindahkan buku dengan rating lebih tinggi ke posisi yang sesuai dalam daftar, hingga seluruh daftar terurut dari yang tertinggi ke terendah.
   
   ```go
	func Cetak5Terbaru(pustaka DaftarBuku, n int) {
		fmt.Println("5 Buku dengan rating tertinggi :")
		for i := 0; i < n && i < 5; i++ {
			fmt.Printf("Judul: %s,		Penulis: %s,		Penerbit: %s,		Rating: %d\n", 
				pustaka[i].Judul, pustaka[i].Penulis, pustaka[i].Penerbit, pustaka[i].Rating)
		}
	}
   ```
   Kode di atas adalah fungsi Cetak5Terbaru yang mencetak 5 buku dengan rating tertinggi dari daftar pustaka. Fungsi ini menampilkan judul, penulis, penerbit, dan rating untuk setiap buku, hingga 5 buku pertama (atau jumlah buku yang ada jika kurang dari 5) setelah daftar diurutkan berdasarkan rating.

   ```go
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
   ```
   Kode di atas adalah fungsi CariBuku yang mencari buku berdasarkan rating tertentu dalam daftar pustaka. Fungsi ini memeriksa setiap buku, dan jika ditemukan buku dengan rating yang sesuai, informasi buku tersebut (judul, penulis, penerbit, tahun, eksemplar, dan rating) akan ditampilkan. Jika tidak ada buku dengan rating tersebut, maka pesan "Tidak ada buku dengan rating seperti itu" akan ditampilkan.

   ```go
	func CetakDaftarBuku(pustaka DaftarBuku, n int) {
		fmt.Println("Daftar Buku :")
		for i := 0; i < n; i++ {
			fmt.Printf("Judul: %s,		Penulis: %s,		Penerbit: %s,		Tahun: %d,		Eksemplar: %d,		Rating: %d\n",
				pustaka[i].Judul, pustaka[i].Penulis, pustaka[i].Penerbit, pustaka[i].Tahun, pustaka[i].Eksemplar, pustaka[i].Rating)
		}
	}
   ```
   Kode di atas adalah fungsi CetakDaftarBuku yang mencetak daftar semua buku yang ada dalam pustaka. Fungsi ini menampilkan informasi setiap buku, termasuk judul, penulis, penerbit, tahun terbit, jumlah eksemplar, dan rating. Fungsi ini akan mengulang dan mencetak data untuk semua buku yang ada, berdasarkan jumlah buku n.

   ```go
	var pustaka DaftarBuku
	var n, r int
   ```
   Kode di atas mendeklarasikan dua variabel, yaitu pustaka yang bertipe DaftarBuku (sebuah slice dari struct Buku), dan n serta r yang bertipe int. Variabel n digunakan untuk menyimpan jumlah buku yang akan dimasukkan, sementara r digunakan untuk menyimpan rating yang ingin dicari dalam daftar buku.

   ```go
	fmt.Print("Masukkan jumlah buku : ")
	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
   ```
   Kode di atas meminta pengguna untuk memasukkan jumlah buku yang akan didaftarkan ke dalam sistem dengan menggunakan fmt.Print untuk menampilkan prompt dan fmt.Scan untuk membaca input dari pengguna ke dalam variabel n. Kemudian, fungsi DaftarkanBuku dipanggil untuk memasukkan data buku sebanyak n ke dalam variabel pustaka, yang merupakan slice dari tipe Buku. Fungsi DaftarkanBuku akan mengisi daftar buku dengan data yang dimasukkan oleh pengguna.

   ```go
	fmt.Println("\n=== Buku Terfavorit ===")
	CetakTerfavorit(pustaka, n)
   ```
   Kode di atas menampilkan pesan "=== Buku Terfavorit ===" menggunakan fmt.Println, kemudian memanggil fungsi CetakTerfavorit dengan argumen pustaka (daftar buku) dan n (jumlah buku). Fungsi CetakTerfavorit akan mencari buku dengan rating tertinggi dari daftar pustaka dan menampilkan informasi tentang buku-buku tersebut. Jika tidak ada buku, fungsi akan menampilkan pesan bahwa tidak ada buku yang terdaftar.

   ```go
	fmt.Println("\n=== Mengurutkan Buku Berdasarkan Rating ===")
	UrutBuku(&pustaka, n)
	CetakDaftarBuku(pustaka, n)
   ```
   Kode di atas pertama-tama menampilkan pesan "=== Mengurutkan Buku Berdasarkan Rating ===" menggunakan fmt.Println. Kemudian, fungsi UrutBuku dipanggil dengan argumen &pustaka (pointer ke daftar buku) dan n (jumlah buku), yang akan mengurutkan buku berdasarkan rating menggunakan algoritma insertion sort. Setelah buku diurutkan, fungsi CetakDaftarBuku dipanggil untuk mencetak daftar buku yang sudah diurutkan, dengan menampilkan informasi lengkap tentang setiap buku.

   ```go
	fmt.Println("\n=== 5 Buku dengan Rating Tertinggi ===")
	Cetak5Terbaru(pustaka, n)
   ```
   Kode di atas akan menampilkan pesan "=== 5 Buku dengan Rating Tertinggi ===" menggunakan fmt.Println. Setelah itu, fungsi Cetak5Terbaru dipanggil dengan argumen pustaka (daftar buku) dan n (jumlah buku). Fungsi ini akan mencetak lima buku dengan rating tertinggi dari daftar yang telah diurutkan, atau sebanyak buku yang tersedia jika jumlah buku kurang dari lima. Setiap buku yang tercetak akan menunjukkan informasi seperti judul, penulis, penerbit, dan rating.
   
   ```go
	fmt.Print("\nMasukkan rating buku yang ingin dicari : ")
	fmt.Scan(&r)
	CariBuku(pustaka, n, r)
   ```
   Kode di atas meminta pengguna untuk memasukkan rating buku yang ingin dicari. Setelah itu, nilai rating tersebut disimpan dalam variabel r melalui fmt.Scan(&r). Fungsi CariBuku kemudian dipanggil dengan argumen pustaka (daftar buku), n (jumlah buku), dan r (rating yang dicari). Fungsi CariBuku akan mencari dan mencetak buku yang memiliki rating yang sama dengan nilai r. Jika ada buku dengan rating tersebut, informasi buku (judul, penulis, penerbit, tahun, eksemplar, dan rating) akan ditampilkan. 
