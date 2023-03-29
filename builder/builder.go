package main

type IBuilder interface {
    setWindowType()
    setDoorType()
    setNumFloor()
    getHouse() House
}

func getBuilder(builderType string) IBuilder {
	switch builderType {
	case "normal":
		return &NormalBuilder{}
	case "igloo":
		return &IglooBuilder{}
	default:
		return nil
	}
}