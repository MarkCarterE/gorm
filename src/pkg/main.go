package main

import (
	"log"
	"net/http"

	// "gorm/pkg/handlers"
	// "gorm/pkg/db"
	"github.com/MarkCarterE/gorm/src/pkg/db" 
	"github.com/MarkCarterE/gorm/src/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	//init the DB
	DB := db.Init()
	//main.Init()
	//init the Handlers and inject the initialized GORM DB
	h := handlers.New(DB)

	router := mux.NewRouter()

	//with the handler (h) initialized we can now utilize the h.Addbook, GetAllBooks etc.
	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	//http.ListenAndServe(":4000", router)
	http.ListenAndServe(":8080", router)

	//kafkaClients.Produce("ledgerString", "localhost:9092", "AccountWithdrawlRequest")
}
