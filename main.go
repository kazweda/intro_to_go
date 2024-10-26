package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func short_variable() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

func Multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)
	short_variable()

	var m = Multiply(2, 3)
	fmt.Printf("m is of type %T %v\n", m, m)
}
