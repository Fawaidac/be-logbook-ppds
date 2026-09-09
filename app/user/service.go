package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"be-logbook-ppds/pkg/email"
	"be-logbook-ppds/pkg/utils"
)

type Service interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*UserResponse, error)
	GetAllUsers(ctx context.Context) ([]UserResponse, error)
	GetUserByID(ctx context.Context, id int) (*UserResponse, error)
	UpdateUser(ctx context.Context, id int, req UpdateUserRequest) (*UserResponse, error)
	DeleteUser(ctx context.Context, id int) error

	CheckUniqueCredentials(ctx context.Context, username string, email string) error

	RegisterPPDS(ctx context.Context, req CreateRegistrationRequest, selfiePath, strPath, sipPath string) (*UserRegistrationResponse, error)
	GetRegistrations(ctx context.Context, status string) ([]UserRegistrationResponse, error)
	ApproveRegistration(ctx context.Context, id int) (*UserResponse, error)
	RejectRegistration(ctx context.Context, id int, reason string) error

	GetProfile(ctx context.Context, username string) (*UserProfileResponse, error)
	UpdateProfile(ctx context.Context, username string, req UpdateProfileRequest) (*UserProfileResponse, error)
	GetWorkHistories(ctx context.Context, username string) ([]WorkHistoryResponse, error)
	CreateWorkHistory(ctx context.Context, username string, req CreateWorkHistoryRequest) (*WorkHistoryResponse, error)
	UpdateWorkHistory(ctx context.Context, username string, id int, req CreateWorkHistoryRequest) (*WorkHistoryResponse, error)
	DeleteWorkHistory(ctx context.Context, username string, id int) error
}

type service struct {
	repo   Repository
	mailer email.Mailer
}

func NewService(repo Repository, mailer email.Mailer) Service {
	return &service{repo: repo, mailer: mailer}
}

func toUserResponse(u *User) *UserResponse {
	return &UserResponse{
		ID:           u.ID,
		Username:     u.Username,
		Name:         u.Name,
		Email:        u.Email,
		Role:         u.Role,
		NimNip:       u.NimNip,
		Jabatan:      u.Jabatan,
		ProgramStudi: u.ProgramStudi,
		CreatedAt:    u.CreatedAt,
	}
}

func toRegistrationResponse(reg *UserRegistration) *UserRegistrationResponse {
	return &UserRegistrationResponse{
		ID:              reg.ID,
		Name:            reg.Name,
		Nik:             reg.Nik,
		Str:             reg.Str,
		Sip:             reg.Sip,
		Username:        reg.Username,
		Email:           reg.Email,
		Specialty:       reg.Specialty,
		ProgramStudi:    reg.ProgramStudi,
		University:      reg.University,
		SelfiePath:      reg.SelfiePath,
		StrFilePath:     reg.StrFilePath,
		SipFilePath:     reg.SipFilePath,
		Status:          reg.Status,
		RejectionReason: reg.RejectionReason,
		CreatedAt:       reg.CreatedAt,
	}
}

func (s *service) CreateUser(ctx context.Context, req CreateUserRequest) (*UserResponse, error) {
	existingUser, _ := s.repo.FindByUsername(ctx, req.Username)
	if existingUser != nil {
		return nil, errors.New("username sudah digunakan")
	}

	hashedPassword, err := utils.HashPasswordArgon2(req.Password)
	if err != nil {
		return nil, errors.New("gagal memproses kata sandi")
	}

	u := &User{
		Username:     req.Username,
		Name:         req.Name,
		Email:        req.Email,
		Password:     hashedPassword,
		Role:         req.Role,
		NimNip:       req.NimNip,
		Jabatan:      req.Jabatan,
		ProgramStudi: req.ProgramStudi,
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}

	return toUserResponse(u), nil
}

func (s *service) GetAllUsers(ctx context.Context) ([]UserResponse, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var res []UserResponse
	for _, u := range users {
		res = append(res, *toUserResponse(&u))
	}
	return res, nil
}

func (s *service) GetUserByID(ctx context.Context, id int) (*UserResponse, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("pengguna tidak ditemukan")
	}
	return toUserResponse(u), nil
}

func (s *service) UpdateUser(ctx context.Context, id int, req UpdateUserRequest) (*UserResponse, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("pengguna tidak ditemukan")
	}

	u.Name = req.Name
	u.Email = req.Email
	u.Role = req.Role
	u.NimNip = req.NimNip
	u.Jabatan = req.Jabatan
	u.ProgramStudi = req.ProgramStudi

	if req.Password != "" {
		hashed, err := utils.HashPasswordArgon2(req.Password)
		if err != nil {
			return nil, errors.New("gagal memproses kata sandi baru")
		}
		u.Password = hashed
	}

	if err := s.repo.Update(ctx, u); err != nil {
		return nil, err
	}

	return toUserResponse(u), nil
}

func (s *service) DeleteUser(ctx context.Context, id int) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("pengguna tidak ditemukan")
	}

	return s.repo.Delete(ctx, id)
}

