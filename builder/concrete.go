package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// main meth
func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

// sub meth
func (b *NormalBuilder) setWindowType() {
	b.windowType = "Window Glass"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

// from sub to main
func (b *NormalBuilder) getHouse() House {
	return House{
		doorType: b.doorType,
		windowType: b.windowType,
		floor: b.floor,
	}
}