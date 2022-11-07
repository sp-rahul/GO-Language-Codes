package main

import "fmt"

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() set {
	s := set{}
	s.m = make(map[string]struct{})
	return s
}

func (s set) Add(value string) set {
	s.m[value] = exists
	return s
}

func (s set) Remove(value string) set {
	delete(s.m, value)
	return s
}

func (s set) Contains(value string) (bool, set) {
	_, c := s.m[value]
	return c, s
}

func main() {
	s := NewSet()

	s.Add("Peter")
	s.Add("David")
	s.Add("wss")
	s.Add("acx")
	s.Add("Acd")
	s.Add("David")

	fmt.Println(s.m)

	// fmt.Println(s.Contains("Peter"))  // True
	// fmt.Println(s.Contains("George")) // False

	s.Remove("David")
	//fmt.Println(s.Contains("David")) // False
	fmt.Println(s)

}
