package main

import "fmt"

func main() {
	type Person struct {
		name  string
		addr  string
		phone string
	}
	var p1 Person
	p1.name = "leonardo"
	p1.addr = "Maceió"
	p1.phone = "82999999999"
	fmt.Println(p1)
	p2 := new(Person)
	fmt.Println(p2)
	p3 := Person{name: "izabelly", addr: "uniao", phone: "8299999999"}
	fmt.Println(p3)
}
