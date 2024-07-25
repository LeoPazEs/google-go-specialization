package main

import "fmt"

func main() {
	x := [...]int{1, 2, 3, 4}
	for i, v := range x {
		fmt.Printf("ind %d, val %d", i, v)
	}
	// slices are pointers
	x1 := x[0:3]
	x1[0] = 0
	fmt.Println(x[0])
	// manually creating a slice == dynamically-sized arrays.
	// sli := make([]int, 0, 100)
	var sli []int
	sli = append(sli, 5)
	sli = append(sli, 6)
}
