package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func get_file_path_input() string {
	var file_path string
	fmt.Println("Path to file:")
	fmt.Scan(&file_path)
	return file_path
}

func open_file(file_path string) *os.File {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	return file
}

func read_names(file *os.File) []Name {
	names := make([]Name, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fname_lname := strings.Split(scanner.Text(), " ")
		names = append(names, Name{Fname: fname_lname[0], Lname: fname_lname[1]})
	}
	return names
}

func print_names(names []Name) {
	for _, value := range names {
		fmt.Printf(value.Fname + " " + value.Lname + "\n")
	}
}

func main() {
	file_path := get_file_path_input()
	file := open_file(file_path)
	defer file.Close()
	names := read_names(file)
	print_names(names)
}
