package user

import (
	"github.com/google/uuid"
	customErrors "go-task/core/pkg/errors"
	"go-task/core/pkg/models"
	"go-task/user/internal/event"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

type (
	service struct {
		store Store
	}

	Service interface {
		CreateUserWithHashedPassword(email, password, fullName string) (*models.User, error)
		AuthenticateUser(email, password string) (*models.User, error)
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

func (s *service) CreateUserWithHashedPassword(email, password, fullName string) (*models.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, customErrors.NewHashingError("failed to hash password", err)
	}

	user := &models.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: hashedPassword,
		FullName:     fullName,
	}

	if err := s.store.CreateUser(user); err != nil {
		return nil, customErrors.NewDatabaseError("failed to create user", err)
	}

	event.PublishUserCreated(user.ID.String(), user.Email, user.FullName)

	user.PasswordHash = ""
	return user, nil
}

func (s *service) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.store.GetUserByEmail(email)
	if err != nil {
		return nil, customErrors.NewAuthenticationError("Invalid email or password", err)
	}

	valid := checkPasswordHash(password, user.PasswordHash)
	if valid != true {
		return nil, customErrors.NewAuthenticationError("Invalid email or password", err)
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
