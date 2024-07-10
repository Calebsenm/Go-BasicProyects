package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "caleb"
	DB_PASSWORD = "12345"
	DB_NAME     = "test"
)

func ConnectPostgres() (*sql.DB, error) {
	// Crear cadena de conexión
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, DB_USER, DB_PASSWORD, DB_NAME)

	// Abrir conexión a la base de datos
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al abrir conexión a la base de datos:", err)
	}

	// Comprobar que la conexión funciona correctamente
	err = db.Ping()
	if err != nil {
		log.Fatal("Error al hacer ping a la base de datos:", err)
	}

	// Devolver la conexión
	return db, err
}



