package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	animals := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	var animal string
	var request string
	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s", &animal, &request)
		chosen_animal, exists := animals[animal]
		if !exists {
			fmt.Println("Animal not found")
		} else {
			if request == "eat" {
				fmt.Println(chosen_animal.Eat())
			} else if request == "move" {
				fmt.Println(chosen_animal.Move())
			} else if request == "speak" {
				fmt.Println(chosen_animal.Speak())
			} else {
				fmt.Println("Request not found")
			}
		}
	}
}
