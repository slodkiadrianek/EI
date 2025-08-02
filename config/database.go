package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


type Db struct {
	DbConnection *sql.DB
}

func NewDb(databaseLink string) *Db {
	fmt.Println("DB Link:", databaseLink)

	dbConnection, err := sql.Open("postgres", databaseLink)
	if err != nil {
		panic(err)
	}
	 if err := dbConnection.Ping(); err != nil {
        panic(err)
    }

	return &Db{
		DbConnection: dbConnection,
	}
}