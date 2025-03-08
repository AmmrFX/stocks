package rabbitmq

import (
	"encoding/json"
	"finance/internal/types"
	"log"
)

// ConsumeOrders listens for trade orders
func ConsumeOrders() {
	conn, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// Declare queue
	q, err := ch.QueueDeclare("trade_orders", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Waiting for orders...")

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			var order types.Order
			err := json.Unmarshal(msg.Body, &order)
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}

			log.Printf("Processing order: %+v\n", order)

			// Process order in business logic
		//	types.ExecuteTrade(order)
		}
	}()

	<-forever
}
