package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID              int          `db:"id" json:"id"`
	Username        string       `db:"username" json:"username"`
	Name            string       `db:"name" json:"name"`
	Email           string       `db:"email" json:"email"`
	Password        string       `db:"password" json:"-"`
	NimNip          string       `db:"nim_nip" json:"nim_nip,omitempty"`
	Role            string       `db:"role" json:"role"`
	Jabatan         string       `db:"jabatan" json:"jabatan"`
	ProgramStudi    string       `db:"program_studi" json:"program_studi,omitempty"`
	Nik             string       `db:"nik" json:"nik,omitempty"`
	BirthPlace      string       `db:"birth_place" json:"birth_place,omitempty"`
	BirthDate       sql.NullTime `db:"birth_date" json:"birth_date,omitempty"`
	Gender          string       `db:"gender" json:"gender,omitempty"`
	BloodType       string       `db:"blood_type" json:"blood_type,omitempty"`
	Religion        string       `db:"religion" json:"religion,omitempty"`
	University      string       `db:"university" json:"university,omitempty"`
	Citizenship     string       `db:"citizenship" json:"citizenship,omitempty"`
	EmailPersonal   string       `db:"email_personal" json:"email_personal,omitempty"`
	PhoneMobile     string       `db:"phone_mobile" json:"phone_mobile,omitempty"`
	PhoneHome       string       `db:"phone_home" json:"phone_home,omitempty"`
	AddressKtp      string       `db:"address_ktp" json:"address_ktp,omitempty"`
	AddressDomicile string       `db:"address_domicile" json:"address_domicile,omitempty"`
	StrNumber       string       `db:"str_number" json:"str_number,omitempty"`
	StrIssued       sql.NullTime `db:"str_issued" json:"str_issued,omitempty"`
	StrExpired      sql.NullTime `db:"str_expired" json:"str_expired,omitempty"`
	StrNote         string       `db:"str_note" json:"str_note,omitempty"`
	Profession      string       `db:"profession" json:"profession,omitempty"`
	Competency      string       `db:"competency" json:"competency,omitempty"`
	College         string       `db:"college" json:"college,omitempty"`
	Stage           string       `db:"stage" json:"stage,omitempty"`
	Dpjp            string       `db:"dpjp" json:"dpjp,omitempty"`
	CreatedAt       time.Time    `db:"created_at" json:"created_at"`
}

type UserRegistration struct {
	ID              int       `db:"id" json:"id"`
	Name            string    `db:"name" json:"name"`
	Nik             string    `db:"nik" json:"nik"`
	Str             string    `db:"str" json:"str"`
	Sip             string    `db:"sip" json:"sip"`
	Username        string    `db:"username" json:"username"`
	Email           string    `db:"email" json:"email"`
	Password        string    `db:"password" json:"-"`
	Specialty       string    `db:"specialty" json:"specialty"`
	ProgramStudi    string    `db:"program_studi" json:"program_studi"`
	University      string    `db:"university" json:"university"`
	SelfiePath      string    `db:"selfie_path" json:"selfie_path"`
	StrFilePath     string    `db:"str_file_path" json:"str_file_path"`
	SipFilePath     string    `db:"sip_file_path" json:"sip_file_path"`
	Status          string    `db:"status" json:"status"`
	RejectionReason string    `db:"rejection_reason" json:"rejection_reason,omitempty"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

type WorkHistory struct {
	ID           int          `db:"id" json:"id"`
	UserUsername string       `db:"user_username" json:"user_username"`
	Position     string       `db:"position" json:"position"`
	Institution  string       `db:"institution" json:"institution"`
	StartDate    sql.NullTime `db:"start_date" json:"start_date"`
	Status       string       `db:"status" json:"status"`
	Sip          string       `db:"sip" json:"sip"`
	Location     string       `db:"location" json:"location"`
	CreatedAt    time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at" json:"updated_at"`
}
