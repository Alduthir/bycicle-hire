package order

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
	"van-der-binckes/items"
	orderStruct "van-der-binckes/order/structs"
	"van-der-binckes/people"
	peoplseStruct "van-der-binckes/people/structs"
)

func ShowOrderCollection(db *sql.DB) {
	orderCollection := getOrderCollection(db)
	printOrderCollection(orderCollection)
}

func AddOrder(employee peoplseStruct.Employee, db *sql.DB) {

}

func getOrderCollection(db *sql.DB) []orderStruct.OrderLine {
	query := "SELECT * FROM OrderLine;"
	result, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	var (
		orderLineId int
		bycicleId   int
		customerId  int
		employeeId  int
		startDate   time.Time
		days        int
		totalPrice  float64
	)

	orderCollection := make([]orderStruct.OrderLine, 0)

	for result.Next() {
		var order orderStruct.OrderLine
		err := result.Scan(&orderLineId, &bycicleId, &customerId, &employeeId, &startDate, &days, &totalPrice)
		if err != nil {
			panic(err)
		}

		order.SetOrderLineId(orderLineId)
		order.SetStartDate(startDate)
		order.SetDays(days)
		order.SetTotalPrice(totalPrice)
		order.SetBycicle(items.GetBycicleById(db, bycicleId))
		order.SetCustomer(people.GetCustomerById(db, customerId))
		order.SetEmployee(people.GetEmployeeById(db, employeeId))
		order.SetOrderAccessoryCollection(getOrderAccessoryCollection(db, order))

		orderCollection = append(orderCollection, order)
	}
	result.Close()
	return orderCollection
}

func printOrderCollection(orderLines []orderStruct.OrderLine) {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 10, 10, 2, ' ', tabwriter.Debug) //Debug flag for lines

	fmt.Fprintln(writer, "ordernummer\tklant\tmedewerker\tfiets\taccessoires\tstartdatum\teindDatum\tprijs")

	for _, orderLine := range orderLines {
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

		bycicle := orderLine.Bycicle()

		output := fmt.Sprintf(
			"%d\t%s\t%s\t%s\t%s\t%s\t%.2f",
			orderLine.OrderLineId(),
			customerName,
			employeeName,
			bycicle.BycicleType(),
			accessories,
			endDate.String(),
			orderLine.TotalPrice())

		fmt.Fprintln(writer, output)
	}

	writer.Flush()
}
