package provider

type Database struct{}

func (db Database) Boot() {
	// driver := viper.GetString("db.driver")
}

func InitDatabase() *Database {
	return &Database{}
}
