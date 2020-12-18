package main

import "fmt"

var x2 int
var y2 string
var z2 bool

var x3 int = 42
var y3 string = "James Bond"
var z3 bool = true

type tipo int

var x4 tipo

var y5 int

func main() {
	ex1()
	ex2()
	ex3()
	ex4e5()
}

func ex1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

func ex2() {
	fmt.Println(x2, y2, z2)
	fmt.Printf("%T\n", x2)
	fmt.Printf("%T\n", y2)
	fmt.Printf("%T\n", z2)

}

func ex3() {
	sprint := fmt.Sprintf("%v, %v, %v", x3, y3, z3)
	fmt.Println(sprint)
}

func ex4e5() {
	fmt.Println(x4)
	fmt.Printf("%T\n", x4)

	x4 = 42
	fmt.Println(x4)

	y5 = int(x4)

	fmt.Println(y5)
	fmt.Printf("%T\n", y5)
}
