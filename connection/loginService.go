package connection

import (
	"database/sql"
	"fmt"
	"time"
	"van-der-binckes/people/structs"
)

// Requests a username and password. Then tries to use those to fetch the employee trying to login.
func Login(db *sql.DB) structs.Employee {
	userName := enterUserName()
	password := enterPassword()

	return fetchEmployee(db, userName, password)
}

// Simply requests a username string and returns it
func enterUserName() string {
	fmt.Println("Wat is uw gebruikersnaam?")

	var userName string
	fmt.Scanf("%s", &userName)

	return userName
}

// Takes the entered password string and attempts to format it to a Time object. If this fails, reruns the function.
func enterPassword() time.Time {
	fmt.Println("Wat is uw wachtwoord?")
	var password string
	fmt.Scanf("%s", &password)

	date, err := time.Parse("2006-01-02", password)
	if err != nil {
		fmt.Println("Het door u ingevoerde wachtwoord in ongeldig, probeer het opnieuw.")
	}

	return date
}

// Attempts to fetch an Employee with the given username and password. If none is found, returns to login to try again.
func fetchEmployee(db *sql.DB, userName string, password time.Time) structs.Employee {
	query := "SELECT employeeId, name, surname, employee_since FROM Employee WHERE name = ? AND employee_since = ? LIMIT 1;"
	result, err := db.Query(query, userName, password)
	if err != nil {
		panic(err)
	}

	var (
		employeeId    int
		name          string
		surname       string
		employeeSince time.Time
		employee      structs.Employee
	)

	if result.Next() {
		err := result.Scan(&employeeId, &name, &surname, &employeeSince)
		if err != nil {
			panic(err)
		}

		employee.SetEmployeeId(employeeId)
		employee.SetName(name)
		employee.SetSurname(surname)
		employee.SetEmployeeSince(employeeSince)
	} else {
		fmt.Println("De door u opgegeven combinatie van gebruikersnaam en wachtwoord is onjuist. Probeer opnieuw.")
		return Login(db)
	}
	result.Close()
	return employee
}
