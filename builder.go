package main

import "fmt"

type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buiildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func getBuilder(builderType string) Builder {
	switch builderType {
	case "normal":
		return &NormalBuilder{}
	case "igloo":
		return &IglooBuilder{}
	default:
		return nil
	}
}

// NormalBuilder is a builder for building normal houses
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// setWindowType sets the window type of the house
func (b *NormalBuilder) setWindowType() {
	b.windowType = "Window Glass"
}

// setDoorType sets the door type of the house
func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

// setNumFloor sets the number of floors of the house
func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

// getHouse returns the built house
func (b *NormalBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}

// IglooBuilder is a builder for building igloo houses
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// setWindowType sets the window type of the house
func (b *IglooBuilder) setWindowType() {
	b.windowType = "Ice Window"
}

// setDoorType sets the door type of the house
func (b *IglooBuilder) setDoorType() {
	b.doorType = "Ice Door"
}

// setNumFloor sets the number of floors of the house
func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

// getHouse returns the built house
func (b *IglooBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}

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
