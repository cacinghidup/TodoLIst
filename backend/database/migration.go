package database

import (
	"Moonlay/models"
	"Moonlay/pkg/mysql"
	"fmt"
)

func MigrationDB() {
	err := mysql.DB.AutoMigrate(
		&models.List{},
		&models.SubList{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Success")

}
