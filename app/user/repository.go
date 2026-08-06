package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindAll(ctx context.Context) ([]User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id int) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, name, email, password, role, jabatan) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`
	return r.db.QueryRowContext(ctx, query, user.Username, user.Name, user.Email, user.Password, user.Role, user.Jabatan).Scan(&user.ID, &user.CreatedAt)
}

func (r *repository) FindAll(ctx context.Context) ([]User, error) {
	var users []User
	query := `SELECT id, username, name, email, password, role, jabatan, created_at FROM users ORDER BY id ASC`
	err := r.db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) FindByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	query := `SELECT id, username, name, email, password, role, jabatan, created_at FROM users WHERE username = $1`
	err := r.db.GetContext(ctx, &user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := `SELECT id, username, name, email, password, role, jabatan, created_at FROM users WHERE email = $1`
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByID(ctx context.Context, id int) (*User, error) {
	var user User
	query := `SELECT id, username, name, email, password, role, jabatan, created_at FROM users WHERE id = $1`
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Update(ctx context.Context, user *User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3, role = $4, jabatan = $5 WHERE id = $6`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Role, user.Jabatan, user.ID)
	return err
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}