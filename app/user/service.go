package user

import (
	"context"
	"errors"

	"be-logbook-ppds/pkg/utils"
)

type Service interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*UserResponse, error)
	GetAllUsers(ctx context.Context) ([]UserResponse, error)
	GetUserByID(ctx context.Context, id int) (*UserResponse, error)
	UpdateUser(ctx context.Context, id int, req UpdateUserRequest) (*UserResponse, error)
	DeleteUser(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(ctx context.Context, req CreateUserRequest) (*UserResponse, error) {
	// Cek apakah username sudah ada
	existingUser, _ := s.repo.FindByUsername(ctx, req.Username)
	if existingUser != nil {
		return nil, errors.New("username sudah digunakan")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("gagal memproses kata sandi")
	}

	u := &User{
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
		Jabatan:  req.Jabatan,
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		Jabatan:   u.Jabatan,
		CreatedAt: u.CreatedAt,
	}, nil
}

func (s *service) GetAllUsers(ctx context.Context) ([]UserResponse, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var res []UserResponse
	for _, u := range users {
		res = append(res, UserResponse{
			ID:        u.ID,
			Username:  u.Username,
			Name:      u.Name,
			Email:     u.Email,
			Role:      u.Role,
			Jabatan:   u.Jabatan,
			CreatedAt: u.CreatedAt,
		})
	}
	return res, nil
}

func (s *service) GetUserByID(ctx context.Context, id int) (*UserResponse, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("pengguna tidak ditemukan")
	}
	return &UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		Jabatan:   u.Jabatan,
		CreatedAt: u.CreatedAt,
	}, nil
}

func (s *service) UpdateUser(ctx context.Context, id int, req UpdateUserRequest) (*UserResponse, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("pengguna tidak ditemukan")
	}

	u.Name = req.Name
	u.Email = req.Email
	u.Role = req.Role
	u.Jabatan = req.Jabatan

	if req.Password != "" {
		hashed, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, errors.New("gagal memproses kata sandi baru")
		}
		u.Password = hashed
	}

	if err := s.repo.Update(ctx, u); err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		Jabatan:   u.Jabatan,
		CreatedAt: u.CreatedAt,
	}, nil
}

func (s *service) DeleteUser(ctx context.Context, id int) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("pengguna tidak ditemukan")
	}

	return s.repo.Delete(ctx, id)
}


