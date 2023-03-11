package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
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

	path, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}
	path = path + "/migrations"

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Fatal(err)
		}

		if info.IsDir() {
			return nil
		}

		if strings.Contains(path, flag_value) {
			logger.Println(path)
		}

		return nil
	})
}
