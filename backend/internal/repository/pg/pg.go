package pg

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() func() {
	dsn := os.Getenv("POSTGRES_URL")

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d

	return func() {
		dbSQL, _ := db.DB()
		err := dbSQL.Close()
		log.Println("Failed to close DB by error", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
