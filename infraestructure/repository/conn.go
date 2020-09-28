package repository

import (
	"os"

	_ "github.com/lib/pq"
)

func init() {
	SetDbFireBird(os.Getenv("APPSETTINGS__CONNECTIONSTRING"))
	_conn = newConnection()
}

// SetDbFireBird set db postgres and connection string
func SetDbFireBird(connectionString string) {
	setDb("postgres", connectionString)
}
