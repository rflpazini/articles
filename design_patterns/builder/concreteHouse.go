package builder

type ConcreteHouseBuilder struct {
	house House
}

func (b *ConcreteHouseBuilder) SetFoundation() {
	b.house.Foundation = "Concrete, brick, and stone"
}

func (b *ConcreteHouseBuilder) SetStructure() {
	b.house.Structure = "Wood and brick"
}

func (b *ConcreteHouseBuilder) SetRoof() {
	b.house.Roof = "Concrete and reinforced steel"
}

func (b *ConcreteHouseBuilder) SetInterior() {
	b.house.Interior = "Gypsum board, plywood, and paint"
}

func (b *ConcreteHouseBuilder) GetHouse() House {
	return b.house
}
