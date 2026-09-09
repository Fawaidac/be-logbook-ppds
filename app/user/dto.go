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

type PreRegisterRequest struct {
	Username             string `json:"username" binding:"required"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
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
	Specialty       string    `json:"specialty"`
	ProgramStudi    string    `json:"program_studi"`
	University      string    `json:"university"`
	SelfiePath      string    `json:"selfie_path"`
	StrFilePath     string    `json:"str_file_path"`
	SipFilePath     string    `json:"sip_file_path"`
	Status          string    `json:"status"`
	RejectionReason string    `json:"rejection_reason,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}

type RejectRegistrationRequest struct {
	Reason string `json:"reason" form:"reason" binding:"required"`
}

// UserProfileResponse adalah data identitas lengkap residen untuk Master Data.
// email_institution dipetakan ke kolom `email` (dipakai juga untuk login).
type UserProfileResponse struct {
	Username         string `json:"username"`
	FullName         string `json:"full_name"`
	EmailInstitution string `json:"email_institution"`
	Nik              string `json:"nik"`
	NimNip           string `json:"nim_nip"`
	ProgramStudi     string `json:"program_studi"`
	BirthPlace       string `json:"birth_place"`
	BirthDate        string `json:"birth_date"`
	Gender           string `json:"gender"`
	BloodType        string `json:"blood_type"`
	Religion         string `json:"religion"`
	University       string `json:"university"`
	Citizenship      string `json:"citizenship"`
	EmailPersonal    string `json:"email_personal"`
	PhoneMobile      string `json:"phone_mobile"`
	PhoneHome        string `json:"phone_home"`
	AddressKtp       string `json:"address_ktp"`
	AddressDomicile  string `json:"address_domicile"`
	StrNumber        string `json:"str_number"`
	StrIssued        string `json:"str_issued"`
	StrExpired       string `json:"str_expired"`
	StrNote          string `json:"str_note"`
	Profession       string `json:"profession"`
	Competency       string `json:"competency"`
	College          string `json:"college"`
	Stage            string `json:"stage"`
	Dpjp             string `json:"dpjp"`
}

type UpdateProfileRequest struct {
	FullName         string `json:"full_name" form:"full_name"`
	EmailInstitution string `json:"email_institution" form:"email_institution"`
	Nik              string `json:"nik" form:"nik"`
	BirthPlace       string `json:"birth_place" form:"birth_place"`
	BirthDate        string `json:"birth_date" form:"birth_date"`
	Gender           string `json:"gender" form:"gender"`
	BloodType        string `json:"blood_type" form:"blood_type"`
	Religion         string `json:"religion" form:"religion"`
	University       string `json:"university" form:"university"`
	Citizenship      string `json:"citizenship" form:"citizenship"`
	EmailPersonal    string `json:"email_personal" form:"email_personal"`
	PhoneMobile      string `json:"phone_mobile" form:"phone_mobile"`
	PhoneHome        string `json:"phone_home" form:"phone_home"`
	AddressKtp       string `json:"address_ktp" form:"address_ktp"`
	AddressDomicile  string `json:"address_domicile" form:"address_domicile"`
	StrNumber        string `json:"str_number" form:"str_number"`
	StrIssued        string `json:"str_issued" form:"str_issued"`
	StrExpired       string `json:"str_expired" form:"str_expired"`
	StrNote          string `json:"str_note" form:"str_note"`
	Profession       string `json:"profession" form:"profession"`
	Competency       string `json:"competency" form:"competency"`
	College          string `json:"college" form:"college"`
	Stage            string `json:"stage" form:"stage"`
	Dpjp             string `json:"dpjp" form:"dpjp"`
}

type WorkHistoryResponse struct {
	ID          int    `json:"id"`
	Position    string `json:"position"`
	Institution string `json:"institution"`
	StartDate   string `json:"start_date"`
	Status      string `json:"status"`
	Sip         string `json:"sip"`
	Location    string `json:"location"`
}

type CreateWorkHistoryRequest struct {
	Position    string `json:"position" form:"position" binding:"required"`
	Institution string `json:"institution" form:"institution" binding:"required"`
	StartDate   string `json:"start_date" form:"start_date"`
	Status      string `json:"status" form:"status"`
	Sip         string `json:"sip" form:"sip"`
	Location    string `json:"location" form:"location"`
}
