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

func (p Philo) eat(host chan int, done chan int, eatTimes int) {
	for i := 0; i < eatTimes; i++ {
		<-host
		p.LeftCS.Lock()
		p.RightCS.Lock()
		fmt.Printf("%s starting to eat %d.\n", p.Name, p.id)
		fmt.Printf("%s finishing eating %d.\n", p.Name, p.id)
		p.RightCS.Unlock()
		p.LeftCS.Unlock()
	}
	done <- p.id
}

func host(hostChan chan int) {
	for {
		hostChan <- 1
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

	hostChan := make(chan int, 2)
	doneChan := make(chan int, len(philos))
	go host(hostChan)
	for _, v := range philos {
		go v.eat(hostChan, doneChan, 3)
	}
	for i := 0; i < len(philos); i++ {
		<-doneChan
	}
}
