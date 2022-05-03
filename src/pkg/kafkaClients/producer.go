package kafkaClients

import (
	"log"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
)

// var ctx context.Context

// func init() {
// 	//init the Kafka context
// 	ctx = context.Background()
// }

//The function in another package must start with a capital letter so that it is public outside its package

func Produce(ledgerString string, brokerAddress string, topic string) {
	// initialize a counter
	i := 0
	//The kafka-go library comes with the option of providing a logger that can provide more detailed information about the state of your kafka brokers.
	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// can be set to -1, 0, or 1
		// 1 is a good default for most non-transactional data - see: https://www.sohamkamani.com/golang/working-with-kafka/
		// for financial data however, we may want to use -1 to ensure all brokers acknowlege they have received the msg
		RequiredAcks: 1,
		// assign the logger to the writer
		Logger: l,
	})

	//for {
	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(i)),
		// create an arbitrary message payload for the value
		Value: []byte(ledgerString),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

	// log a confirmation once the message is written
	//fmt.Println(ledgerString)
	//sagaLedgerEntry(ledgerString)
	//i++
	// sleep for a second
	//time.Sleep(time.Second)
	//}
}
