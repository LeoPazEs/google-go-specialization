package main

import (
	"fmt"
	"reflect"
)

type IAnimal interface {
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
	Cows   []IAnimal
	Birds  []IAnimal
	Snakes []IAnimal
}

func (aDB *AnimalDB) appendAnimal(a IAnimal) *[]IAnimal {
	var new_slice *[]IAnimal
	switch a.(type) {
	case Cow:
		aDB.Cows = append(aDB.Cows, a)
		new_slice = &aDB.Cows
	case Bird:
		aDB.Birds = append(aDB.Birds, a)
		new_slice = &aDB.Birds
	case Snake:
		aDB.Snakes = append(aDB.Snakes, a)
		new_slice = &aDB.Snakes
	default:
		fmt.Println("Animal type not found.")
		return nil
	}
	return new_slice
}

func (aDB *AnimalDB) queryAnimal(n string) *Animal {
	db := reflect.ValueOf(aDB).Elem()
	for i := 0; i < db.NumField(); i++ {
		animalList := db.Field(i)
		for j := 0; j < animalList.Len(); j++ {
			animal := animalList.Index(j).Interface().(IAnimal)
			if animal.(Animal).Name == n {
				return &animal
			}
		}
	}
	fmt.Println("Animal not found.")
	return nil
}

func main() {
	animals := AnimalDB{Cows: []IAnimal{}, Birds: []IAnimal{}, Snakes: []IAnimal{}}
	request, method, name := "newanimal", "cow", "izabelly"
	for {
		// fmt.Print("> ")
		// fmt.Scanf("%s %s %s", &request, &name, &method)
		switch request {
		case "newanimal":
			animals.appendAnimal(animalFactory(method, name))
		case "query":
			animals.queryAnimal(name).Eat()
		default:
			fmt.Println("Invalid request.")
		}
	}
}
