package main

import (
	"fmt"
	"reflect"
)

type IAnimal interface {
	getName() string
	Eat() string
	Move() string
	Speak() string
}

type Animal struct {
	Name       string
	food       string
	locomotion string
	noise      string
}

func (a Animal) getName() string {
	return a.Name
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

type Cow struct {
	Animal
}

type Bird struct {
	Animal
}
type Snake struct {
	Animal
}

func animalFactory(animal string, name string) IAnimal {
	var created_animal IAnimal
	switch animal {
	case "cow":
		created_animal = Cow{Animal{Name: name, food: "grass", locomotion: "walk", noise: "moo"}}
	case "bird":
		created_animal = Bird{Animal{Name: name, food: "worms", locomotion: "fly", noise: "peep"}}
	case "snake":
		created_animal = Snake{
			Animal{Name: name, food: "mice", locomotion: "slither", noise: "hsss"},
		}
	default:
		fmt.Println("Unknown animal type:", animal)
		return nil
	}
	return created_animal
}

type AnimalDB struct {
	Cows   []Cow
	Birds  []Bird
	Snakes []Snake
}

func (aDB *AnimalDB) appendAnimal(a IAnimal) *IAnimal {
	switch a.(type) {
	case Cow:
		aDB.Cows = append(aDB.Cows, a.(Cow))
	case Bird:
		aDB.Birds = append(aDB.Birds, a.(Bird))
	case Snake:
		aDB.Snakes = append(aDB.Snakes, a.(Snake))
	default:
		fmt.Println("Animal type not found.")
		return nil
	}
	return &a
}

func (aDB *AnimalDB) queryAnimal(n string) *IAnimal {
	db := reflect.ValueOf(aDB).Elem()
	for i := 0; i < db.NumField(); i++ {
		animalList := db.Field(i)
		for j := 0; j < animalList.Len(); j++ {
			animal := animalList.Index(j).Interface().(IAnimal)
			if animal.getName() == n {
				return &animal
			}
		}
	}
	fmt.Println("Animal not found.")
	return nil
}

func executeMethod(a IAnimal, method string) {
	switch method {
	case "eat":
		fmt.Println(a.Eat())
	case "speak":
		fmt.Println(a.Speak())
	case "move":
		fmt.Println(a.Move())

	}
}

func main() {
	animals := AnimalDB{Cows: []Cow{}, Birds: []Bird{}, Snakes: []Snake{}}
	var request, method, name string
	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s %s", &request, &name, &method)
		switch request {
		case "newanimal":
			animals.appendAnimal(animalFactory(method, name))
		case "query":
			animalPtr := animals.queryAnimal(name)
			if animalPtr != nil {
				animal := *animalPtr
				executeMethod(animal, method)
			}
		default:
			fmt.Println("Invalid request.")
		}
	}
}
