package config

import (
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseMigration(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
	)
}

var DB *gorm.DB

func DatabaseConnection() {
	dsn := "host=localhost user=postgres password=postgres dbname=notes_db port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DatabaseMigration(db)
	DB = db
}
