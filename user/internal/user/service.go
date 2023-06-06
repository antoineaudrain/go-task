package user

import (
	"github.com/google/uuid"
	"go-task/core/pkg/models"
	"go-task/user/internal/event"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

type (
	service struct {
		store Store
	}

	Service interface {
		Create(email, password, fullName string) (*models.User, error)
		Authenticate(email, password string) (*models.User, error)
		GetUserByID(userID string) (*models.User, error)
		HashPassword(password string) (string, error)
		CheckPasswordHash(password, hash string) bool
	}
)

var _ Service = (*service)(nil)

func NewService() Service {
	s, err := NewStore(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to create user store: %v", err)
	}

	return &service{
		store: s,
	}
}

func (s *service) Create(email, password, fullName string) (*models.User, error) {
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

	if err := s.store.CreateUser(user); err != nil {
		return nil, err
	}

	event.PublishUserCreated(user.ID.String(), user.Email, user.PasswordHash)

	user.PasswordHash = "" // Don't return hashed password
	return user, nil
}

func (s *service) Authenticate(email, password string) (*models.User, error) {
	user, err := s.store.GetUserByEmail(email)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid email or password")
	}

	valid := s.CheckPasswordHash(user.PasswordHash, password)
	if valid != true {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid email or password")
	}

	return user, nil
}

func (s *service) GetUserByID(userID string) (*models.User, error) {
	user, err := s.store.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *service) CheckPasswordHash(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
