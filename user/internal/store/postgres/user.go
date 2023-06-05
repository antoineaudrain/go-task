package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-task/user/pkg/models"
)

type UserStore struct {
	db *pgxpool.Pool
}

func NewUserStore(connStr string) (*UserStore, error) {
	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &UserStore{db: db}, nil
}

func (s *UserStore) CreateUser(user *models.User) error {
	sqlStatement := `
		INSERT INTO users (id, email, password, full_name)
		VALUES ($1, $2, $3, $4)
	`

	_, err := s.db.Exec(context.Background(), sqlStatement, user.ID, user.Email, user.Password, user.FullName)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetUserByEmail(email string) (*models.User, error) {
	sqlStatement := `
		SELECT id, email, password, full_name
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	row := s.db.QueryRow(context.Background(), sqlStatement, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.FullName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStore) GetUserByID(userID string) (*models.User, error) {
	sqlStatement := `
		SELECT id, email, password, full_name
		FROM users
		WHERE id = $1
		LIMIT 1
	`

	row := s.db.QueryRow(context.Background(), sqlStatement, userID)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.FullName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
