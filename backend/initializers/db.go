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
	// Read the connection string from .env
	dsn := os.Getenv("DB_URL")

	// Connect using GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatal(":( —— Failed to connect to database:", err)
	}

	log.Println(":DD —— Connected to Supabase (Postgres) successfully!")

	log.Println("v2 CONNECTED TO SUPABASE (NO PREPARE)")
}
