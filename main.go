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

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)
}