func (s *service) CheckUniqueCredentials(ctx context.Context, username string, email string) error {
    // check username
    if u, _ := s.repo.FindByUsername(ctx, username); u != nil {
        return errors.New("username sudah digunakan")
    }
    // check email
    if u, _ := s.repo.FindByEmail(ctx, email); u != nil {
        return errors.New("email sudah terdaftar")
    }
    return nil
}

func (s *service) RegisterPPDS(ctx context.Context, req CreateRegistrationRequest, selfiePath, strPath, sipPath string) (*UserRegistrationResponse, error) {
	// Validasi konfirmasi password
	if req.Password != req.PasswordConfirmation {
		return nil, errors.New("kata sandi dan konfirmasi kata sandi tidak cocok")
	}

	// Periksa keunikan username dan email
	if err := s.CheckUniqueCredentials(ctx, req.Username, req.Email); err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPasswordArgon2(req.Password)
	if err != nil {
		return nil, errors.New("gagal memproses kata sandi")
	}

	progStudi := req.ProgramStudi
	if progStudi == "" {
		progStudi = req.Specialty
	}

	reg := &UserRegistration{
		Name:         req.Name,
		Nik:          req.Nik,
		Str:          req.Str,
		Sip:          req.Sip,
		Username:     req.Username,
		Email:        req.Email,
		Password:     hashedPassword,
		Specialty:    req.Specialty,
		ProgramStudi: progStudi,
		University:   req.University,
		SelfiePath:   selfiePath,
		StrFilePath:  strPath,
		SipFilePath:  sipPath,
		Status:       "pending",
	}

	if err := s.repo.CreateRegistration(ctx, reg); err != nil {
		return nil, err
	}

	return toRegistrationResponse(reg), nil
}

func (s *service) GetRegistrations(ctx context.Context, status string) ([]UserRegistrationResponse, error) {
	regs, err := s.repo.FindAllRegistrations(ctx, status)
	if err != nil {
		return nil, err
	}

	var res []UserRegistrationResponse
	for _, r := range regs {
		res = append(res, *toRegistrationResponse(&r))
	}
	return res, nil
}

func (s *service) ApproveRegistration(ctx context.Context, id int) (*UserResponse, error) {
	reg, err := s.repo.FindRegistrationByID(ctx, id)
	if err != nil {
		return nil, errors.New("permintaan registrasi tidak ditemukan")
	}

	if reg.Status == "approved" {
		return nil, errors.New("registrasi ini sudah disetujui sebelumnya")
	}

	// Gunakan username yang sudah dipilih pendaftar; fallback ke NIK/email jika kosong
	username := strings.TrimSpace(reg.Username)
	if username == "" {
		username = strings.TrimSpace(reg.Nik)
	}
	if username == "" {
		username = strings.Split(reg.Email, "@")[0]
	}

	// Pastikan username unik (sangat jarang collision karena sudah dicek saat daftar)
	usernameBase := username
	counter := 1
	for {
		existing, _ := s.repo.FindByUsername(ctx, username)
		if existing == nil {
			break
		}
		username = fmt.Sprintf("%s_%d", usernameBase, counter)
		counter++
	}

	progStudi := reg.ProgramStudi
	if progStudi == "" {
		progStudi = reg.Specialty
	}

	nimNip := reg.Nik
	if nimNip == "" {
		nimNip = reg.Str
	}

	// Pakai password yang sudah di-hash saat pendaftaran (bukan default)
	u := &User{
		Username:     username,
		Name:         reg.Name,
		Email:        reg.Email,
		Password:     reg.Password,
		Role:         "residen",
		NimNip:       nimNip,
		Jabatan:      "Residen PPDS",
		ProgramStudi: progStudi,
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, errors.New("gagal membuat akun pengguna: " + err.Error())
	}

	_ = s.repo.UpdateRegistrationStatus(ctx, id, "approved", "")

	// Kirim notifikasi email persetujuan (tanpa menampilkan password)
	if s.mailer != nil {
		_ = s.mailer.SendApprovalEmail(reg.Email, reg.Name, username, "")
	}

	return toUserResponse(u), nil
}

func (s *service) RejectRegistration(ctx context.Context, id int, reason string) error {
	reg, err := s.repo.FindRegistrationByID(ctx, id)
	if err != nil {
		return errors.New("permintaan registrasi tidak ditemukan")
	}

	if reg.Status == "approved" {
		return errors.New("registrasi ini sudah disetujui, tidak dapat ditolak")
	}

	if strings.TrimSpace(reason) == "" {
		reason = "Dokumen/identitas tidak memenuhi persyaratan administrasi."
	}

	if err := s.repo.UpdateRegistrationStatus(ctx, id, "rejected", reason); err != nil {
		return err
	}

	// Kirim notifikasi email penolakan
	if s.mailer != nil {
		_ = s.mailer.SendRejectionEmail(reg.Email, reg.Name, reason)
	}

	return nil
}

// ------------------------- PROFILE (MASTER DATA) -------------------------

