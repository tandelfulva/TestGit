package main

import "fmt"

type Trunk string

func (t Trunk) Accelerate() {
	fmt.Println("Speeding up")
}

func (t Trunk) Brake() {
	fmt.Println("Stopping")
}

func (t Trunk) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Trunk) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()

	trunk, ok := vehicle.(Trunk)
	if ok {
		trunk.LoadCargo("test cargo")
	}
}

func main() {
	TryVehicle(Trunk("Fnord F180"))
}
