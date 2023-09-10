package main

import (
	"follwit/config"
	"follwit/infra/database"
	"follwit/migrations"
	"follwit/routers"
	"net/http"
)

func main() {
	dbDSN := config.DBConfiguration("dev")

	if _, err := database.DBConnection(dbDSN); err != nil {
		println("Error connecting to database")
		panic(err)
	}

	migrations.Migrate()

	// Initialize router
	router := routers.SetupRoute()

	err := http.ListenAndServe(":5200", router)

	if err != nil {
		panic(err)
	}
}
