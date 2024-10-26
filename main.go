package main

import (
	"fmt"
	"math"
)

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

func sum_loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("result of sum loop:", sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func fizzbuzz() {
	i := 1
	for i <= 30 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
		i++
	}
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)
	short_variable()

	var m = Multiply(2, 3)
	fmt.Printf("m is of type %T %v\n", m, m)

	sum_loop()
	fmt.Println(sqrt(2), sqrt(-4))

	fizzbuzz()
}
