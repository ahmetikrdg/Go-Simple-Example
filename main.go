package main

import "fmt"

func sortFunc(function func(it []int), a []int) {
	function(a)
}

func insertionSort(b []int) {

	for i := 1; i <= len(b)-1; i++ {
		temp := b[i]
		j := i - 1
		for j >= 0 && b[j] > temp {
			b[j+1] = b[j]
			j--
		}
		b[j+1] = temp
	}
	fmt.Println("Insertion Sort")
	sortShowArray(b)
}

func bubbleSort(a []int) {
	for i := len(a); i > 0; i-- {
		for j := 1; j < i; j++ {
			if a[j-1] > a[j] {
				carry := a[j]
				a[j] = a[j-1]
				a[j-1] = carry
			}
		}
	}
	fmt.Println("Bubble Sort")
	sortShowArray(a)
}

func sortShowArray(show []int) {
	for i := 0; i < len(show); i++ {
		fmt.Printf("%v ", show[i])
	}
	fmt.Println()
}

func factorization(numb int) {
	for i := 0; i < numb; i++ {
		for j := numb; j > i; j-- {
			multiplying := i * j
			if multiplying == numb {
				fmt.Println(i, "x", j, "=", multiplying)
			}
		}
	}
}


func main() {
	numbers := []int{1, 5, 3, 7, 9, 2, 15}
	sortFunc(bubbleSort, numbers)
	sortFunc(insertionSort, numbers)

	fmt.Println("Factorization")
	factorization(1000)
}
