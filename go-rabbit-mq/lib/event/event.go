package event

import (
	"github.com/streadway/amqp"
)

/*
	return exchange name
	can be any value
*/
func getExchangeName() string {
	return "logs_topic"
}

/*
	Mthod to create namless (random name) queue;
	exclusive queue - only one subscriber is possible
*/
func declareRandonQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

/*
	method to declare exchange
	method is idempotent and will not create duplicate of exchange

	[ param 1 ] - pointer to AMQP connection channel
*/
func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		getExchangeName(), // name
		"topic",           // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
}
