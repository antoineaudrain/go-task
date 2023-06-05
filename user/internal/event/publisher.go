package event

import (
	"log"

	"github.com/nats-io/nats.go"
	pb "go-task/user/api"
	"google.golang.org/protobuf/proto"
)

func PublishUserCreated(userId, email, fullName string) {
	event := &pb.UserCreatedEvent{
		UserId:   userId,
		Email:    email,
		FullName: fullName,
	}

	eventData, err := proto.Marshal(event)
	if err != nil {
		log.Printf("Failed to serialize event: %v", err)
		return
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Printf("Failed to connect to NATS server: %v", err)
		return
	}
	defer nc.Close()

	err = nc.Publish("user-created-events", eventData)
	if err != nil {
		log.Printf("Failed to publish event: %v", err)
	}
}
