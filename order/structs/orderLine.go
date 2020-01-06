package structs

import (
	"time"
	itemStructs "van-der-binckes/items/structs"
	peopleStructs "van-der-binckes/people/structs"
)

type OrderLine struct {
	orderLineId              int
	bicycle                  itemStructs.Bicycle
	customer                 peopleStructs.Customer
	employee                 peopleStructs.Employee
	startDate                time.Time
	days                     int
	totalPrice               float64
	orderAccessoryCollection []OrderAccessory
}

func (o *OrderLine) OrderLineId() int {
	return o.orderLineId
}

func (o *OrderLine) SetOrderLineId(orderLineId int) {
	o.orderLineId = orderLineId
}

func (o *OrderLine) Bicycle() itemStructs.Bicycle {
	return o.bicycle
}

func (o *OrderLine) SetBicycle(bicycle itemStructs.Bicycle) {
	o.bicycle = bicycle
}

func (o *OrderLine) Customer() peopleStructs.Customer {
	return o.customer
}

func (o *OrderLine) SetCustomer(customer peopleStructs.Customer) {
	o.customer = customer
}

func (o *OrderLine) Employee() peopleStructs.Employee {
	return o.employee
}

func (o *OrderLine) SetEmployee(employee peopleStructs.Employee) {
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

func (o *OrderLine) OrderAccessoryCollection() []OrderAccessory {
	return o.orderAccessoryCollection
}

func (o *OrderLine) SetOrderAccessoryCollection(orderAccessoryCollection []OrderAccessory) {
	o.orderAccessoryCollection = orderAccessoryCollection
}
