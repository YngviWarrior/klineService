package services

import (
	"context"
	"encoding/json"
	"klineService/entities/kline"
	rabbitmqstructs "klineService/services/rabbitMQStructs"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func BrokerConnect() (conn *amqp.Connection, ch *amqp.Channel) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ"))

	if err != nil {
		log.Println(err, "Failed to connect to RabbitMQ")
		return
	}

	ch, err = conn.Channel()

	if err != nil {
		log.Println(err, "Failed to open a channel")
		return
	}

	return
}

func (r *RabbitMQ) SendCotation(kline *kline.Kline) {
	conn, ch := BrokerConnect()

	defer conn.Close()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		"crypto_data", // name
		"direct",      // type
		false,         // durable
		true,          // auto-deleted
		false,         // internal
		true,          // no-wait
		nil,           // arguments
	)

	if err != nil {
		log.Println(err, "Failed to declare a queue")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bytes, err := json.Marshal(kline)

	if err != nil {
		log.Println(err, "Failed to encode struct")
		return
	}

	err = ch.PublishWithContext(ctx,
		"crypto_data",     // exchange
		"crypto_cotation", // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})

	if err != nil {
		log.Println(err, "Failed to publish a message")
		return
	}
}

func (r *RabbitMQ) SendAveragePrice(avgMessage *rabbitmqstructs.InputAvgMessageDto) {
	conn, ch := BrokerConnect()

	defer conn.Close()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		"crypto_data", // name
		"direct",      // type
		false,         // durable
		true,          // auto-deleted
		false,         // internal
		true,          // no-wait
		nil,           // arguments
	)

	if err != nil {
		log.Println(err, "Failed to declare a queue")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bytes, err := json.Marshal(avgMessage)

	if err != nil {
		log.Println(err, "Failed to encode struct")
		return
	}

	err = ch.PublishWithContext(ctx,
		"crypto_data", // exchange
		"crypto_avg",  // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})

	if err != nil {
		log.Println(err, "Failed to publish a message")
		return
	}
}
