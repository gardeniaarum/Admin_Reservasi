package config

import (
	"fmt"
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB{
	
	db, err := sql.Open("mysql", "root:@tcp(localhost:127.0.0.1:3306)/dbname")
	if err != err{
		log.Fatal(err)
	}

	fmt.Println("Koneksi Database Sukses")

	return db
}