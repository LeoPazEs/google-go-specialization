package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func divideSlice(intSlice []int, slicesPart int) [][]int {
	intSlices := [][]int{}

	sliceLength := len(intSlice)
	slicePartSize := sliceLength / slicesPart
	left := sliceLength % slicesPart

	for i, j := 0, 0; j < slicesPart-1; i, j = i+slicePartSize, j+1 {
		intSlices = append(intSlices, intSlice[i:i+slicePartSize])
	}
	intSlices = append(intSlices, intSlice[sliceLength-slicePartSize-left:sliceLength])

	return intSlices
}

func makeSlice(integers string, splitString string) []int {
	splitIntegers := strings.Split(integers, splitString)
	integersSlice := []int{}
	for _, v := range splitIntegers {
		integer, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		integersSlice = append(integersSlice, integer)

	}
	return integersSlice
}

func sortSlices(intSlices [][]int) [][]int {
	var wg sync.WaitGroup
	wg.Add(len(intSlices))
	for _, v := range intSlices {
		go func(intSlice []int) {
			defer wg.Done()
			fmt.Printf("Goroutine sorting: %v \n", v)
			slices.Sort(v)
		}(v)
	}
	wg.Wait()
	return intSlices
}

func mergeSortedSlices(intSlices [][]int) []int {
	var result []int
	for _, slice := range intSlices {
		result = mergeTwoSortedSlices(result, slice)
	}
	return result
}

func mergeTwoSortedSlices(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

func main() {
	var integers string
	fmt.Println("Integers separeted by ',':")
	fmt.Scan(&integers)
	intSlice := makeSlice(integers, ",")

	intSlices := divideSlice(intSlice, 4)
	sortSlices(intSlices)
	fmt.Println(mergeSortedSlices(intSlices))
}
