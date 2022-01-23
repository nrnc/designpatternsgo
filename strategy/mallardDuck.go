package main

import "fmt"

// mallardDuck is different implementation of duck
type mallardDuck struct {
	duck
}

// display displays the duck
func (m *mallardDuck) display() {
	fmt.Println("I am mallard duck")
}
