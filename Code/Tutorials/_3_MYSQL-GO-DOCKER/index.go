package main

import (
	"database/sql" // Interactuar con bases de datos
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id       int
	Name     string
	Password string
}

func main() {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/Test")
	if err != nil {
		panic(err.Error())
	}

	data, err := db.Query("SELECT * FROM users")

	defer data.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Usuarios Registrados En la base de datos ")
	for data.Next() {
		var user User

		err := data.Scan(&user.Id, &user.Name, &user.Password)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", user)
	}
}
