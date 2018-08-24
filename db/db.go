package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/spf13/viper"
)

func Init(cog *viper.Viper) *gorm.DB {
	driver := cog.GetString("db.driver")

	var connection string

	switch driver {
	case "mysql":
		host := cog.GetString("mysql.host")
		port := cog.GetString("mysql.port")
		user := cog.GetString("mysql.user")
		pass := cog.GetString("mysql.password")
		dbname := cog.GetString("mysql.dbname")
		connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", user, pass, host, port, dbname)
	default:
		connection = cog.GetString("sqlite3.db")
	}

	db, err := gorm.Open(driver, connection)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	return db
}