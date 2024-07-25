package main

import "fmt"

func main() {
	var m map[string]int     // this is a refence type so behaves like a pointer, runtime erros if try to append for example
	m = make(map[string]int) // now its initialized
	fmt.Println(m)

	literal := map[string]int{"leo": 123, "izabelly": 321}
	fmt.Println(literal)
	for key, val := range literal {
		fmt.Println(key, val)
	}

	var id int
	var p bool
	id, p = literal["leo"] // testing if key exists with p as bool
	fmt.Println("id:", id, "present:", p)

	delete(literal, "joe")
}
