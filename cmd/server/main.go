package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Starting Peril server...")

	const connString = "amqp://guest:guest@localhost:5672/"

	amqpConn, err := amqp.Dial(connString)
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ: ", err)
	}

	defer amqpConn.Close()

	fmt.Println("Success connecting to RabbitMQ!")

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt)

	<-signalChan

	fmt.Println("RabbitMQ connection closed.")
}