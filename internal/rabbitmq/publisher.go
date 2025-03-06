package rabbitmq

import (
	"encoding/json"
	"finance/internal/types"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func publishOrder(order *types.Order) error {
	conn, err := Connect()
	if err != nil {
		log.Fatalf("error %v", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("error %v", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"trade_orders", // Queue name
		true,           // Durable (ensures queue is not lost on restart)
		false,          // Auto-delete (false = keep the queue even if no consumers)
		false,          // Exclusive (false = multiple connections can use it)
		false,          // No-wait (waits for a response)
		nil,            // Additional arguments
	)
	if err != nil {
		log.Fatalf("error %v", err)
		return err
	}

	// 4️⃣ Convert the order struct to JSON format
	body, err := json.Marshal(order)
	if err != nil {
		log.Printf("Failed to serialize order: %v", err)
		return err
	}
	err = ch.Publish("", // Exchange (empty = default exchange)
		q.Name, // Routing key (queue name)
		false,  // Mandatory (false = discard if no queue)
		false,  // Immediate (false = store if no consumers available)
		amqp091.Publishing{
			ContentType: "application/json", // Message format
			Body:        body,               // Message content
		},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("📩 Order published: %+v\n", order)
	return nil
}
