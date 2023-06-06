package service

import (
	"github.com/google/uuid"
	"go-task/user/internal/event"
	"go-task/user/internal/store"
	"go-task/user/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

type UserService struct {
	userStore *store.UserStore
}

type Service interface {
	Create(email, password, fullName string) (*models.User, error)
	Authenticate(email, password string) (*models.User, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

var _ Service = (*UserService)(nil)

func NewUserService() *UserService {
	userStore, err := store.NewUserStore(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to create user store: %v", err)
	}

	return &UserService{
		userStore: userStore,
	}
}

func (s *UserService) Create(email, password, fullName string) (*models.User, error) {
	hashedPassword, err := s.HashPassword(password)
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

func (s *UserService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.userStore.GetUserByEmail(email)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid email or password")
	}

	valid := s.CheckPasswordHash(user.PasswordHash, password)
	if valid != true {
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

func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *UserService) CheckPasswordHash(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
