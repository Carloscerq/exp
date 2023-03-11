package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
  "github.com/Carloscerq/exp/migrations"

	_ "github.com/go-sql-driver/mysql"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	Db            *sql.DB
)

func main() {
	InfoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	flag_value_pointer := flag.String("action", "", "")
	flag.Parse()
	flag_value := string(*flag_value_pointer)
	if flag_value != "up" && flag_value != "down" {
		ErrorLogger.Fatal("Invalid operation")
	}

	port, portPresent := os.LookupEnv("PORT")
	if !portPresent {
		InfoLogger.Println("Using default port 8080")
		port = "8000"
	}

	dbStr, dbStrPresent := os.LookupEnv("DB_STR")
	if !dbStrPresent {
		ErrorLogger.Fatal("Missing DB_STR")
	}

	Db, err := sql.Open("mysql", dbStr)
	if err != nil {
    ErrorLogger.Fatal(err)
	}

	ping := Db.Ping()
	if ping != nil {
		ErrorLogger.Fatal(ping)
	}

	migrations.Migrate(flag_value, Db)
	InfoLogger.Println("Connected to db instance")
	InfoLogger.Println("Starting server on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
