package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danielemery/rabbitmq-publish-cron/util"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func main() {
	config, configErr := util.LoadConfig()

	if configErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config")
		log.Panic(configErr)
	}

	log.Printf("config loaded successfully, attempting to connect to %s \n", config.RabbitUrl)

	conn, err := amqp.Dial(config.RabbitUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	ch.ExchangeDeclare(
		config.ExchangeName, // name
		amqp.ExchangeFanout, // kind
		false,               // durable
		false,               // autoDelete
		false,               // internal
		false,               // nowait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare an exchange")
	log.Printf("Created exchange with name %s", config.ExchangeName)

	err = ch.Publish(
		config.ExchangeName, // exchange
		"",                  // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(config.MessageBody),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", config.MessageBody)
}
