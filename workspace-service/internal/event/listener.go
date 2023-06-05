package event

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func ListenToUserRegisteredEvents() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Printf("Failed to connect to NATS server: %v", err)
		return
	}
	defer nc.Close()

	// Subscribe to the NATS subject for user registered events
	sub, err := nc.SubscribeSync("user-registered-events")
	if err != nil {
		log.Printf("Failed to subscribe to user-registered-events: %v", err)
		return
	}

	// Infinite loop to receive and process messages
	for {
		msg, err := sub.NextMsg(5 * time.Second)
		if err != nil {
			//log.Printf("Failed to receive message: %v", err)
			continue
		}

		// Process the received message
		log.Printf("Received message: %s", string(msg.Data))
	}
}
