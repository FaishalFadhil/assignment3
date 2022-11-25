package main

import (
	"assignment3/controllers"
	"assignment3/database"
)

func main() {
	database.StartDB()
	controllers.UpdateData()
}
