package structs

import "database/sql"

type Customer struct {
	customerId        int
	name              string
	surname           string
	postalCode        string
	houseNumber       int
	houseNumberSuffix sql.NullString
}

func (c *Customer) CustomerId() int {
	return c.customerId
}

func (c *Customer) SetCustomerId(customerId int) {
	c.customerId = customerId
}

func (c *Customer) Name() string {
	return c.name
}

func (c *Customer) SetName(name string) {
	c.name = name
}

func (c *Customer) Surname() string {
	return c.surname
}

func (c *Customer) SetSurname(surname string) {
	c.surname = surname
}

func (c *Customer) PostalCode() string {
	return c.postalCode
}

func (c *Customer) SetPostalCode(postalCode string) {
	c.postalCode = postalCode
}

func (c *Customer) HouseNumber() int {
	return c.houseNumber
}

func (c *Customer) SetHouseNumber(houseNumber int) {
	c.houseNumber = houseNumber
}

func (c *Customer) HouseNumberSuffix() sql.NullString {
	return c.houseNumberSuffix
}

func (c *Customer) SetHouseNumberSuffix(houseNumberSuffix sql.NullString) {
	c.houseNumberSuffix = houseNumberSuffix
}