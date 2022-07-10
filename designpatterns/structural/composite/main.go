package main

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

type Swimmer interface {
	Swim()
}
type Trainer interface {
	Train()
}
type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

func Swim() {
	fmt.Println("Swimming!")
}

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p Parent) int {
	return p.SomeField
}

func main() {
	swimmerA := CompositeSwimmerA{
		MySwim: Swim,
	}
	swimmerA.MyAthlete.Train()
	swimmerA.MySwim()

	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()

	swimmerB := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmerB.Train()
	swimmerB.Swim()

	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
			Left:      nil,
		},
		Left: &Tree{4, nil, nil},
	}
	fmt.Println(root.Right.Right.LeafValue)

	son := Son{}
	fmt.Println(GetParentField(son.P))
}
