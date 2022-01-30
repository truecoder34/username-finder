package event

import (
	"log"

	"github.com/streadway/amqp"
)

/*
	sctruct of consumer. reciver of AMQP events
	1 - connection to AMQP server
	2 - random queue name
*/
type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

/*
	like in Emitter - be sure that ecxchage is declared
*/
func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	return declareExchange(channel)
}

/*
	create and return new consumer
	[ param 1 ] - new Consumer
	[ out 1 ] - consumer
	[ out 2 ] - error
*/
func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}
	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

/*
	Listener - listen all new Queue publications and print  to console
	[ param 1 ] - topics string arr - binding keys
*/
func (consumer *Consumer) Listen(topics []string) error {
	// take channel from cntn pool
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// declare namless queue
	q, err := declareRandonQueue(ch)
	if err != nil {
		return err
	}

	for _, s := range topics {
		// for each bnding key - bind queue to echange  . specify what we went to recieve
		err = ch.QueueBind(
			q.Name,
			s, // key
			getExchangeName(),
			false,
			nil,
		)

		if err != nil {
			return err
		}
	}

	// start listening on the queue and rcieve messages
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	// itterate through messages andprint
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for message [Exchange, Queue][%s, %s]. To exit press CTRL+C", getExchangeName(), q.Name)
	<-forever
	return nil
}
