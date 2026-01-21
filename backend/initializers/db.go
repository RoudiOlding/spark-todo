package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable to access the database from anywhere
var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatal(":( —— Failed to connect to database:", err)
	}

	log.Println("✨✨ v3 CONNECTED WITH SIMPLE PROTOCOL ✨✨")
}
