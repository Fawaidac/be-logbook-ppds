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
	CreateRegistration(ctx context.Context, reg *UserRegistration) error
	FindAllRegistrations(ctx context.Context, status string) ([]UserRegistration, error)
	FindRegistrationByID(ctx context.Context, id int) (*UserRegistration, error)
	UpdateRegistrationStatus(ctx context.Context, id int, status string, rejectionReason string) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, name, email, password, role, jabatan, program_studi, nim_nip) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at`
	return r.db.QueryRowContext(ctx, query, user.Username, user.Name, user.Email, user.Password, user.Role, user.Jabatan, user.ProgramStudi, user.NimNip).Scan(&user.ID, &user.CreatedAt)
}

func (r *repository) FindAll(ctx context.Context) ([]User, error) {
	var users []User
	query := `SELECT id, username, name, email, password, role, jabatan, COALESCE(nim_nip, '') AS nim_nip, COALESCE(program_studi, '') AS program_studi, created_at FROM users ORDER BY id ASC`
	err := r.db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) FindByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	query := `SELECT id, username, name, email, password, role, jabatan, COALESCE(nim_nip, '') AS nim_nip, COALESCE(program_studi, '') AS program_studi, created_at FROM users WHERE username = $1`
	err := r.db.GetContext(ctx, &user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := `SELECT id, username, name, email, password, role, jabatan, COALESCE(nim_nip, '') AS nim_nip, COALESCE(program_studi, '') AS program_studi, created_at FROM users WHERE email = $1`
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByID(ctx context.Context, id int) (*User, error) {
	var user User
	query := `SELECT id, username, name, email, password, role, jabatan, COALESCE(nim_nip, '') AS nim_nip, COALESCE(program_studi, '') AS program_studi, created_at FROM users WHERE id = $1`
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Update(ctx context.Context, user *User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3, role = $4, jabatan = $5, program_studi = $6, nim_nip = $7 WHERE id = $8`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Role, user.Jabatan, user.ProgramStudi, user.NimNip, user.ID)
	return err
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *repository) CreateRegistration(ctx context.Context, reg *UserRegistration) error {
	query := `INSERT INTO user_registrations (name, nik, str, sip, username, email, password, specialty, program_studi, university, selfie_path, str_file_path, sip_file_path, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query, reg.Name, reg.Nik, reg.Str, reg.Sip, reg.Username, reg.Email, reg.Password, reg.Specialty, reg.ProgramStudi, reg.University, reg.SelfiePath, reg.StrFilePath, reg.SipFilePath, reg.Status).Scan(&reg.ID, &reg.CreatedAt, &reg.UpdatedAt)
}

func (r *repository) FindAllRegistrations(ctx context.Context, status string) ([]UserRegistration, error) {
	var regs []UserRegistration
	var query string
	var err error
	if status != "" {
		query = `SELECT id, name, COALESCE(nik, '') AS nik, COALESCE(str, '') AS str, COALESCE(sip, '') AS sip, COALESCE(username, '') AS username, email, password, COALESCE(specialty, '') AS specialty, COALESCE(program_studi, '') AS program_studi, COALESCE(university, '') AS university, COALESCE(selfie_path, '') AS selfie_path, COALESCE(str_file_path, '') AS str_file_path, COALESCE(sip_file_path, '') AS sip_file_path, status, COALESCE(rejection_reason, '') AS rejection_reason, created_at, updated_at FROM user_registrations WHERE status = $1 ORDER BY id DESC`
		err = r.db.SelectContext(ctx, &regs, query, status)
	} else {
		query = `SELECT id, name, COALESCE(nik, '') AS nik, COALESCE(str, '') AS str, COALESCE(sip, '') AS sip, COALESCE(username, '') AS username, email, password, COALESCE(specialty, '') AS specialty, COALESCE(program_studi, '') AS program_studi, COALESCE(university, '') AS university, COALESCE(selfie_path, '') AS selfie_path, COALESCE(str_file_path, '') AS str_file_path, COALESCE(sip_file_path, '') AS sip_file_path, status, COALESCE(rejection_reason, '') AS rejection_reason, created_at, updated_at FROM user_registrations ORDER BY id DESC`
		err = r.db.SelectContext(ctx, &regs, query)
	}
	if err != nil {
		return nil, err
	}
	return regs, nil
}

func (r *repository) FindRegistrationByID(ctx context.Context, id int) (*UserRegistration, error) {
	var reg UserRegistration
	query := `SELECT id, name, COALESCE(nik, '') AS nik, COALESCE(str, '') AS str, COALESCE(sip, '') AS sip, COALESCE(username, '') AS username, email, password, COALESCE(specialty, '') AS specialty, COALESCE(program_studi, '') AS program_studi, COALESCE(university, '') AS university, COALESCE(selfie_path, '') AS selfie_path, COALESCE(str_file_path, '') AS str_file_path, COALESCE(sip_file_path, '') AS sip_file_path, status, COALESCE(rejection_reason, '') AS rejection_reason, created_at, updated_at FROM user_registrations WHERE id = $1`
	err := r.db.GetContext(ctx, &reg, query, id)
	if err != nil {
		return nil, err
	}
	return &reg, nil
}

func (r *repository) UpdateRegistrationStatus(ctx context.Context, id int, status string, rejectionReason string) error {
	query := `UPDATE user_registrations SET status = $1, rejection_reason = $2, updated_at = NOW() WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, status, rejectionReason, id)
	return err
}
