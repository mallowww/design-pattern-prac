package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// main meth - instance struct
func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

// sub meth - config fields for client requirement
func (b *NormalBuilder) setWindowType() {
	b.windowType = "Window Glass"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

// from sub to main - return object House
func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
