package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Swap(numbers []int, firstIndex int, secondIndex int) {
	var temp int
	temp = numbers[secondIndex]

	numbers[secondIndex] = numbers[firstIndex]
	numbers[firstIndex] = temp
}

func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < (len(numbers)-1)-i; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j, j+1)
			}
		}
	}
}

func inputNumbersConverter(inputNumbers *string) []int {
	numbersInput := strings.Split(*inputNumbers, ",")
	var numbers []int
	for _, i := range numbersInput {
		number, _ := strconv.Atoi(i)
		numbers = append(numbers, number)
	}
	return numbers
}

func main() {
	var inputNumbers string
	fmt.Println("Write numbers separeted by ',' to sort:")
	fmt.Scan(&inputNumbers)
	var numbers []int
	numbers = inputNumbersConverter(&inputNumbers)
	BubbleSort(numbers)
	fmt.Println(numbers)
}
