package user

import "time"

type UserResponse struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	NimNip       string    `json:"nim_nip,omitempty"`
	Jabatan      string    `json:"jabatan"`
	ProgramStudi string    `json:"program_studi,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	Username     string `json:"username" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
	Role         string `json:"role" binding:"required"`
	NimNip       string `json:"nim_nip"`
	Jabatan      string `json:"jabatan"`
	ProgramStudi string `json:"program_studi"`
}

type UpdateUserRequest struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password"`
	Role         string `json:"role" binding:"required"`
	NimNip       string `json:"nim_nip"`
	Jabatan      string `json:"jabatan"`
	ProgramStudi string `json:"program_studi"`
}
