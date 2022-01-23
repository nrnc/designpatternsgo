package main

func main() {
	//  this is the client in strategy pattern
	var m mallardDuck
	// m.flyBehaviour = &flyNoWay{}
	// m.quackBehaviour = &quack{}
	m.setFlyBehaviour(&flyNoWay{})
	m.setQuackBehaviour(&quack{})
	m.display()
	m.performFly()
	m.performQuack()
	m.swim()
	m.setFlyBehaviour(&flyWithWings{})
	m.setQuackBehaviour(&squeak{})
	m.performFly()
	m.performQuack()
	m.setQuackBehaviour(&mute{})
	m.performFly()
	m.performQuack()
}
