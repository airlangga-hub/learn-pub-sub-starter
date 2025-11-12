package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")

	const connString = "amqp://guest:guest@localhost:5672/"

	clientConn, err := amqp.Dial(connString)
	if err != nil {
		log.Fatal("Error creating RabbitMQ connection: ", err)
	}

	defer clientConn.Close()

	fmt.Println("Success connecting client to RabbitMQ!")

	username, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Fatal("Error creating username: ", err)
	}

	_, queue, err := pubsub.DeclareAndBind(
		clientConn,
		routing.ExchangePerilDirect,
		routing.PauseKey+"."+username,
		routing.PauseKey,
		pubsub.SimpleQueueTransient,
	)
	if err != nil {
		log.Fatal("could not subscribe to pause: ", err)
	}

	fmt.Printf("Queue %v declared and bound!\n", queue.Name)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("RabbitMQ connection closed.")
}
