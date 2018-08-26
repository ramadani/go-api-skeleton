package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"  // mysql dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite dialect

	"github.com/spf13/viper"
)

// Database contains db.
type Database struct {
	DB *gorm.DB
}

// Init the database connection and returns db.
func Init(cog *viper.Viper) *Database {
	driver := cog.GetString("db.driver")

	var connection string

	switch driver {
	case "mysql":
		host := cog.GetString("mysql.host")
		port := cog.GetString("mysql.port")
		user := cog.GetString("mysql.user")
		pass := cog.GetString("mysql.password")
		dbName := cog.GetString("mysql.dbname")
		connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", user, pass, host, port, dbName)
	default:
		connection = cog.GetString("sqlite3.db")
	}

	db, err := gorm.Open(driver, connection)

	if err != nil {
		panic("failed to connect database")
	}

	return &Database{db}
}
