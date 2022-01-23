package main

import "fmt"

// flyBehaviour is the strategy interface in the strategy pattern
type flyBehaviour interface {
	fly()
}

// flyWithWings, flyNoWay these are concrete strategies
type flyWithWings struct{}
type flyNoWay struct{}

func (w *flyWithWings) fly() {
	fmt.Println("This duck flies with wings")
}

func (n *flyNoWay) fly() {
	fmt.Println("This duck can't fly :(")
}
