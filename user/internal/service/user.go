package service

import (
	"github.com/google/uuid"
	"go-task/user/internal/event"
	"go-task/user/internal/store/postgres"
	"go-task/user/pkg/models"
	"go-task/user/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

type UserService struct {
	userStore *postgres.UserStore
}

type Service interface {
	Create(email, password, fullName string) (*models.User, error)
	GetUserByEmail(email, password string) (*models.User, error)
}

var _ Service = (*UserService)(nil)

func NewUserService() *UserService {
	store, err := postgres.NewUserStore(os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatalf("Failed to create user store: %v", err)
	}

	return &UserService{
		userStore: store,
	}
}

func (s *UserService) Create(email, password, fullName string) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: hashedPassword,
		FullName:     fullName,
	}

	if err := s.userStore.CreateUser(user); err != nil {
		return nil, err
	}

	event.PublishUserCreated(user.ID.String(), user.Email, user.PasswordHash)

	user.PasswordHash = "" // Don't return hashed password
	return user, nil
}

func (s *UserService) GetUserByEmail(email, password string) (*models.User, error) {
	user, err := s.userStore.GetUserByEmail(email)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid email or password")
	}

	return user, nil
}

func (s *UserService) GetUserByID(userID string) (*models.User, error) {
	user, err := s.userStore.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
