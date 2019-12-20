package order

import (
	"database/sql"
	"van-der-binckes/items"
	orderStruct "van-der-binckes/order/structs"
)

func getOrderAccessoryCollection(db *sql.DB, orderLine orderStruct.OrderLine) []orderStruct.OrderAccessory {
	query := "SELECT accessoryId, amount FROM OrderAccessory WHERE orderLineId = ?;"
	result, err := db.Query(query, orderLine.OrderLineId())
	if err != nil {
		panic(err)
	}

	var (
		accessoryId int
		amount  int
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


