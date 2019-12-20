package structs

type Bycicle struct {
	name        string
	bycicleType string
	price       float64
	amount      int
}

func (b *Bycicle) Name() string {
	return b.name
}

func (b *Bycicle) SetName(name string) {
	b.name = name
}

func (b *Bycicle) BycicleType() string {
	return b.bycicleType
}

func (b *Bycicle) SetBycicleType(bycicleType string) {
	b.bycicleType = bycicleType
}

func (b *Bycicle) Price() float64 {
	return b.price
}

func (b *Bycicle) SetPrice(price float64) {
	b.price = price
}

func (b *Bycicle) Amount() int {
	return b.amount
}

func (b *Bycicle) SetAmount(amount int) {
	b.amount = amount
}
