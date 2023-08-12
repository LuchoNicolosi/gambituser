package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func DbConnect() error {

	Db, err = sql.Open("mysql", ConnStr())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexion exitosa a la db.")
	return nil
}

func ConnStr() string {

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, dbPassword, host, dbName)
	fmt.Println(url)

	return url
}
