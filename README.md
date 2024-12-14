# MODUL TIGA BELAS
  ## SOAL 1
  Program di atas adalah program untuk mengurutkan nomor rumah di setiap daerah dalam urutan membesar (ascending). 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Konstanta maxSize, Konstanta ini mendefinisikan ukuran maksimum array (datInt), yang bernilai 100. Konstanta ini digunakan sebagai batas untuk jumlah rumah di setiap daerah.
      - Tipe Data datInt, Sebuah tipe data berupa array dengan ukuran tetap ([maxSize]int) yang digunakan untuk menyimpan daftar nomor rumah di setiap daerah.
      - Fungsi pengurutanMembesar, Fungsi ini mengimplementasikan algoritma selection sort untuk mengurutkan elemen-elemen array dalam urutan membesar (ascending order). Fungsi menerima referensi ke array A (dari tipe datInt) dan jumlah elemen N yang akan diurutkan.
      
   ## Code Explanation
   ```go
	const maxSize = 100
   ```
   Kode di atas adalah kode deklarasi konstanta yang menetapkan nilai tetap 100. Konstanta ini digunakan untuk menentukan batas maksimum ukuran array dalam program, memastikan bahwa jumlah elemen tidak melebihi kapasitas yang ditentukan. 
  
   ```go
	type datInt [maxSize]int
   ```
   Kode di atas adalah kode deklarasi tipe data baru.Tipe ini mendefinisikan datInt sebagai array dengan ukuran tetap sebesar maxSize (yang sebelumnya telah dideklarasikan sebagai konstanta). Elemen-elemen dalam array ini bertipe int.

   ```go
	func pengurutanMembesar(A *datInt, N int) {
		for i := 0; i < N-1; i++ {
			minIdx := i
			for j := i + 1; j < N; j++ {
				if A[j] < A[minIdx] {
					minIdx = j
				}
			}
			A[i], A[minIdx] = A[minIdx], A[i]
		}
	}
   ```
   Kode di atas adalah kode implementasi algoritma selection sort. Algoritma ini digunakan untuk mengurutkan elemen-elemen array dalam urutan membesar (ascending order). Fungsi pengurutanMembesar bekerja dengan cara memilih elemen terkecil dari bagian array yang belum diurutkan dan menukarnya dengan elemen pertama yang belum terurut, hingga seluruh array terurut.
   
   ```go
	var n int
	fmt.Scan(&n)
   ```
   Kode di atas adalah kode untuk membaca input dari pengguna. Kode ini mendeklarasikan variabel n bertipe int dan kemudian menggunakan fungsi fmt.Scan(&n) untuk mengambil nilai input dari pengguna melalui konsol dan menyimpannya dalam variabel 

   ```go
	if n <= 0 || n >= 1000000 || n > maxSize {
		fmt.Println("Jumlah daerah tidak memenuhi syarat")
		return
	}
   ```
   Kode di atas adalah kode validasi input, Kode ini memeriksa apakah nilai n yang dimasukkan memenuhi syarat atau tidak. Jika n lebih kecil dari atau sama dengan 0, lebih besar dari atau sama dengan 1.000.000, atau lebih besar dari nilai maksimum yang ditetapkan oleh konstanta maxSize, maka program akan mencetak pesan kesalahan "Jumlah daerah tidak memenuhi syarat" dan menghentikan eksekusi lebih lanjut dengan return. 

   ```go
	result := make([][]int, n)
   ```
   Kode di atas adalah kode untuk membuat slice dua dimensi. Dengan menggunakan fungsi make, kode ini mendeklarasikan sebuah slice bernama result yang memiliki panjang n dan menyimpan elemen-elemen yang merupakan slice bertipe int. 
   
   ```go
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m) 

		if m <= 0 || m >= 1000000 || m > maxSize {
			fmt.Println("Jumlah rumah tidak memenuhi syarat")
			return
		}

		var houses datInt
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		pengurutanMembesar(&houses, m) 
		result[i] = houses[:m]
	}
   ```
   Kode di atas adalah kode untuk memproses input dan mengurutkan data dalam array. Dalam kode ini, program pertama-tama menerima input jumlah daerah (n). Kemudian, untuk setiap daerah, program membaca jumlah rumah (m) dan memvalidasi apakah jumlah rumah tersebut memenuhi syarat. Jika jumlah rumah tidak valid, program akan mencetak pesan kesalahan dan menghentikan eksekusi. Setelah itu, program membaca nomor rumah untuk setiap daerah dan menyimpannya dalam array houses bertipe datInt. Array ini kemudian diurutkan menggunakan fungsi pengurutanMembesar, yang mengurutkan elemen dalam urutan membesar. Hasil yang telah diurutkan kemudian disalin ke dalam slice dua dimensi result.

   ```go
	for i := 0; i < len(result); i++ {
		houses := result[i]
		for j := 0; j < len(houses); j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(houses[j])
		}
		fmt.Println()
	}
   ```
   Kode di atas adalah kode untuk menampilkan hasil yang telah diproses dan disimpan dalam slice dua dimensi result.
