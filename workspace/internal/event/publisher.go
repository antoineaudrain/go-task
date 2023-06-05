package event

import (
	"log"

	"github.com/nats-io/nats.go"
	pb "go-task/auth-service/pkg/pb"
	"google.golang.org/protobuf/proto"
)

func PublishUserRegistered(userId, email string) {
	event := &pb.UserRegisteredEvent{
		UserId: userId,
		Email:  email,
	}

	// Serialize the event as a protobuf message
	eventData, err := proto.Marshal(event)
	if err != nil {
		log.Printf("Failed to serialize event: %v", err)
		return
	}

	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Printf("Failed to connect to NATS server: %v", err)
		return
	}
	defer nc.Close()

	// Publish the serialized event to NATS
	err = nc.Publish("user-registered-events", eventData)
	if err != nil {
		log.Printf("Failed to publish event: %v", err)
	}
}
