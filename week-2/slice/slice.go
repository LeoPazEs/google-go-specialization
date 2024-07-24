package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)

	var data_in string
	fmt.Println("Insert a integer or 'X' to exit: ")
	fmt.Scan(&data_in)
	for data_in != "X" {
		data_in_int, err := strconv.Atoi(data_in)
		_ = err
		slice = append(slice, data_in_int)

		sort.Ints(slice)
		fmt.Println(slice)

		fmt.Println("Insert a integer or 'X' to exit: ")
		fmt.Scan(&data_in)
	}
}
