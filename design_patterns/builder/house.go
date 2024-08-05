package builder

type House struct {
	Foundation string
	Structure  string
	Roof       string
	Interior   string
}

type HouseBuilder interface {
	SetFoundation()
	SetStructure()
	SetRoof()
	SetInterior()
	GetHouse() House
}
