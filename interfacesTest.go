//trying interface as a type

package main

import "fmt"

type Vehicle interface {
	move ()
}

type Car struct{}

type Aircraft struct {}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}

func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {
	var tesla Vehicle = Car {}
	var boing Vehicle = Aircraft{}

	var boingVhicle Vehicle
	boingVhicle = boing

	tesla.move()
	boing.move()
	boingVhicle.move()
}