package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const DATABASE_DRIVER_NAME = "postgres"

func Init(dbHost, dbUser, dbPassword, dbName string, dbPort int64) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	connection, err := sqlx.Open(DATABASE_DRIVER_NAME, dsn)
	if err != nil {
		return connection, err
	}

	err = connection.Ping()
	if err != nil {
		return connection, err
	}

	log.Println("Database connection established")

	return connection, nil
}