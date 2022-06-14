package strategy

import "fmt"

func RunStrategy() {
	//  this is the client in strategy pattern
	fmt.Println("***********************************************")
	fmt.Println("Strategy Pattern")
	fmt.Println("***********************************************")
	var m mallardDuck
	// m.flyBehaviour = &flyNoWay{}
	// m.quackBehaviour = &quack{}
	//setting the algorithm dynamically
	m.setFlyBehaviour(&flyNoWay{})
	m.setQuackBehaviour(&quack{})

	m.display()
	m.performFly()
	m.performQuack()
	m.swim()
	fmt.Println("****Changed algorithm for Fly and squeak **********")
	m.setFlyBehaviour(&flyWithWings{})
	m.setQuackBehaviour(&squeak{})

	m.performFly()
	m.performQuack()
	m.setQuackBehaviour(&mute{})
	m.performFly()
	m.performQuack()
}
