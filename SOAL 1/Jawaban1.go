package main

import "fmt"

const MAX = 100

type arrayInt [MAX]int

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

func checkDataOrder(T arrayInt, n int) string {
	for i := 1; i < n-1; i++ {
		if T[i+1]-T[i] != T[i]-T[i-1] {
			return "Data berjarak tidak tetap"
		}
	}
	return "Data berjarak tetap"
}

func main() {
	var input arrayInt
	var n, x int

	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")
	fmt.Print("Input: ")
	
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

	fmt.Println("Sebelum pengurutan:", input[:n])

	insertionSort(&input, n)

	fmt.Println("Setelah pengurutan:", input[:n])

	status := checkDataOrder(input, n)
	fmt.Println(status)
}
