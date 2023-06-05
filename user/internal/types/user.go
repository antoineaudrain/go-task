package types

import (
	"github.com/jackc/pgx/v4/pgxpool"
	pb "go-task/user/api"
	"go-task/user/internal/service"
	"go-task/user/internal/store/postgres"
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
