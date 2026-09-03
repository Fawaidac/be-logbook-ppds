package user

import "time"

type User struct {
	ID        int       `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	NimNip       string    `db:"nim_nip" json:"nim_nip,omitempty"`
	Role      string    `db:"role" json:"role"`
	Jabatan   string    `db:"jabatan" json:"jabatan"`
	ProgramStudi string `db:"program_studi" json:"program_studi,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type UserRegistration struct {
	ID           int       `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Nik          string    `db:"nik" json:"nik"`
	Str          string    `db:"str" json:"str"`
	Sip          string    `db:"sip" json:"sip"`
	Username     string    `db:"username" json:"username"`
	Email        string    `db:"email" json:"email"`
	Password     string    `db:"password" json:"-"`
	Specialty    string    `db:"specialty" json:"specialty"`
	ProgramStudi string    `db:"program_studi" json:"program_studi"`
	University   string    `db:"university" json:"university"`
	SelfiePath   string    `db:"selfie_path" json:"selfie_path"`
	StrFilePath  string    `db:"str_file_path" json:"str_file_path"`
	SipFilePath  string    `db:"sip_file_path" json:"sip_file_path"`
	Status          string    `db:"status" json:"status"`
	RejectionReason string    `db:"rejection_reason" json:"rejection_reason,omitempty"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}