package handler

import (
	"encoding/json"
	"fmt"
	"github.com/mrjonze/goexpert-clean-architecture/pkg/events"
	"github.com/rabbitmq/amqp091-go"
	"sync"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp091.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp091.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Order created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitMq := amqp091.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish("amq.direct",
		"",
		false,
		false,
		msgRabbitMq,
	)
}
