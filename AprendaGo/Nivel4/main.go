package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}

func ex1() {

	array := [5]int{1, 2, 3, 4, 5}

	for _, v := range array {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", array)
}

func ex2() {

	slice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	for _, v := range slice {
		fmt.Println(v)
	}
	fmt.Printf("%T\n %v\n", slice, cap(slice))

}

func ex3() {

	slice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(slice[:3])
	fmt.Println(slice[5:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[3:8])
	fmt.Println(slice[3 : len(slice)-1])

}
