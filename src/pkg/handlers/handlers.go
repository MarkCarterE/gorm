package handlers

import "gorm.io/gorm"

// Handler Struct used to intialize the DB once and inject the initialzed, open connection string, into each handler

type handler struct {
	DB *gorm.DB
}

//New function signature with pointer to the GORM DB instance
func New(db *gorm.DB) handler {
	return handler{db}
}
