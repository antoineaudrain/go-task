package auth

import pb "go-task/auth-service/pkg/pb"

//go:generate mockery --name Service --inpackage-suffix --output ../mocks --case underscore --with-expecter
type (
	Server struct {
		pb.UnimplementedAuthServiceServer
		Service Service
	}

	Service interface {
		Create(email, password string) (string, error)
		Login(email, password string) (string, error)
	}
)
