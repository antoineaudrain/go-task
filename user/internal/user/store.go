package user

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-task/core/pkg/models"
)

type (
	store struct {
		db *pgxpool.Pool
	}

	Store interface {
		CreateUser(user *models.User) error
		GetUserByEmail(email string) (*models.User, error)
	}
)

var _ Store = (*store)(nil)

func NewStore(connStr string) (Store, error) {
	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &store{db: db}, nil
}

func (s *store) CreateUser(user *models.User) error {
	sqlStatement := `
		INSERT INTO users (id, email, password_hash, full_name)
		VALUES ($1, $2, $3, $4)
	`

	_, err := s.db.Exec(context.Background(), sqlStatement, user.ID, user.Email, user.PasswordHash, user.FullName)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) GetUserByEmail(email string) (*models.User, error) {
	sqlStatement := `
		SELECT id, email, password_hash, full_name
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	row := s.db.QueryRow(context.Background(), sqlStatement, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.FullName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
