package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := "localhost"
	port := 5432
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := "ProyectoBD"

	if user == "" {
		user = "postgres"
	}
	if password == "" {
		password = "trabajos380"
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Printf("Intentando conexión: %s", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error al abrir conexión: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %v", err)
	}

	log.Println("✅ Conexión exitosa a la base de datos.")
	return db, nil
}
