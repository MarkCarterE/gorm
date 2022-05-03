package db

import (
	"log"

	//"github.com/tutorials/go/crud/pkg/models"
	//"gorm/pkg/models"
	"github.com/dna/gorm/pkg/models v1.8.0"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres Initialization utilizing GORM
func Init() *gorm.DB {
	//dbURL := "postgres://pg:pass@localhost:5432/crud"
	dbURL := "postgres://postgres:Admin@1234@localhost:5432/postgres"

	//open a connection with the Postgres Driver
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// with the PK defined, this should migrate the Model
	db.AutoMigrate(&models.Book{})

	return db
}