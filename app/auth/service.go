package auth

import (
	"context"
	"errors"

	"be-logbook-ppds/app/user"
	"be-logbook-ppds/configs"
	"be-logbook-ppds/pkg/utils"
)

type Service interface {
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
	GetProfile(ctx context.Context, userID int) (*user.UserResponse, error)
}

type service struct {
	userRepo user.Repository
	cfg      *configs.Config
}

var akunDemoMap = map[string]struct {
	Password string
	Role     string
	Name     string
	Jabatan  string
}{
	"fawaid":       {Password: "fawaid", Role: "superadmin", Name: "Achmad Fawaid, S.Tr.Kom", Jabatan: "Super Admin"},
	"superadmin01": {Password: "superadmin01", Role: "superadmin", Name: "Super Admin Utama", Jabatan: "Super Admin System"},
	"residen01":    {Password: "residen01", Role: "residen", Name: "dr. Ratna Puspita", Jabatan: "Residen Bedah Th.2"},
	"supervisor01": {Password: "supervisor01", Role: "supervisor", Name: "dr. Budi Santoso, Sp.B", Jabatan: "DPJP Bedah Umum"},
	"admin01":      {Password: "admin01", Role: "admin", Name: "Siti Amalia, S.KM", Jabatan: "Admin Komkordik"},
}

func NewService(userRepo user.Repository, cfg *configs.Config) Service {
	return &service{userRepo: userRepo, cfg: cfg}
}

func (s *service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	var userID int
	var username, name, role, jabatan, storedPassword string
	// 1. Coba cari dari database terlebih dahulu
	u, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err == nil && u != nil {
		userID = u.ID
		username = u.Username
		name = u.Name
		role = u.Role
		jabatan = u.Jabatan
		storedPassword = u.Password
	} else {
		demo, exists := akunDemoMap[req.Username]
		if !exists {
			return nil, errors.New("Username atau password salah")
		}
		userID = 0
		username = req.Username
		name = demo.Name
		role = demo.Role
		jabatan = demo.Jabatan
		storedPassword = demo.Password
	}
	// 3. Verifikasi password (bcrypt hash dengan fallback plain-text)
	errBcrypt := utils.CheckPassword(storedPassword, req.Password)
	if errBcrypt != nil && storedPassword != req.Password {
		return nil, errors.New("Username atau password salah")
	}

	token, err := GenerateToken(userID, username, role, s.cfg.JWTSecret)
	if err != nil {
		return nil, errors.New("Gagal membuat token")
	}

	return &LoginResponse{
		User:  LoginData{Username: username, Name: name, Role: role, Jabatan: jabatan},
		Token: token,
	}, nil
}

func (s *service) GetProfile(ctx context.Context, userID int) (*user.UserResponse, error) {
	u, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, errors.New("Pengguna tidak ditemukan")
	}
	return &user.UserResponse{ID: u.ID, Username: u.Username, Name: u.Name, Email: u.Email, Role: u.Role, Jabatan: u.Jabatan}, nil
}
