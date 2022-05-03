package handlers

import (
	"encoding/json"
	"fmt"
	//"gorm/pkg/kafkaClients"
	//"gorm/pkg/models"
	"github.com/MarkCarterE/gorm/src/pkg/kafkaClients"
	"github.com/MarkCarterE/gorm/src/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	//"github.com/tutorials/go/crud/pkg/models"
	//NOTE - I need to create a new package - asyncMsgHandlers (for Kafka consumer and producer funcs)
	//You should have created your package with go mod init e.g. go mod init github.com/my-org/my-package
	//https://stackoverflow.com/questions/35480623/how-to-import-local-packages-in-go
)

// Defining AddBook as a Handler Function on the Handler Struct - this is the way we reuse the initialized DB and as such reuse the existing connection
func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the Books table - NOTE: with this func declared on the Handler Struct we now have access to the DB object declared and initialized on the Handler Struct
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
		// Send a 500 error response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result.Error)
		return
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

	// Write updated response to Kafka
	fmt.Println("Writing AddBook to Kafka " + string(body))
	kafkaClients.Produce(string(body), "localhost:9092", "go_demo_account_request")

}
