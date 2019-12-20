package items

import (
	"database/sql"
	items "van-der-binckes/items/structs"
)

func GetBycicleById(db *sql.DB, bycicleId int) items.Bycicle {
	query := "SELECT name, type, price, amount FROM Bycicle WHERE bycicleId = ?;"
	result, err := db.Query(query, bycicleId)
	if err != nil {
		panic(err)
	}

	var (
		name        string
		bycicleType string
		price       float64
		amount      int
		bycicle     items.Bycicle
	)

	if result.Next() {
		err := result.Scan(&name, &bycicleType, &price, &amount)
		if err != nil {
			panic(err)
		}

		bycicle.SetName(name)
		bycicle.SetBycicleType(bycicleType)
		bycicle.SetPrice(price)
		bycicle.SetAmount(amount)
	}
	result.Close()
	return bycicle
}
