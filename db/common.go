package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luchonicolosi/gambituser/env"
)

var Db *sql.DB

func DbConnect() error {

	Db, err := sql.Open("mysql", ConnStr())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//Buena practica
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexion exitosa a la db.")
	return nil
}

func ConnStr() string {
	//La mejor practica es guardarlo en un SecretManager
	dbUser := env.GetEnv("DB_USER", "default-user")
	dbName := env.GetEnv("DB_NAME", "default-dbName")
	dbPassword := env.GetEnv("DB_PASSWORD", "default-password")
	host := env.GetEnv("DB_HOST", "default-host")

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, dbPassword, host, dbName)
	fmt.Println(url)

	return url
}
