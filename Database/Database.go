package Database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Connection(iType string) *gorm.DB {

	if (iType == "first") {

	} else {

		var err error
		DB, err = gorm.Open("mysql", "DBUsername:DBPassword@/DBName?charset=utf8&parseTime=True&loc=Local") // this does not really open a new connection
		if err != nil {
			log.Fatalf("Error on initializing database connection: %s", err.Error())
		}

		DB.DB().SetMaxIdleConns(10)
		DB.DB().SetMaxOpenConns(100)

		err = DB.DB().Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
		if err != nil {
			log.Fatalf("Error on opening database connection: %s", err.Error())
		}
	}
	return DB
}
