package store

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  []byte    `json:"-"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) (int, error) {
	query := `INSERT INTO users (username,email,password) 
				values($1,$2,$3) RETURNING id,created_at , updated_at`
	result := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	)

	err := result.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return -1, err
	}
	return user.ID, nil
}
func (s *UsersStore) GetByID(ctx context.Context, id int) (User, error) {
	query := `Select username , email,created_at , updated_at ,password , id  from  users where id = $1`
	var user User
	result := s.db.QueryRowContext(
		ctx,
		query,
		id,
	)

	err := result.Scan(&user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Password, &user.ID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UsersStore) GetByEmail(ctx context.Context, email string) (User, error) {
	query := `Select username , email,created_at , updated_at ,password , id  from  users where email= $1`
	var user User
	result := s.db.QueryRowContext(
		ctx,
		query,
		email,
	)

	err := result.Scan(&user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Password, &user.ID)
	if err != nil {
		return user, err
	}
	return user, nil
}
