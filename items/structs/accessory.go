package structs

type Accessory struct {
	name  string
	price float64
}

func (a *Accessory) Name() string {
	return a.name
}

func (a *Accessory) SetName(name string) {
	a.name = name
}

func (a *Accessory) Price() float64 {
	return a.price
}

func (a *Accessory) SetPrice(price float64) {
	a.price = price
}
