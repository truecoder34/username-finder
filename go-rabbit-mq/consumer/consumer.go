package main

import (
	"os"

	"username-finder/go-rabbit-mq/lib/event"

	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	consumer, err := event.NewConsumer(connection)
	if err != nil {
		panic(err)
	}
	consumer.Listen(os.Args[1:])
}
