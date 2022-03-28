package database

import (
	"log"
	"user-crpyto-portfolio/entity"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

func MigrateUser(table *entity.User) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}

func MigratePortfolio(table *entity.Portfolio) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
