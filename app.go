package main

import (
	"log"
	"net/http"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func main() {
	port, portPresent := os.LookupEnv("PORT")
	if !portPresent {
		port = "8000"
	}

	InfoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger.Println("Starting server on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
