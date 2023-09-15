package utils

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func GetDb(env string) *sql.DB {
	var connStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		"postgres",
		"postgres",
		"service-core-db",
		"service-core")

	if env != "test" {
		if err := godotenv.Load(".env"); err != nil {
			log.Panic(err)
		}

		connStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"))
	}

	//db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Panic(err)
	}

	return db
}
