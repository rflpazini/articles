package builder

type Director struct {
	Builder HouseBuilder
}

func (d *Director) Build() {
	d.Builder.SetFoundation()
	d.Builder.SetStructure()
	d.Builder.SetRoof()
	d.Builder.SetInterior()
}

func (d *Director) SetBuilder(b HouseBuilder) {
	d.Builder = b
}
