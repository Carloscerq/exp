package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
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

	port, portPresent := os.LookupEnv("PORT")
	if !portPresent {
		InfoLogger.Println("Using default port 8080")
		port = "8000"
	}

	dbStr, dbStrPresent := os.LookupEnv("DB_STR")
	if !dbStrPresent {
		ErrorLogger.Fatal("Missing DB_STR")
		panic("Missing DB_STR")
	}

	Db, err := sql.Open("mysql", dbStr)
	if err != nil {
		panic(err)
	}

	ping := Db.Ping()
	if ping != nil {
		ErrorLogger.Fatal(ping)
	}

	InfoLogger.Println("Connected to db instance")
	InfoLogger.Println("Starting server on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
