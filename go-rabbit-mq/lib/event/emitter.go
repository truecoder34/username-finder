package event

/*
	create PUBLISHER here
*/

import (
	"log"

	"github.com/streadway/amqp"
)

/*
	Publisher == Emmiter  declaration. to publish AMWP events
	Emmitter structure contains amqp connection
*/
type Emitter struct {
	connection *amqp.Connection
}

/*
	func :
		retrive a channel from connection pool
		call declare Exchange for channel
*/
func (e *Emitter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()
	return declareExchange(channel)
}

/*
	Func to publish exact message to the AMQP exchange
	- get channel from ppol
	[ param 1 ] - event - msg to be sent
	[ param 2 ] - severity - logging severity - to define which msg sent by which subscriber
*/
func (e *Emitter) Push(event string, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	err = channel.Publish(
		getExchangeName(),
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	log.Printf("Sending message: %s -> %s", event, getExchangeName())
	return nil
}

/*
	func to return new event.Emitter object
	detect is connection is esteblished
*/
func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}

	// init object and check that it was initialized without error
	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}
