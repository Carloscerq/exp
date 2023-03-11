package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "MIGRATIONS: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Running migrations...")

	flag_value_pointer := flag.String("action", "", "")
	flag.Parse()
	flag_value := string(*flag_value_pointer)
	if flag_value != "up" && flag_value != "down" {
		logger.Fatal("Invalid operation")
	}


}
