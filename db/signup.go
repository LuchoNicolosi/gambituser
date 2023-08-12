package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luchonicolosi/gambituser/models"
	"github.com/luchonicolosi/gambituser/tools"
)

func SignUp(data models.SignUp) error {
	fmt.Println("Comienza registro")
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	INSERT_USER := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + data.UserEmail + "','" + data.UserUUID + "','" + tools.FechaMySQL() + "')"

	fmt.Println(INSERT_USER)
	
	_, err = Db.Exec(INSERT_USER)
	if err != nil {
		fmt.Println("Error en ejecutar la query: ", err.Error())
		return err
	}

	fmt.Println("Signup > Ejecucion exitosa!")
	return nil
}
