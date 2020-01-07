package items

import (
	"database/sql"
	"fmt"
	"os"
	"text/tabwriter"
	items "van-der-binckes/items/structs"
)

func GetBicycleById(db *sql.DB, bicycleId int) items.Bicycle {
	query := "SELECT name, type, price, amount FROM Bicycle WHERE bicycleId = ?;"
	result, err := db.Query(query, bicycleId)
	if err != nil {
		panic(err)
	}

	var (
		name        string
		bicycleType string
		price       float64
		amount      int
		bicycle     items.Bicycle
	)

	if result.Next() {
		err := result.Scan(&name, &bicycleType, &price, &amount)
		if err != nil {
			panic(err)
		}

		bicycle.SetBicycleId(bicycleId)
		bicycle.SetName(name)
		bicycle.SetBicycleType(bicycleType)
		bicycle.SetPrice(price)
		bicycle.SetAmount(amount - getCurrentlyHiredStock(bicycleId, db))
	}
	result.Close()
	return bicycle
}

func AddBicycleToOrder(db *sql.DB) (bicycle items.Bicycle) {
	printBicycleCollection(db)
	fmt.Println("Welke fiets wil de klant huren? (fietsnummer)")

	var bicycleId int
	fmt.Scanf("%d", &bicycleId)

	defer func(db *sql.DB, bicycleId int) {
		r := recover()
		if r != nil {
			fmt.Println(fmt.Printf("Geen fiets met nummer %d gevonden.", bicycleId))
			bicycle = AddBicycleToOrder(db)
		}
	}(db, bicycleId)

	bicycle = GetBicycleById(db, bicycleId)
	return bicycle
}

// uses the golang tabwriter to display all bicycles and their current stock.
func printBicycleCollection(db *sql.DB) {
	bicycleCollection := getAllBicycles(db)
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 10, 10, 2, ' ', tabwriter.Debug) //Debug flag for lines

	fmt.Fprintln(writer, "fietsnummer\tnaam\ttype\tprijs\tvoorraad")

	for _, bicycle := range bicycleCollection {
		output := fmt.Sprintf(
			"%d\t%s\t%s\t%.2f\t%d",
			bicycle.BicycleId(),
			bicycle.Name(),
			bicycle.BicycleType(),
			bicycle.Price(),
			bicycle.Amount())

		fmt.Fprintln(writer, output)
	}

	writer.Flush()
}

// Retrieves all bicycles in a slice
func getAllBicycles(db *sql.DB) []items.Bicycle {
	query := "SELECT * FROM Bicycle;"
	result, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	var (
		bicycleId   int
		name        string
		bicycleType string
		price       float64
		amount      int
		bicycle     items.Bicycle
	)

	bicycleCollection := make([]items.Bicycle, 0)
	for result.Next() {
		err := result.Scan(&bicycleId, &name, &bicycleType, &price, &amount)
		if err != nil {
			panic(err)
		}

		bicycle.SetBicycleId(bicycleId)
		bicycle.SetAmount(bicycleId)
		bicycle.SetName(name)
		bicycle.SetBicycleType(bicycleType)
		bicycle.SetPrice(price)
		bicycle.SetAmount(amount - getCurrentlyHiredStock(bicycleId, db))

		bicycleCollection = append(bicycleCollection, bicycle)
	}
	result.Close()
	return bicycleCollection
}

// Checks which bicycles are not currently used in an Order which has not yet ended.
func getCurrentlyHiredStock(bicycleId int, db *sql.DB) int {
	query := "SELECT COUNT(ol.bicycleId) stock FROM OrderLine ol WHERE DATE_ADD(startDate, INTERVAL days DAY) > CURDATE()  AND bicycleId = ? GROUP BY ol.bicycleId;"
	result, err := db.Query(query, bicycleId)
	if err != nil {
		panic(err)
	}

	stock := 0
	if result.Next() {
		err := result.Scan(&stock)
		if err != nil {
			panic(err)
		}
	}
	result.Close()
	return stock
}
