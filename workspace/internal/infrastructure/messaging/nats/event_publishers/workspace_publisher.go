package event_publishers

import (
	"github.com/antoineaudrain/go-task/workspace/internal/domain/workspace"
	"github.com/google/uuid"
	"log"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"

	pb "github.com/antoineaudrain/go-task/workspace/api"
)

type WorkspacePublisher struct {
	NatsConn *nats.Conn
}

func NewWorkspacePublisher(nc *nats.Conn) *WorkspacePublisher {
	return &WorkspacePublisher{NatsConn: nc}
}

func (wp *WorkspacePublisher) PublishWorkspaceCreated(workspace *workspace.Workspace) error {
	event := &pb.WorkspaceCreatedEvent{
		Id:            uuid.New().String(),
		WorkspaceId:   workspace.ID.String(),
		WorkspaceName: workspace.Name,
	}

	data, err := proto.Marshal(event)
	if err != nil {
		return err
	}

	if err := wp.NatsConn.Publish("WorkspaceCreated", data); err != nil {
		log.Printf("Error publishing WorkspaceCreated event: %v", err)
		return err
	}

	return nil
}
