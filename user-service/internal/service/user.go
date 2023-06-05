package service

import (
	"github.com/google/uuid"
	"go-task/user-service/internal/event"
	"go-task/user-service/internal/store/postgres"
	"go-task/user-service/pkg/models"
	"go-task/user-service/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
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
	store, err := postgres.NewUserStore("postgresql://go_task:go_task@localhost:5432/user_db")
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
		ID:       uuid.New(),
		Email:    email,
		Password: hashedPassword,
		FullName: fullName,
	}

	if err := s.userStore.CreateUser(user); err != nil {
		return nil, err
	}

	event.PublishUserCreated(user.ID.String(), user.Email, user.Password)

	user.Password = "" // Don't return hashed password
	return user, nil
}

func (s *UserService) GetUserByEmail(email, password string) (*models.User, error) {
	user, err := s.userStore.GetUserByEmail(email)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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
