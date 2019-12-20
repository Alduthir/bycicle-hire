package main

import (
	"database/sql"
	"fmt"
	"os"
	"van-der-binckes/connection"
	"van-der-binckes/order"
	"van-der-binckes/people"
	peopleStructs "van-der-binckes/people/structs"
)

/*
	The main function simply fetches the (long living) DB object and uses this to require a login.
	After login the user may perform actions through the main menu of the application.
*/
func main() {
	db := connection.FetchDatabaseConnection()
	defer db.Close()

	fmt.Println("Om dit systeem te kunnen gebruiken moet u eerst inloggen.")
	user := connection.Login(db)
	requestAction(user, db)
}

// The main menu/control loop of the application. Requests an action and will continue to be called until exit.
func requestAction(user peopleStructs.Employee, db *sql.DB) {
	fmt.Println("Wat wilt u doen?")
	fmt.Println("(1) Klanten inzien")
	fmt.Println("(2) Verhuren inzien")
	fmt.Println("(3) Klant toevoegen")
	fmt.Println("(4) Verhuur toevoegen")
	fmt.Println("(5) Applicatie afsluiten")

	var response int
	fmt.Scanf("%d", &response)

	switch response {
	case 1:
		people.ShowCustomerCollection(db)
	case 2:
		order.ShowOrderCollection(db)
	case 3:
		people.AddCustomer(db)
	case 4:
		order.AddOrder(user, db)
	case 5:
		os.Exit(0)
	default:
		fmt.Println("Uw invoer werd niet herkend. " +
			"Voer alstublieft het getal in van de actie die u wilt uitvoeren.")
	}
	fmt.Println("------------------------------------------------------------------") // For easier reading
	requestAction(user, db)
}