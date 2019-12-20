package items

import (
	"database/sql"
	"van-der-binckes/items/structs"
)

func GetAccessoryById(db *sql.DB, accessoryId int) structs.Accessory {
	query := "SELECT name, price FROM Accessory WHERE accessoryId = ?;"
	result, err := db.Query(query, accessoryId)
	if err != nil {
		panic(err)
	}

	var (
		name  string
		price float64
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
	}
	result.Close()
	return accessory
}
