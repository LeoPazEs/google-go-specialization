package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	person_map := map[string]string{"name": "", "address": ""}

	scanner := bufio.NewScanner(os.Stdin)
	var input_name string
	fmt.Println("Person name:")
	scanner.Scan()
	input_name = scanner.Text()
	person_map["name"] = input_name

	var input_address string
	fmt.Println("Person address:")
	scanner.Scan()
	input_address = scanner.Text()
	person_map["address"] = input_address

	person_json, err := json.Marshal(person_map)
	_ = err
	fmt.Println(string(person_json))
}
