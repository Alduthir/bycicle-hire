package structs

import (
	"van-der-binckes/items/structs"
)

type OrderAccessory struct {
	order     OrderLine
	accessory structs.Accessory
	amount    int
}

func (o *OrderAccessory) Order() OrderLine {
	return o.order
}

func (o *OrderAccessory) SetOrder(order OrderLine) {
	o.order = order
}

func (o *OrderAccessory) Accessory() structs.Accessory {
	return o.accessory
}

func (o *OrderAccessory) SetAccessory(accessory structs.Accessory) {
	o.accessory = accessory
}

func (o *OrderAccessory) Amount() int {
	return o.amount
}

func (o *OrderAccessory) SetAmount(amount int) {
	o.amount = amount
}
