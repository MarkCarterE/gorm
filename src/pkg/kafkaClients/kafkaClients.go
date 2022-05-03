package kafkaClients

import (
	"context"
	//"gormTutorial/gorm/pkg/kafkaClients"
)

var ctx context.Context

func init() {
	//init the Kafka context
	ctx = context.Background()
}

// Handler Struct used to intialize the DB once and inject the initialzed, open connection string, into each handler

//type handler struct {
type kafkaClient struct {
	//DB *gorm.DB
	// create a new context for Kafka producer (not sure is this is the write location however meaning, is the context static and can it
	//therefore be created once during service startup?). For now, leaving it here
	CTX context.Context
}

// func New(db *gorm.DB) handler {
// 	return handler{db}
// }

//
//Adding & before a variable name evaluates to it's address
//Adding * before a pointer variable dereferences the pointer to the address it is holding

func New() kafkaClient {
	return kafkaClient{ctx}
}
