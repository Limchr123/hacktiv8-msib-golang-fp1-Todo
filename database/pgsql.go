package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"todo/entity"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	var (
		host     = os.Getenv("PGHOST")
		port     = os.Getenv("PGPORT")
		dbname   = os.Getenv("PGDATABASE")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Error occurred while trying to connect to database:", err)
	}

	if err := db.AutoMigrate(&entity.Todo{}); err != nil {
		log.Panic("Error occurred while trying to perform database migrations:", err)
	}
}

func GetDataBaseInstance() *gorm.DB {
	return db
}
