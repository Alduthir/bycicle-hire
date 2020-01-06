package structs

type Bicycle struct {
	bicycleId   int
	name        string
	bicycleType string
	price       float64
	amount      int
}

func (b *Bicycle) BicycleId() int {
	return b.bicycleId
}

func (b *Bicycle) SetBicycleId(bicycleId int) {
	b.bicycleId = bicycleId
}

func (b *Bicycle) Name() string {
	return b.name
}

func (b *Bicycle) SetName(name string) {
	b.name = name
}

func (b *Bicycle) BicycleType() string {
	return b.bicycleType
}

func (b *Bicycle) SetBicycleType(bicycleType string) {
	b.bicycleType = bicycleType
}

func (b *Bicycle) Price() float64 {
	return b.price
}

func (b *Bicycle) SetPrice(price float64) {
	b.price = price
}

func (b *Bicycle) Amount() int {
	return b.amount
}

func (b *Bicycle) SetAmount(amount int) {
	b.amount = amount
}
