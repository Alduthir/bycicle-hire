package order

import (
	"database/sql"
	"fmt"
	"van-der-binckes/items"
	"van-der-binckes/items/structs"
	orderStruct "van-der-binckes/order/structs"
)

// Gets a slice of all OrderAccessories for a specific orderLine
func getOrderAccessoryCollection(db *sql.DB, orderLine orderStruct.OrderLine) []orderStruct.OrderAccessory {
	query := "SELECT accessoryId, amount FROM OrderAccessory WHERE orderLineId = ?;"
	result, err := db.Query(query, orderLine.OrderLineId())
	if err != nil {
		panic(err)
	}

	var (
		accessoryId int
		amount      int
	)

	orderAccessoryCollection := make([]orderStruct.OrderAccessory, 0)

	for result.Next() {
		var orderAccessory orderStruct.OrderAccessory
		err := result.Scan(&accessoryId, &amount)
		if err != nil {
			panic(err)
		}

		orderAccessory.SetOrder(orderLine)
		orderAccessory.SetAmount(amount)
		orderAccessory.SetAccessory(items.GetAccessoryById(db, accessoryId))

		orderAccessoryCollection = append(orderAccessoryCollection, orderAccessory)
	}
	result.Close()
	return orderAccessoryCollection
}

// Recursively adds accessories to the order.
func addAccessoryToOrder(
	db *sql.DB,
	orderAccessoryCollection []orderStruct.OrderAccessory) []orderStruct.OrderAccessory {
	fmt.Println("Wilt u accessoires toevoegen aan de bestelling? (ja/nee)")

	var response string
	fmt.Scanf("%s", &response)

	if response == "nee" {
		return orderAccessoryCollection
	}

	accessory := requestAccessory(db)
	fmt.Println(fmt.Sprintf("Hoeveel accessoires van het type %s wil de klant huren?", accessory.Name()))

	var amount int
	fmt.Scanf("%d", &amount)

	orderAccessory := orderStruct.NewOrderAccessory(accessory, amount)
	orderAccessoryCollection = append(orderAccessoryCollection, orderAccessory)

	return addAccessoryToOrder(db, orderAccessoryCollection)
}

// Request and retrieve a single accessory by id
func requestAccessory(db *sql.DB) (accessory structs.Accessory) {
	items.PrintAccessoryCollection(db)
	fmt.Println("Welke accessoire wil de klant huren? (accessoirenummer)")

	var accessoryId int
	fmt.Scanf("%d", &accessoryId)

	defer func(db *sql.DB, accessoryId int) {
		r := recover()
		if r != nil {
			fmt.Println(fmt.Printf("Geen accessoire met nummer %d gevonden.", accessoryId))
			accessory = requestAccessory(db)
		}
	}(db, accessoryId)

	accessory = items.GetAccessoryById(db, accessoryId)
	return accessory
}

// Inserts a single orderAccessory linked to the given orderLineId
func insertOrderAccessory(db *sql.DB, orderAccessory orderStruct.OrderAccessory, orderLineId int) {
	query := "INSERT INTO OrderAccessory (orderLineId, accessoryId, amount) values (?, ?, ?)"

	accessory := orderAccessory.Accessory()
	_, err := db.Query(
		query,
		orderLineId,
		accessory.AccessoryId(),
		orderAccessory.Amount())

	if err != nil {
		panic(err)
	}
}
