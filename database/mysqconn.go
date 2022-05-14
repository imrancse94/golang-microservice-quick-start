package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
	"go.quick.start/config"
)

// ConnectDB to sql database
func ConnectDB() *gorm.DB {
	connectionString, driver := createConnectionString()
	db, err := gorm.Open(driver, connectionString)

	if err != nil {
		log.Error(err)
	}

	log.Info("Database connection successful")
	return db
}

// Create string for SQL connection
func createConnectionString() (string, string) {
	conf := config.GetSQL()
	return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	), conf.Driver
}
