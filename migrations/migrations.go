package migrations

import (
	"follwit/infra/database"
	"follwit/models"
)

func Migrate() {
	migrations := []interface{}{
		&models.Boilerplate{},
	}

	if err := database.GetDB().AutoMigrate(migrations...); err != nil {
		println("Error migrating database")

		panic(err)

		return
	}
}
