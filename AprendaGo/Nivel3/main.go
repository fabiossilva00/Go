package main

import (
	"fmt"
	"time"
)

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
	ex9()
}

func ex1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func ex2() {
	initial := 65
	for i := 65; i <= initial+25; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U \n", i)
		}
	}
}

func ex3() {
	born := 1992
	now := time.Now()

	for born <= now.Year() {
		fmt.Println(born)
		born++
	}
}

func ex4() {

	born := 1992
	now := time.Now()

	for {
		fmt.Println(born)
		if born >= now.Year() {
			break
		}
		born++
	}
}

func ex5() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

func ex6() {
	if true {
		fmt.Println("Pulei essa atividade\n")
	}
}

func ex7() {
	if false {
		fmt.Println("Pulei essa atividade\n")
	} else {
		fmt.Println("Pulei essa também\n")
	}
}

func ex8() {
	x := 1
	switch {
	case x == 1:
		fmt.Println("One")
	case x == 2:
		fmt.Println("Two")
	default:
		fmt.Println("None")
	}
}

func ex9() {
	sport := "skate"

	switch sport {
	case "basket":
		fmt.Println("Basket")
	case "soccer":
		fmt.Println("Soccer")
	case "skate":
		fmt.Println("Skate")
	}
}
