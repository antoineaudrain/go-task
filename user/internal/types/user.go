package types

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go-task/user/internal/service"
	"go-task/user/internal/store/postgres"
	pb "go-task/user/pkg/pb"
)

//go:generate mockery --name Service --inpackage-suffix --output ../mocks --case underscore --with-expecter
type (
	Server struct {
	}

	Handler struct {
		pb.UnimplementedUserServiceServer
		userService *service.UserService
	}

	Service struct {
		userStore *postgres.UserStore
	}

	Store struct {
		db *pgxpool.Pool
	}
)
