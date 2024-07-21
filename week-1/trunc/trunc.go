package main

import "fmt"

func main() {
	var float_number float32
	fmt.Println("Float number:")
	fmt.Scan(&float_number)
	var truncated_number int = int(float_number)
	fmt.Println(truncated_number)
}
