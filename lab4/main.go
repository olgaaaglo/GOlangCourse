package main

import (
	"os"
)

func main() {
	var database Database
	var peopleStruct PeopleStruct
	var offersStruct OffersStruct

	database.initialize(&peopleStruct, &offersStruct)
	if len(os.Args) == 3 {
		database.fill(os.Args[1], os.Args[2])
	} else {
		database.fill("People.json", "Offers.json")
	}
	database.print()

	var offersStruct2 OffersStruct
	database.addToDatabase(&offersStruct2)
	database.print()
	database.saveToFiles()
}