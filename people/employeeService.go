package people

import (
	"database/sql"
	"time"
	people "van-der-binckes/people/structs"
)

func GetEmployeeById(db *sql.DB, employeeId int) people.Employee {
	query := "SELECT * FROM Employee WHERE employeeId = ?;"
	result, err := db.Query(query, employeeId)
	if err != nil {
		panic(err)
	}

	var (
		name          string
		employeeSince time.Time
		surname       string
		employee      people.Employee
	)

	if result.Next() {
		err := result.Scan(&employeeId, &name, &employeeSince, &surname)
		if err != nil {
			panic(err)
		}

		employee.SetEmployeeId(employeeId)
		employee.SetName(name)
		employee.SetSurname(surname)
		employee.SetEmployeeSince(employeeSince)
	}
	result.Close()
	return employee
}
