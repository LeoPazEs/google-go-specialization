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
	// Arbitrator solution
	// Another approach is to guarantee that a philosopher can only pick up both forks or none by introducing an arbitrator to replace circular waiting, e.g., a waiter. In order to pick up the forks, a philosopher must ask permission of the waiter. The waiter gives permission to only one philosopher at a time until the philosopher has picked up both of his forks. Putting down a fork is always allowed. The waiter can be implemented as a mutex. In addition to introducing a new central entity (the waiter), this approach can result in reduced parallelism: if a philosopher is eating and one of his neighbors is requesting the forks, all other philosophers must wait until this request has been fulfilled even if forks for them are still available.
	var wg sync.WaitGroup
	philosToEatChan := make(chan int, simultaneously)

	currPhilo := 0
	firstPhilosopherAte := 0
	for firstPhilosopherAte < eatTimes {
		for i := 0; i < simultaneously; i++ {
			if currPhilo == 0 {
				firstPhilosopherAte += 1
			}
			philosToEatChan <- currPhilo
			currPhilo = (currPhilo + 2) % philosQtd
		}

		wg.Add(simultaneously)

		for i := 0; i < simultaneously; i++ {
			go philos[<-philosToEatChan].eat(&wg)
		}
		wg.Wait()
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
	host(philos, 3, 2)
}
