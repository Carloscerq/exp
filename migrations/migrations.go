package migrations

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Migrate(f string, db *sql.DB) {
	logger := log.New(os.Stderr, "MIGRATIONS: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Running migrations...")

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

		if strings.Contains(path, f+".sql") {
			logger.Println(path)
			data_binary, err := os.ReadFile(path)

			if err != nil {
				logger.Fatal(err)
			}

			data := string(data_binary)
			res, err := (*db).Exec(data)

			if err != nil {
				logger.Fatal(err)
			}

			logger.Println(data)
			logger.Println(res)
		}

		return nil
	})
}
