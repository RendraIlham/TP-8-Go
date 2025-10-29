package main

import "fmt"

const NMAX int = 20

type tabInt [NMAX]int

func main () {
	var data tabInt
	var nData int
	
	baca(&data, &nData)
	
	cetakElemen(data, nData)
	
	cetakInfo(data, nData)
}

func baca(A *tabInt, n *int) {
	
	var x int
	
	*n = 0
	
	for *n < NMAX {
		fmt.Scan(&x)
		
		if x <= 0 {
			break
		}
		
		(*A)[*n] = x
		*n = *n + 1
	}
}

func cetakElemen(A tabInt, n int) {
	fmt.Print("Elemen array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	
	fmt.Println()
}

func maksimum(A tabInt, n int) int {
	max := 0
	
	for i := 1; i < n; i++ {
		if A[i] > A[max] {
			max = i 
		}
	}
	
	return A[max]
}

func minimum(A tabInt, n int) int {
	min := 0
	
	for i := 1; i <= n-1; i++ {
		if A[i] < A[min] {
			min = i
		}
	}
	
	return A[min]
}

func cetakInfo(A tabInt, n int) {
	fmt.Println("Nilai maksimum: ", maksimum(A,n))
	fmt.Println("Nilai minimum: ", minimum(A,n))
	fmt.Println("Banyak elemen: ", n)
}