package sagas

//import (
//	"github.com/golang/protobuf/proto"
//	"log"
//
//	"github.com/antoineaudrain/go-task/workspace/internal/domain/models"
//	"github.com/antoineaudrain/go-task/workspace/internal/domain/services"
//	"github.com/nats-io/nats.go"
//)
//
//type WorkspaceSaga struct {
//	workspaceService *services.WorkspaceService
//	NatsConn         *nats.Conn
//}
//
//func NewWorkspaceSaga(ws *services.WorkspaceService, nc *nats.Conn) *WorkspaceSaga {
//	return &WorkspaceSaga{
//		workspaceService: ws,
//		NatsConn:         nc,
//	}
//}
//
//func (ws *WorkspaceSaga) Start() {
//	if _, err := ws.NatsConn.Subscribe("BoardCreated", ws.handleBoardCreated); err != nil {
//		log.Printf("Error subscribing to BoardCreated event: %v", err)
//		return
//	}
//}
//
//func (ws *WorkspaceSaga) handleBoardCreated(m *nats.Msg) {
//	event := &pb.BoardCreatedEvent{}
//	if err := proto.Unmarshal(m.Data, event); err != nil {
//		log.Printf("Error unmarshaling BoardCreated event: %v", err)
//		return
//	}
//
//	workspaceId := event.GetWorkspaceId()
//	workspace, err := ws.CommandHandler.GetWorkspace(workspaceId)
//	if err != nil {
//		log.Printf("Error retrieving workspace: %v", err)
//		return
//	}
//
//	if workspace != nil {
//		workspace.Status = models.WorkspaceStatusActive
//
//		if err := ws.CommandHandler.UpdateWorkspace(workspace); err != nil {
//			log.Printf("Error updating workspace: %v", err)
//		}
//	}
//}
