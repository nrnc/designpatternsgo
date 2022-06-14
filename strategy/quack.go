package strategy

import "fmt"

// quackBehaviour is the strategy interface in strategy pattern
type quackBehaviour interface {
	quack()
}

//  quack, squeak, mute are the concrete strategies
type quack struct{}
type squeak struct{}
type mute struct{}

func (q *quack) quack() {
	fmt.Println("This duck quacks")
}

func (s *squeak) quack() {
	fmt.Println("This duck squeaks")
}

func (m *mute) quack() {
	fmt.Println("This duck can't quack :(")
}
