package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
)

func main() {
	fmt.Println("Starting Peril server...")

	const connString = "amqp://guest:guest@localhost:5672/"

	amqpConn, err := amqp.Dial(connString)
	if err != nil {
		log.Fatal("Error connecting server to RabbitMQ: ", err)
	}

	defer amqpConn.Close()

	fmt.Println("Success connecting server to RabbitMQ!")

	publishChannel, err := amqpConn.Channel()
	if err != nil {
		log.Fatal("Error creating RabbitMQ Channel: ", err)
	}

	err = pubsub.PublishJSON(
		publishChannel,
		routing.ExchangePerilDirect,
		routing.PauseKey,
		routing.PlayingState{
			IsPaused: true,
		},
	)
	if err != nil {
		log.Printf("could not publish time: %v", err)
	}

	fmt.Println("Pause message sent!")
}