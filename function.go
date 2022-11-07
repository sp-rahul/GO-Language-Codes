package main

import "fmt"

func calc(a, b int) (int, int) {
	add := a + b
	sub := a - b
	return add, sub
}

func main() {
	a := 5
	b := 12
	result, sub := calc(a, b)
	//result := a + b

	fmt.Println(result, sub)
}
