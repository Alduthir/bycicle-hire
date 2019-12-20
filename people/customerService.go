package people

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	people "van-der-binckes/people/structs"
)

// Retrieves the customercollection and passes them to a custom print function
func ShowCustomerCollection(db *sql.DB) {
	customerCollection := getCustomerCollection(db)
	printCustomerCollection(customerCollection)
}

func AddCustomer(db *sql.DB) {

}

func GetCustomerById(db *sql.DB, customerId int) people.Customer {
	query := "SELECT * FROM Customer WHERE customerId = ?;"
	result, err := db.Query(query, customerId)
	if err != nil {
		panic(err)
	}

	var (
		name              string
		surname           string
		postalCode        string
		houseNumber       int
		houseNumberSuffix sql.NullString
		customer          people.Customer
	)

	if result.Next() {
		err := result.Scan(&customerId, &name, &surname, &postalCode, &houseNumber, &houseNumberSuffix)
		if err != nil {
			panic(err)
		}

		customer.SetCustomerId(customerId)
		customer.SetName(name)
		customer.SetSurname(surname)
		customer.SetPostalCode(postalCode)
		customer.SetHouseNumber(houseNumber)
		customer.SetHouseNumberSuffix(houseNumberSuffix)
	}
	result.Close()
	return customer
}

//Retrieves a slice of Customer objects from the database.
func getCustomerCollection(db *sql.DB) []people.Customer {
	query := "SELECT * FROM Customer;"
	result, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	var (
		customerId        int
		name              string
		surname           string
		postalCode        string
		houseNumber       int
		houseNumberSuffix sql.NullString
	)

	customerCollection := make([]people.Customer, 0)

	for result.Next() {
		var customer people.Customer
		err := result.Scan(&customerId, &name, &surname, &postalCode, &houseNumber, &houseNumberSuffix)
		if err != nil {
			panic(err)
		}

		customer.SetCustomerId(customerId)
		customer.SetName(name)
		customer.SetSurname(surname)
		customer.SetPostalCode(postalCode)
		customer.SetHouseNumber(houseNumber)
		customer.SetHouseNumberSuffix(houseNumberSuffix)

		customerCollection = append(customerCollection, customer)
	}
	result.Close()
	return customerCollection
}

// uses the golang tabwriter to display all customers in a table.
func printCustomerCollection(customers []people.Customer) {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 10, 10, 2, ' ', tabwriter.Debug) //Debug flag for lines

	fmt.Fprintln(writer, "klantnummer\tnaam\tpostcode\thuisnummer")

	for _, customer := range customers {
		name := customer.Name() + " " + customer.Surname()
		houseNumber := strconv.Itoa(customer.HouseNumber())

		if customer.HouseNumberSuffix().Valid {
			houseNumber += customer.HouseNumberSuffix().String
		}

		output := fmt.Sprintf("%d\t%s\t%s\t%s", customer.CustomerId(), name, customer.PostalCode(), houseNumber)
		fmt.Fprintln(writer, output)
	}

	writer.Flush()
}