package structs

type Accessory struct {
	accessoryId int
	name  string
	price float64
}

func (a *Accessory) AccessoryId() int {
	return a.accessoryId
}

func (a *Accessory) SetAccessoryId(accessoryId int) {
	a.accessoryId = accessoryId
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
