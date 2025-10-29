package main

import "fmt"

func main () {
	const NMAX int = 10
	
	var A [NMAX]int
	
	var i,n,max int
	
	fmt.Scan(&n)
	
	if n > NMAX {
		n = NMAX
	}
	
	for i = 0; i <= n-1; i++ {
		fmt.Scan(&A[i])
	}
	
	max = 0
	
	for i = 1; i <= n-1; i++ {
		if A[i] > A[max] {
			max = i
		}
	}
	
	fmt.Print(A[max], max)
}