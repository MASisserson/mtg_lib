package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "aikanu"
	password = "1234"
	dbname   = "mtg_lib"
)

var DB *sql.DB

func initDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func closeDB() {
	DB.Close()
}

func Create(query string, objects []any) error {
	initDB()
	defer closeDB()

	_, err := DB.Exec(query, objects...)
	if err != nil {
		return err
	}

	return nil
}

func Read(query string) (*sql.Rows, error) {
	initDB()
	defer closeDB()

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
