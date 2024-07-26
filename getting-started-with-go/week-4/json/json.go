package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Public uppercase start Private lowercase letter
	type Person struct {
		Name  string
		Addr  string
		Phone string
	}

	p1 := Person{Name: "izabelly", Addr: "a st.", Phone: "123"}
	barr, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	fmt.Println(string(barr))

	var p2 Person
	err = json.Unmarshal(barr, &p2)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}
	fmt.Println(p2)
}
