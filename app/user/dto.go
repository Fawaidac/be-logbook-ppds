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

type CreateRegistrationRequest struct {
	Name                 string `json:"name" form:"name" binding:"required"`
	Nik                  string `json:"nik" form:"nik"`
	Str                  string `json:"str" form:"str"`
	Sip                  string `json:"sip" form:"sip"`
	Username             string `json:"username" form:"username" binding:"required"`
	Email                string `json:"email" form:"email" binding:"required,email"`
	Password             string `json:"password" form:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" binding:"required"`
	Specialty            string `json:"specialty" form:"specialty"`
	ProgramStudi         string `json:"program_studi" form:"program_studi"`
	University           string `json:"university" form:"university"`
}

type UserRegistrationResponse struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Nik             string    `json:"nik"`
	Str             string    `json:"str"`
	Sip             string    `json:"sip"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Specialty    string    `json:"specialty"`
	ProgramStudi string    `json:"program_studi"`
	University   string    `json:"university"`
	SelfiePath   string    `json:"selfie_path"`
	StrFilePath  string    `json:"str_file_path"`
	SipFilePath  string    `json:"sip_file_path"`
	Status          string    `json:"status"`
	RejectionReason string    `json:"rejection_reason,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}

type RejectRegistrationRequest struct {
	Reason string `json:"reason" form:"reason" binding:"required"`
}
