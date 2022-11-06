package main

import "fmt"

func main() {
	mymap := make(map[int]string)
	mymap[2] = "Hello Man"
	mymap[10] = "Ohh no"
	fmt.Println(mymap)

	myList := make([]string, 5)
	myList[1] = "London"
	myList = append(myList, "Babb")
	myList = append(myList, "Go")
	fmt.Println(myList)

	lst := make([]int, 5)
	lst = append(lst, 10)
	fmt.Println(lst)

	myChannel := make(chan int, 5)
	myChannel <- 10
	myChannel <- 12
	fmt.Println(<-myChannel)
}
