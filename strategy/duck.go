package main

import "fmt"

/*
	duck is the context in the strategy pattern
*/
type duck struct {
	flyBehaviour   flyBehaviour
	quackBehaviour quackBehaviour
}

func (d *duck) swim() {
	fmt.Println("Haha i am swimming")
}

func (d *duck) performQuack() {
	d.quackBehaviour.quack()
}

func (d *duck) performFly() {
	d.flyBehaviour.fly()
}

func (d *duck) setFlyBehaviour(f flyBehaviour) {
	d.flyBehaviour = f
}

func (d *duck) setQuackBehaviour(q quackBehaviour) {
	d.quackBehaviour = q
}
