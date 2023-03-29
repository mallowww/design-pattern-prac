package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buiildHouse()

	fmt.Printf("Normal House Window type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Door type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Num floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buiildHouse()

	fmt.Printf("\nIgloo House Window type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Door type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Num floor: %d\n", iglooHouse.floor)

}
