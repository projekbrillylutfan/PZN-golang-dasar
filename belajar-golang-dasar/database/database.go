package database

var connection string

func init() {
	connection = "sqlite"
}

func GetDatabase() string {
	return connection
}