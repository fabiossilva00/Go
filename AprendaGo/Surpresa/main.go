package main

import "fmt"

func main() {
	clock()
	ascii()
	akgi()
}

func clock() {

	for i := 1; i <= 12; i++ {
		fmt.Println("Hora ", i)
		for j := 1; j <= 60; j++ {
			fmt.Print(j, "\t")
			// for k := 1; k <= 60; k++ {
			// 	fmt.Print("Seconds: ", k)
			// }
		}
		fmt.Println(" ")
	}
}

func ascii() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v \n", string(i))
	}
}
