package main

import "fmt"

func push(st []string, s string) []string {
	st = append(st, s)
	return st
}

func main() {
	var stack []string = []string{"Hi", "Vaiya", "Kemon"}
	fmt.Println(len(stack))
	// stack = append(stack, "Hello")
	// stack = append(stack, "Ruhul")
	stack = push(stack, "ruhul")
	stack = push(stack, "Amin")
	fmt.Println(stack)
	for len(stack) > 0 {
		n := len(stack) - 1
		fmt.Println(stack[n])
		stack = stack[:n]
	}
	fmt.Println(len(stack))

}
