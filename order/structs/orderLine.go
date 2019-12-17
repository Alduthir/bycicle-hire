package structs

import (
	"time"
	structs2 "van-der-binckes/items/structs"
	"van-der-binckes/people/structs"
)

type OrderLine struct {
	bycicle                  structs2.Bycicle
	customer                 structs.Customer
	employee                 structs.Employee
	startDate                time.Time
	days                     int
	totalPrice               float64
	orderAccessoryCollection []structs2.OrderAccessory
}

func (o *OrderLine) Bycicle() structs2.Bycicle {
	return o.bycicle
}

func (o *OrderLine) SetBycicle(bycicle structs2.Bycicle) {
	o.bycicle = bycicle
}

func (o *OrderLine) Customer() structs.Customer {
	return o.customer
}

func (o *OrderLine) SetCustomer(customer structs.Customer) {
	o.customer = customer
}

func (o *OrderLine) Employee() structs.Employee {
	return o.employee
}

func (o *OrderLine) SetEmployee(employee structs.Employee) {
	o.employee = employee
}

func (o *OrderLine) StartDate() time.Time {
	return o.startDate
}

func (o *OrderLine) SetStartDate(startDate time.Time) {
	o.startDate = startDate
}

func (o *OrderLine) Days() int {
	return o.days
}

func (o *OrderLine) SetDays(days int) {
	o.days = days
}

func (o *OrderLine) TotalPrice() float64 {
	return o.totalPrice
}

func (o *OrderLine) SetTotalPrice(totalPrice float64) {
	o.totalPrice = totalPrice
}

func (o *OrderLine) OrderAccessoryCollection() []structs2.OrderAccessory {
	return o.orderAccessoryCollection
}

func (o *OrderLine) SetOrderAccessoryCollection(orderAccessoryCollection []structs2.OrderAccessory) {
	o.orderAccessoryCollection = orderAccessoryCollection
}
