package builder

type WoodHouseBuilder struct {
	house House
}

func (b *WoodHouseBuilder) SetFoundation() {
	b.house.Foundation = "Wooden piles"
}

func (b *WoodHouseBuilder) SetStructure() {
	b.house.Structure = "Wooden frame"
}

func (b *WoodHouseBuilder) SetRoof() {
	b.house.Roof = "Wooden shingles"
}

func (b *WoodHouseBuilder) SetInterior() {
	b.house.Interior = "Wooden panels and paint"
}

func (b *WoodHouseBuilder) GetHouse() House {
	return b.house
}
