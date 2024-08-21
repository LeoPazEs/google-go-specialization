package main

import (
	"fmt"
	"sync"
)

type Chops struct{ sync.Mutex }

type Philo struct {
	id              int
	Name            string
	LeftCS, RightCS *Chops
	ate             int
}

func (p Philo) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	p.LeftCS.Lock()
	p.RightCS.Lock()
	fmt.Printf("%s starting to eat %d.\n", p.Name, p.id)
	fmt.Printf("%s finishing eating %d.\n", p.Name, p.id)
	p.LeftCS.Unlock()
	p.RightCS.Unlock()
}

func host(philos []*Philo, eatTimes int, simultaneously int) {
	philosQtd := len(philos)
	if philosQtd < 3 {
		panic("There are not enough chopsticks on the table.")
	}
	if philosQtd/simultaneously < 2 {
		panic("There are not enough chopsticks on the table for simultaneosly eating quantity.")
	}

	var wg sync.WaitGroup
	totalEat := philosQtd * eatTimes
	for totalEat > 0 {
		wg.Add(simultaneously)
		// How can i prevent the edge case where the philo has already eaten 3 times but is one of the possible to eat
		// channels?
		for i, v := range philos {
		}
		wg.Wait()
		totalEat -= simultaneously
	}
}

func main() {
	philos := []*Philo{
		{id: 1, Name: "Aristotle"},
		{id: 2, Name: "Karl Marx"},
		{id: 3, Name: "Immanuel Kant"},
		{id: 4, Name: "Friedrich Nietzsche"},
		{id: 5, Name: "Plato"},
	}
	chops := []*Chops{new(Chops), new(Chops), new(Chops), new(Chops), new(Chops)}
	for i, v := range philos {
		v.LeftCS = chops[i]
		v.RightCS = chops[(i+1)%5]
	}
}
