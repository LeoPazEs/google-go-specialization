package main

import "fmt"

func getMax(vals ...int) int {
	maxInt := vals[0]
	for _, i := range vals {
		if maxInt < i {
			maxInt = i
		}
	}
	return maxInt
}

func main() {
	vSlice := []int{1, 2, 3, 4, 5}
	defer fmt.Println(getMax(vSlice...))
	fmt.Println(getMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15))
}
