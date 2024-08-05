package factory

type FoodFactory struct{}

func (f FoodFactory) CreateFood(ft FoodType) Food {
	switch ft {
	case PizzaType:
		return &Pizza{}
	case SaladType:
		return &Salad{}
	default:
		return nil
	}
}
