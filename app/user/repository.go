package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindAll(ctx context.Context) ([]User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindProfileByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id int) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
	UpdateProfile(ctx context.Context, username string, req UpdateProfileRequest) error
	CreateRegistration(ctx context.Context, reg *UserRegistration) error
	FindAllRegistrations(ctx context.Context, status string) ([]UserRegistration, error)
	FindRegistrationByID(ctx context.Context, id int) (*UserRegistration, error)
	UpdateRegistrationStatus(ctx context.Context, id int, status string, rejectionReason string) error
	FindWorkHistories(ctx context.Context, username string) ([]WorkHistory, error)
	FindWorkHistoryByID(ctx context.Context, username string, id int) (*WorkHistory, error)
	CreateWorkHistory(ctx context.Context, wh *WorkHistory) error
	UpdateWorkHistory(ctx context.Context, username string, wh *WorkHistory) error
	DeleteWorkHistory(ctx context.Context, username string, id int) error
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

// profileColumns adalah daftar kolom profil untuk SELECT (dengan COALESCE agar aman dari NULL).
const profileColumns = `id, username, name, email, COALESCE(password, '') AS password, role, COALESCE(jabatan, '') AS jabatan,
	COALESCE(nim_nip, '') AS nim_nip, COALESCE(program_studi, '') AS program_studi,
	COALESCE(nik, '') AS nik, COALESCE(birth_place, '') AS birth_place, birth_date,
	COALESCE(gender, '') AS gender, COALESCE(blood_type, '') AS blood_type, COALESCE(religion, '') AS religion,
	COALESCE(university, '') AS university, COALESCE(citizenship, 'Warga Negara Indonesia (WNI)') AS citizenship,
	COALESCE(email_personal, '') AS email_personal, COALESCE(phone_mobile, '') AS phone_mobile,
	COALESCE(phone_home, '') AS phone_home, COALESCE(address_ktp, '') AS address_ktp,
	COALESCE(address_domicile, '') AS address_domicile, COALESCE(str_number, '') AS str_number,
	str_issued, str_expired, COALESCE(str_note, '') AS str_note, COALESCE(profession, '') AS profession,
	COALESCE(competency, '') AS competency, COALESCE(college, '') AS college, COALESCE(stage, '') AS stage,
	COALESCE(dpjp, '') AS dpjp, created_at`

func (r *repository) FindProfileByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	query := `SELECT ` + profileColumns + ` FROM users WHERE username = $1`
	err := r.db.GetContext(ctx, &user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) UpdateProfile(ctx context.Context, username string, req UpdateProfileRequest) error {
	query := `UPDATE users SET
		name = $1,
		email = NULLIF($2, ''),
		nik = $3,
		birth_place = $4,
		birth_date = NULLIF($5, '')::date,
		gender = $6,
		blood_type = $7,
		religion = $8,
		university = $9,
		citizenship = $10,
		email_personal = $11,
		phone_mobile = $12,
		phone_home = $13,
		address_ktp = $14,
		address_domicile = $15,
		str_number = $16,
		str_issued = NULLIF($17, '')::date,
		str_expired = NULLIF($18, '')::date,
		str_note = $19,
		profession = $20,
		competency = $21,
		college = $22,
		stage = $23,
		dpjp = $24
	WHERE username = $25`

	_, err := r.db.ExecContext(ctx, query,
		req.FullName, req.EmailInstitution, req.Nik, req.BirthPlace, req.BirthDate,
		req.Gender, req.BloodType, req.Religion, req.University, req.Citizenship,
		req.EmailPersonal, req.PhoneMobile, req.PhoneHome, req.AddressKtp, req.AddressDomicile,
		req.StrNumber, req.StrIssued, req.StrExpired, req.StrNote,
		req.Profession, req.Competency, req.College, req.Stage, req.Dpjp,
		username,
	)
	return err
}

// ------------------------- WORK HISTORY -------------------------

const workHistoryColumns = `id, user_username, position, institution, start_date, COALESCE(status, 'Aktif') AS status, COALESCE(sip, '') AS sip, COALESCE(location, '') AS location, created_at, updated_at`

func (r *repository) FindWorkHistories(ctx context.Context, username string) ([]WorkHistory, error) {
	list := []WorkHistory{}
	query := `SELECT ` + workHistoryColumns + ` FROM user_work_histories WHERE user_username = $1 ORDER BY id DESC`
	if err := r.db.SelectContext(ctx, &list, query, username); err != nil {
		return []WorkHistory{}, err
	}
	return list, nil
}

func (r *repository) FindWorkHistoryByID(ctx context.Context, username string, id int) (*WorkHistory, error) {
	var wh WorkHistory
	query := `SELECT ` + workHistoryColumns + ` FROM user_work_histories WHERE user_username = $1 AND id = $2`
	if err := r.db.GetContext(ctx, &wh, query, username, id); err != nil {
		return nil, err
	}
	return &wh, nil
}

func (r *repository) CreateWorkHistory(ctx context.Context, wh *WorkHistory) error {
	query := `INSERT INTO user_work_histories (user_username, position, institution, start_date, status, sip, location)
		VALUES ($1, $2, $3, NULLIF($4, '')::date, $5, $6, $7) RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		wh.UserUsername, wh.Position, wh.Institution, wh.StartDate, wh.Status, wh.Sip, wh.Location,
	).Scan(&wh.ID, &wh.CreatedAt, &wh.UpdatedAt)
}

func (r *repository) UpdateWorkHistory(ctx context.Context, username string, wh *WorkHistory) error {
	query := `UPDATE user_work_histories SET
		position = $1, institution = $2, start_date = NULLIF($3, '')::date, status = $4, sip = $5, location = $6, updated_at = NOW()
	WHERE user_username = $7 AND id = $8`
	res, err := r.db.ExecContext(ctx, query,
		wh.Position, wh.Institution, wh.StartDate, wh.Status, wh.Sip, wh.Location, username, wh.ID,
	)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return context.Canceled
	}
	return nil
}

func (r *repository) DeleteWorkHistory(ctx context.Context, username string, id int) error {
	query := `DELETE FROM user_work_histories WHERE user_username = $1 AND id = $2`
	_, err := r.db.ExecContext(ctx, query, username, id)
	return err
}
