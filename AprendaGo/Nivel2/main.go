package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
}

func ex1() {
	n := 101
	fmt.Printf("%b\t %d\t %#x\n", n, n, n)
}

func ex2() {
	a := (1 == 1)
	b := (1 != 1)
	c := (1 <= 1)
	d := (1 < 1)
	e := (1 >= 1)
	f := (1 > 1)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)

}

const a = 1
const b int = 1

func ex3() {
	fmt.Println(a)
	fmt.Println(b)
}

func ex4() {
	var i int = 2
	fmt.Printf("%d\t %b\t %#x\n", i, i, i)
	d := i << 1
	fmt.Printf("%d\t %b\t %#x\n", d, d, d)
}

func ex5() {
	s := `Here`
	fmt.Println(s)
}

const (
	_ = iota + 2020
	y1
	y2
	y3
	y4
)

func ex6() {
	fmt.Println(y1, y2, y3, y4)
}
