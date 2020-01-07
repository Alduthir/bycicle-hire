package order

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
	"van-der-binckes/items"
	itemStruct "van-der-binckes/items/structs"
	orderStruct "van-der-binckes/order/structs"
	"van-der-binckes/people"
	peopleStruct "van-der-binckes/people/structs"
)

// Retrieves all orders and displays them in the console.
func ShowOrderCollection(db *sql.DB) {
	orderCollection := getOrderCollection(db)
	printOrderCollection(orderCollection)
}

// Adds a new order, including accessories.
func AddOrder(employee peopleStruct.Employee, db *sql.DB) {
	customer := people.AddCustomerToOrder(db)
	bicycle := items.AddBicycleToOrder(db)
	orderAccessoryCollection := make([]orderStruct.OrderAccessory, 0)
	orderAccessoryCollection = addAccessoryToOrder(db, orderAccessoryCollection)

	fmt.Println("Voor hoeveel dagen wil de klant deze bestelling afnemen?")

	var days int
	fmt.Scanf("%d", &days)

	totalPrice := calculateTotalPrice(bicycle, orderAccessoryCollection, days)
	fmt.Println(fmt.Sprintf("De totaalprijs is €%.2f. Is dit akkoord? (ja/nee)", totalPrice))

	var response string
	fmt.Scanf("%s", &response)

	if response == "ja" {
		insertOrder(db, bicycle, customer, employee, days, totalPrice, orderAccessoryCollection)

		fmt.Println("Order succesvol toegevoegd.")
		return
	}

	fmt.Println("Order geannuleerd.")
}

//  Inserts a new orderline and instantly retrieves it's id to then insert the orderAccessories
func insertOrder(
	db *sql.DB,
	bicycle itemStruct.Bicycle,
	customer peopleStruct.Customer,
	employee peopleStruct.Employee,
	days int,
	totalPrice float64,
	orderAccessoryCollection []orderStruct.OrderAccessory) {

	query := "INSERT INTO OrderLine (bicycleId, customerId, employeeId, startDate, days, totalPrice) values (?, ?, ?, NOW(), ?, ?);"
	_, err := db.Query(
		query,
		bicycle.BicycleId(),
		customer.CustomerId(),
		employee.EmployeeId(),
		days,
		totalPrice)
	if err != nil {
		panic(err)
	}
	query = "SELECT orderLineId FROM OrderLine WHERE bicycleId = ? AND customerId = ? AND employeeId = ? ORDER BY orderLineId DESC LIMIT 1"
	result, err := db.Query(
		query,
		bicycle.BicycleId(),
		customer.CustomerId(),
		employee.EmployeeId())
	if err != nil {
		panic(err)
	}
	if result.Next() {
		var orderLineId int
		err := result.Scan(&orderLineId)
		if err != nil {
			panic(err)
		}

		for _, orderAccessory := range orderAccessoryCollection {
			insertOrderAccessory(db, orderAccessory, orderLineId)
		}
	}
	result.Close()
}

// Retrieves a slice containing all orders.
func getOrderCollection(db *sql.DB) []orderStruct.OrderLine {
	query := "SELECT * FROM OrderLine;"
	result, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	var (
		orderLineId int
		bicycle     int
		customerId  int
		employeeId  int
		startDate   time.Time
		days        int
		totalPrice  float64
	)

	orderCollection := make([]orderStruct.OrderLine, 0)

	for result.Next() {
		var order orderStruct.OrderLine
		err := result.Scan(&orderLineId, &bicycle, &customerId, &employeeId, &startDate, &days, &totalPrice)
		if err != nil {
			panic(err)
		}

		order.SetOrderLineId(orderLineId)
		order.SetStartDate(startDate)
		order.SetDays(days)
		order.SetTotalPrice(totalPrice)
		order.SetBicycle(items.GetBicycleById(db, bicycle))
		order.SetCustomer(people.GetCustomerById(db, customerId))
		order.SetEmployee(people.GetEmployeeById(db, employeeId))
		order.SetOrderAccessoryCollection(getOrderAccessoryCollection(db, order))

		orderCollection = append(orderCollection, order)
	}
	result.Close()
	return orderCollection
}

// Prints all orders using the tabwriter
func printOrderCollection(orderLines []orderStruct.OrderLine) {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 10, 10, 2, ' ', tabwriter.Debug) //Debug flag for lines

	fmt.Fprintln(writer, "ordernummer\tklant\tmedewerker\tfiets\taccessoires\tstartdatum\teindDatum\tprijs")

	for _, orderLine := range orderLines {
		// Explicitly retrieve all objects before using their getters because go is stupid and can't read pointers.
		customer := orderLine.Customer()
		customerName := customer.Name() + " " + customer.Surname()

		employee := orderLine.Customer()
		employeeName := employee.Name() + " " + employee.Surname()

		startDate := orderLine.StartDate()
		days := orderLine.Days()
		endDate := startDate.AddDate(0, 0, days)

		accessories := ""
		for index, orderAccessory := range orderLine.OrderAccessoryCollection() {
			accessory := orderAccessory.Accessory()
			accessories += strconv.Itoa(orderAccessory.Amount()) + "x " + accessory.Name()

			if index != len(orderLine.OrderAccessoryCollection()) {
				accessories += "; "
			}
		}

		bicycle := orderLine.Bicycle()

		output := fmt.Sprintf(
			"%d\t%s\t%s\t%s\t%s\t%s\t%s\t%.2f",
			orderLine.OrderLineId(),
			customerName,
			employeeName,
			bicycle.BicycleType(),
			accessories,
			startDate.Format("2006-01-02"),
			endDate.Format("2006-01-02"),
			orderLine.TotalPrice())

		fmt.Fprintln(writer, output)
	}

	writer.Flush()
}

// Calculates the totalprice of the bicycle and accessories multiplied by the amount of days
func calculateTotalPrice(
	bycicle itemStruct.Bicycle,
	orderAccessoryCollection []orderStruct.OrderAccessory,
	days int) float64 {
	totalPrice := 0.00
	totalPrice += bycicle.Price()

	for _, orderAccessory := range orderAccessoryCollection {
		accessory := orderAccessory.Accessory()
		totalPrice += accessory.Price()
	}

	return totalPrice * float64(days)
}
