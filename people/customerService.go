package people

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"
	people "van-der-binckes/people/structs"
)

// Retrieves the customercollection and passes them to a custom print function
func ShowCustomerCollection(db *sql.DB) {
	customerCollection := getCustomerCollection(db)
	printCustomerCollection(customerCollection)
}

func AddCustomerToOrder(db *sql.DB) (customer people.Customer) {
	fmt.Println("Voor welke klant wilt u een bestelling toevoegen? (klantnummer)")

	var customerId int
	fmt.Scanf("%d", &customerId)

	defer func(db *sql.DB, customerId int) {
		r := recover()
		if r != nil {
			fmt.Println(fmt.Printf("Geen persoon met klantnummer %d gevonden.", customerId))
			customer = AddCustomerToOrder(db)
		}
	}(db, customerId)

	customer = GetCustomerById(db, customerId)
	return customer
}

// Creates a new customer and inserts it into the database.
func AddCustomer(db *sql.DB) {
	// Use the bufio reader to read multiple words at once.
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hoe heet deze nieuwe klant? Voor & Achternaam invullen.")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	nameCollection := strings.Split(text, " ")
	if len(nameCollection) != 2 {
		fmt.Println("Uw invoer werdt niet herkend. Voer de naam en achternaam in, gescheiden door één spatie.")
		return
	}

	fmt.Println("Wat is de postcode van de nieuwe klant?")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	regex := regexp.MustCompile(`^\W$`) // Remove whitespace
	postalCode := regex.ReplaceAllString(text, "")

	// Matches the input on 4 digits followed by 2 letters.
	matches, _ := regexp.MatchString(`^\d{4}[a-zA-Z]{2}$`, postalCode)

	if matches == false {
		fmt.Println("Een postcode moet bestaan uit 4 cijfers en 2 letters")
		return
	}

	fmt.Println("En tot slot, wat is het huisnummer van de nieuwe klant?")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	regex = regexp.MustCompile(`[^0-9.]`)
	houseNumber, _ := strconv.Atoi(regex.ReplaceAllString(text, "")) //Remove all non numeric characters
	regex = regexp.MustCompile(`[^a-zA-Z]`)                          // Remove all numeric characters
	houseNumberSuffix := sql.NullString{String: regex.ReplaceAllString(text, "")}
	houseNumberSuffix.Valid = false

	if len(houseNumberSuffix.String) > 0 {
		houseNumberSuffix.Valid = true
	}

	query := "INSERT INTO Customer (name, surname, postalcode, housenumber, houseNumberSuffix) values (?,?,UPPER(?),?,?)"

	_, err := db.Query(
		query,
		strings.Title(nameCollection[0]),
		strings.Title(nameCollection[1]),
		postalCode,
		houseNumber,
		houseNumberSuffix)

	if err != nil {
		panic(err)
	}

	query = "SELECT customerId FROM Customer WHERE surname = ? AND postalCode = UPPER(?) AND houseNumber = ?"
	result, err := db.Query(query, strings.Title(nameCollection[1]), postalCode, houseNumber)
	if err != nil {
		panic(err)
	}

	var customerId int

	if result.Next() {
		err := result.Scan(&customerId)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Gebruiker sucessvol toegevoegd met klantnummer " + strconv.Itoa(customerId))
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
