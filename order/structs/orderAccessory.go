package structs

import (
	itemStructs "van-der-binckes/items/structs"
)

type OrderAccessory struct {
	order     OrderLine
	accessory itemStructs.Accessory
	amount    int
}

func NewOrderAccessory(accessory itemStructs.Accessory, amount int) OrderAccessory {
	return OrderAccessory{accessory: accessory, amount: amount}
}

func (o *OrderAccessory) Order() OrderLine {
	return o.order
}

func (o *OrderAccessory) SetOrder(order OrderLine) {
	o.order = order
}

func (o *OrderAccessory) Accessory() itemStructs.Accessory {
	return o.accessory
}

func (o *OrderAccessory) SetAccessory(accessory itemStructs.Accessory) {
	o.accessory = accessory
}

func (o *OrderAccessory) Amount() int {
	return o.amount
}

func (o *OrderAccessory) SetAmount(amount int) {
	o.amount = amount
}
