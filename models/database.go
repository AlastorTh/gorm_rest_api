package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host=db_postgres port=5432 user=postgres dbname=adlist password=postgres sslmode=disable")

	if err != nil {
		panic("Can't establish connection to DB")
	}
	db.AutoMigrate(&Ad{})
	return db
}
