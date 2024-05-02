package handler

import (
	"encoding/json"
	"fmt"
	"github.com/mrjonze/goexpert-clean-architecture/pkg/events"
	"github.com/rabbitmq/amqp091-go"
	"sync"
)

type OrdersListedHandler struct {
	RabbitMQChannel *amqp091.Channel
}

func NewOrdersListedHandler(rabbitMQChannel *amqp091.Channel) *OrdersListedHandler {
	return &OrdersListedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrdersListedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Orders listed: %v", event.GetPayload())
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