func formatDatePtr(t sql.NullTime) string {
	if !t.Valid {
		return ""
	}
	return t.Time.Format("2006-01-02")
}

func toProfileResponse(u *User) *UserProfileResponse {
	return &UserProfileResponse{
		Username:         u.Username,
		FullName:         u.Name,
		EmailInstitution: u.Email,
		Nik:              u.Nik,
		NimNip:           u.NimNip,
		ProgramStudi:     u.ProgramStudi,
		BirthPlace:       u.BirthPlace,
		BirthDate:        formatDatePtr(u.BirthDate),
		Gender:           u.Gender,
		BloodType:        u.BloodType,
		Religion:         u.Religion,
		University:       u.University,
		Citizenship:      u.Citizenship,
		EmailPersonal:    u.EmailPersonal,
		PhoneMobile:      u.PhoneMobile,
		PhoneHome:        u.PhoneHome,
		AddressKtp:       u.AddressKtp,
		AddressDomicile:  u.AddressDomicile,
		StrNumber:        u.StrNumber,
		StrIssued:        formatDatePtr(u.StrIssued),
		StrExpired:       formatDatePtr(u.StrExpired),
		StrNote:          u.StrNote,
		Profession:       u.Profession,
		Competency:       u.Competency,
		College:          u.College,
		Stage:            u.Stage,
		Dpjp:             u.Dpjp,
	}
}

func (s *service) GetProfile(ctx context.Context, username string) (*UserProfileResponse, error) {
	u, err := s.repo.FindProfileByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("profil pengguna tidak ditemukan")
	}
	return toProfileResponse(u), nil
}

func (s *service) UpdateProfile(ctx context.Context, username string, req UpdateProfileRequest) (*UserProfileResponse, error) {
	if strings.TrimSpace(req.FullName) == "" {
		return nil, errors.New("nama lengkap wajib diisi")
	}

	if err := s.repo.UpdateProfile(ctx, username, req); err != nil {
		return nil, errors.New("gagal memperbarui profil: " + err.Error())
	}

	u, err := s.repo.FindProfileByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("profil pengguna tidak ditemukan")
	}
	return toProfileResponse(u), nil
}

// ------------------------- WORK HISTORY -------------------------

func toWorkHistoryResponse(wh *WorkHistory) *WorkHistoryResponse {
	startDate := ""
	if wh.StartDate.Valid {
		startDate = wh.StartDate.Time.Format("2006-01-02")
	}
	return &WorkHistoryResponse{
		ID:          wh.ID,
		Position:    wh.Position,
		Institution: wh.Institution,
		StartDate:   startDate,
		Status:      wh.Status,
		Sip:         wh.Sip,
		Location:    wh.Location,
	}
}

func (s *service) GetWorkHistories(ctx context.Context, username string) ([]WorkHistoryResponse, error) {
	list, err := s.repo.FindWorkHistories(ctx, username)
	if err != nil {
		return []WorkHistoryResponse{}, err
	}

	res := []WorkHistoryResponse{}
	for i := range list {
		res = append(res, *toWorkHistoryResponse(&list[i]))
	}
	return res, nil
}

func (s *service) CreateWorkHistory(ctx context.Context, username string, req CreateWorkHistoryRequest) (*WorkHistoryResponse, error) {
	wh := &WorkHistory{
		UserUsername: username,
		Position:     req.Position,
		Institution:  req.Institution,
		Status:       req.Status,
		Sip:          req.Sip,
		Location:     req.Location,
	}
	if strings.TrimSpace(wh.Status) == "" {
		wh.Status = "Aktif"
	}

	if err := s.repo.CreateWorkHistory(ctx, wh); err != nil {
		return nil, errors.New("gagal menambah riwayat pekerjaan: " + err.Error())
	}
	return toWorkHistoryResponse(wh), nil
}

func (s *service) UpdateWorkHistory(ctx context.Context, username string, id int, req CreateWorkHistoryRequest) (*WorkHistoryResponse, error) {
	existing, err := s.repo.FindWorkHistoryByID(ctx, username, id)
	if err != nil {
		return nil, errors.New("riwayat pekerjaan tidak ditemukan")
	}

	existing.Position = req.Position
	existing.Institution = req.Institution
	existing.Status = req.Status
	existing.Sip = req.Sip
	existing.Location = req.Location
	if strings.TrimSpace(existing.Status) == "" {
		existing.Status = "Aktif"
	}

	if err := s.repo.UpdateWorkHistory(ctx, username, existing); err != nil {
		return nil, errors.New("gagal memperbarui riwayat pekerjaan")
	}
	return toWorkHistoryResponse(existing), nil
}

func (s *service) DeleteWorkHistory(ctx context.Context, username string, id int) error {
	if _, err := s.repo.FindWorkHistoryByID(ctx, username, id); err != nil {
		return errors.New("riwayat pekerjaan tidak ditemukan")
	}
	return s.repo.DeleteWorkHistory(ctx, username, id)
}
