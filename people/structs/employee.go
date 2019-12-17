package structs

import "time"

type Employee struct {
	employeeId    int
	name          string
	surname       string
	employeeSince time.Time
}

func (e *Employee) EmployeeId() int {
	return e.employeeId
}

func (e *Employee) SetEmployeeId(employeeId int) {
	e.employeeId = employeeId
}

func (e *Employee) Name() string {
	return e.name
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) Surname() string {
	return e.surname
}

func (e *Employee) SetSurname(surname string) {
	e.surname = surname
}

func (e *Employee) EmployeeSince() time.Time {
	return e.employeeSince
}

func (e *Employee) SetEmployeeSince(employeeSince time.Time) {
	e.employeeSince = employeeSince
}
