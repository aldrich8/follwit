package repository

import "follwit/infra/database"

func Save(model interface{}) error {
	err := database.DB.Save(model).Error
	if err != nil {
		println("Error: ", err.Error())
	}
	return err
}

func Find(model interface{}, id string) error {
	err := database.DB.First(model, id).Error
	if err != nil {
		println("Error: ", err.Error())
	}
	return err
}
