package items

import (
	"database/sql"
	"fmt"
	"os"
	"text/tabwriter"
	"van-der-binckes/items/structs"
)

func GetAccessoryById(db *sql.DB, accessoryId int) structs.Accessory {
	query := "SELECT name, price FROM Accessory WHERE accessoryId = ?;"
	result, err := db.Query(query, accessoryId)
	if err != nil {
		panic(err)
	}

	var (
		name      string
		price     float64
		accessory structs.Accessory
	)

	if result.Next() {
		err := result.Scan(&name, &price)
		if err != nil {
			panic(err)
		}

		accessory.SetAccessoryId(accessoryId)
		accessory.SetName(name)
		accessory.SetPrice(price)
	} else {
		panic("no accessory found.")
	}
	result.Close()
	return accessory
}

// uses the golang tabwriter to display all accessories
func PrintAccessoryCollection(db *sql.DB) {
	accessoryCollection := getAllAccessories(db)
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 10, 10, 2, ' ', tabwriter.Debug) //Debug flag for lines

	fmt.Fprintln(writer, "accessoirenummer\tnaam\tprijs")

	for _, accessory := range accessoryCollection {
		output := fmt.Sprintf(
			"%d\t%s\t%.2f",
			accessory.AccessoryId(),
			accessory.Name(),
			accessory.Price())

		fmt.Fprintln(writer, output)
	}

	writer.Flush()
}

// Retrieves all accessories in a slice
func getAllAccessories(db *sql.DB) []structs.Accessory {
	query := "SELECT * FROM Accessory;"
	result, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	var (
		accessory   structs.Accessory
		accessoryId int
		name        string
		price       float64
	)

	accessoryCollection := make([]structs.Accessory, 0)
	for result.Next() {
		err := result.Scan(&accessoryId, &name, &price)
		if err != nil {
			panic(err)
		}

		accessory.SetAccessoryId(accessoryId)
		accessory.SetName(name)
		accessory.SetPrice(price)

		accessoryCollection = append(accessoryCollection, accessory)
	}
	result.Close()
	return accessoryCollection
}
