package main

import (
	"Week2/geo"
	"fmt"
)

func main() {
	location := geo.Coordinates{Latitude: 37.42, Longitude: -122.08}
	fmt.Println("Laatitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)
}
