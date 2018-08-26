package db

import (
	"fmt"

	"github.com/ramadani/go-api-skeleton/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"  // mysql dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite dialect
)

// Database contains db library.
type Database struct {
	DB *gorm.DB
}

// Init the database connection and returns db.
func Init(cog *config.Config) *Database {
	driver := cog.Config.GetString("db.driver")

	var connection string

	switch driver {
	case "mysql":
		host := cog.Config.GetString("mysql.host")
		port := cog.Config.GetString("mysql.port")
		user := cog.Config.GetString("mysql.user")
		pass := cog.Config.GetString("mysql.password")
		dbName := cog.Config.GetString("mysql.dbname")
		connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", user, pass, host, port, dbName)
	default:
		connection = cog.Config.GetString("sqlite3.db")
	}

	db, err := gorm.Open(driver, connection)

	if err != nil {
		panic("failed to connect database")
	}

	return &Database{db}
}
