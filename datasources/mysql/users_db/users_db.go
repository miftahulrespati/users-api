package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Client *sql.DB
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		username = os.Getenv("MYSQL_USERS_USERNAME")
		password = os.Getenv("MYSQL_USERS_PASSWORD")
		host     = os.Getenv("MYSQL_USERS_HOST")
		schema   = os.Getenv("MYSQL_USERS_SCHEMA")
	)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username,
		password,
		host,
		schema,
	)

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
