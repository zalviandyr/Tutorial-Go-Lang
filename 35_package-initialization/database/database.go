package database

var connection string

func init() {
	connection = "PostgresSQL"
}

func GetConnection() string {
	return connection
}
