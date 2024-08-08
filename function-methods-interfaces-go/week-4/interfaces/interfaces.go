package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{ name string }

func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("<noise>")
	} else {
		fmt.Println(d.name)
	}
}

// Matches automatically the interface and the type
// Dynamic Type -> Struct Dynamic Value -> Instance
func main() {
	var s1 Speaker
	d1 := Dog{"Brian"}

	s1 = &d1
	s1.Speak()

	var s2 Speaker
	var d2 *Dog

	s2 = d2
	s2.Speak()
}
