package msgbrocker

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type senderInterface interface {
	SendMessage([]string, string)
}

type senderEntity struct{}

var Sender senderInterface = &senderEntity{}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (snd *senderEntity) SendMessage(urls []string, endpoint string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"username-endpoint", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	var i int = 0
	for _, url := range urls {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("[DATA] " + url + "; [API endpoint] " + endpoint),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x - %d] Sent data: [DATA] %s;  [API endpoint] %s\n", i, url, endpoint)
		i++
	}
}
