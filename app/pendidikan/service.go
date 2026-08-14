package pendidikan

import (
	"context"
	"database/sql"
	"time"
)

type KompetensiService interface {
	CreateKompetensi(ctx context.Context, req CreateKompetensiRequest, username string) (*KompetensiResponse, error)
	GetAllKompetensi(ctx context.Context) ([]KompetensiResponse, error)
	GetKompetensiByUsername(ctx context.Context, username string) ([]KompetensiResponse, error)
	DeleteKompetensi(ctx context.Context, id int) error
}

type RotasiService interface {
	CreateRotasi(ctx context.Context, req CreateRotasiRequest, username string) (*RotasiResponse, error)
	GetAllRotasi(ctx context.Context) ([]RotasiResponse, error)
	GetRotasiByUsername(ctx context.Context, username string) ([]RotasiResponse, error)
	DeleteRotasi(ctx context.Context, id int) error
}

type MiniCexService interface {
	CreateMiniCex(ctx context.Context, req CreateMiniCexRequest, username string) (*MiniCexResponse, error)
	GetAllMiniCex(ctx context.Context) ([]MiniCexResponse, error)
	GetMiniCexByUsername(ctx context.Context, username string) ([]MiniCexResponse, error)
	DeleteMiniCex(ctx context.Context, id int) error
}

type DopsService interface {
	CreateDops(ctx context.Context, req CreateDopsRequest, username string) (*DopsResponse, error)
	GetAllDops(ctx context.Context) ([]DopsResponse, error)
	GetDopsByUsername(ctx context.Context, username string) ([]DopsResponse, error)
	DeleteDops(ctx context.Context, id int) error
}

type SeminarService interface {
	CreateSeminar(ctx context.Context, req CreateSeminarRequest, username string) (*SeminarResponse, error)
	GetAllSeminar(ctx context.Context) ([]SeminarResponse, error)
	GetSeminarByUsername(ctx context.Context, username string) ([]SeminarResponse, error)
	DeleteSeminar(ctx context.Context, id int) error
}

type CbdService interface {
	CreateCbd(ctx context.Context, req CreateCbdRequest, username string) (*CbdResponse, error)
	GetAllCbd(ctx context.Context) ([]CbdResponse, error)
	GetCbdByUsername(ctx context.Context, username string) ([]CbdResponse, error)
	DeleteCbd(ctx context.Context, id int) error
}

// Kompetensi Service
type kompetensiService struct {
	repo KompetensiRepository
}

func NewKompetensiService(repo KompetensiRepository) KompetensiService {
	return &kompetensiService{repo: repo}
}

var namaModulMap = map[string]string{
	"KMP-TRM-01": "Debridement & Irigasi Masif Fraktur Terbuka",
	"KMP-TRM-02": "ORIF Plating & Nailing Femur/Tibia",
	"KMP-PED-01": "Reduksi & Gips Serial Ponseti",
	"KMP-SPN-01": "Dekompresi & Pedicle Screw Lumbal",
	"KMP-REC-01": "Artroskopi Diagnostik Genu",
	"KMP-HND-01": "Repair Tendon Flexor & Ekstensor",
}

var domainModulMap = map[string]string{
	"KMP-TRM-01": "Trauma Muskoloskeletal",
	"KMP-TRM-02": "Trauma Muskoloskeletal",
	"KMP-PED-01": "Orthopedi Anak (Pediatric)",
	"KMP-SPN-01": "Spine & Vertebra",
	"KMP-REC-01": "Rekonstruksi & Sport",
	"KMP-HND-01": "Hand & Microsurgery",
}

func (s *kompetensiService) CreateKompetensi(ctx context.Context, req CreateKompetensiRequest, username string) (*KompetensiResponse, error) {
	nama := namaModulMap[req.Kode]
	if nama == "" {
		nama = "Modul Standar Kolegium"
	}

	domain := domainModulMap[req.Kode]
	if domain == "" {
		domain = "Trauma Muskoloskeletal"
	}

	k := &PendidikanKompetensi{
		UserUsername:  sql.NullString{String: username, Valid: username != ""},
		Kode:          req.Kode,
		Nama:          sql.NullString{String: nama, Valid: true},
		Domain:        sql.NullString{String: domain, Valid: true},
		LevelTarget:   req.LevelTarget,
		TargetLog:     10,
		AchievedLog:   req.AchievedLog,
		Evaluator:     sql.NullString{String: req.Evaluator, Valid: true},
		Status:        "Dalam Proses",
		TglVerifikasi: sql.NullString{String: time.Now().Format("02 Jan 2006"), Valid: true},
		Deskripsi:     sql.NullString{String: req.Deskripsi, Valid: true},
	}

	if err := s.repo.Create(ctx, k); err != nil {
		return nil, err
	}

	return s.toKompetensiResponse(k), nil
}

func (s *kompetensiService) GetAllKompetensi(ctx context.Context) ([]KompetensiResponse, error) {
	list, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []KompetensiResponse
	for _, k := range list {
		responses = append(responses, *s.toKompetensiResponse(&k))
	}
	return responses, nil
}

func (s *kompetensiService) GetKompetensiByUsername(ctx context.Context, username string) ([]KompetensiResponse, error) {
	list, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	var responses []KompetensiResponse
	for _, k := range list {
		responses = append(responses, *s.toKompetensiResponse(&k))
	}
	return responses, nil
}

func (s *kompetensiService) DeleteKompetensi(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *kompetensiService) toKompetensiResponse(k *PendidikanKompetensi) *KompetensiResponse {
	return &KompetensiResponse{
		ID:            k.ID,
		UserUsername:  k.UserUsername.String,
		Kode:          k.Kode,
		Nama:          k.Nama.String,
		Domain:        k.Domain.String,
		LevelTarget:   k.LevelTarget,
		TargetLog:     k.TargetLog,
		AchievedLog:   k.AchievedLog,
		Evaluator:     k.Evaluator.String,
		Status:        k.Status,
		TglVerifikasi: k.TglVerifikasi.String,
		Deskripsi:     k.Deskripsi.String,
	}
}

// Rotasi Service
type rotasiService struct {
	repo RotasiRepository
}

func NewRotasiService(repo RotasiRepository) RotasiService {
	return &rotasiService{repo: repo}
}

func (s *rotasiService) CreateRotasi(ctx context.Context, req CreateRotasiRequest, username string) (*RotasiResponse, error) {
	r := &PendidikanRotasi{
		UserUsername: sql.NullString{String: username, Valid: username != ""},
		Stase:        "Stase Poliklinik & Rawat Jalan Ortopedi",
		Lokasi:       sql.NullString{String: "RSUD dr. Soebandi", Valid: true},
		Periode:      sql.NullString{String: time.Now().Format("02 Jan 2006"), Valid: true},
		Pembimbing:   sql.NullString{String: req.Pembimbing, Valid: req.Pembimbing != ""},
		Kehadiran:    sql.NullString{String: "100%", Valid: true},
		Nilai:        sql.NullString{String: "Belum Ada", Valid: true},
		Status:       "Berlangsung",
		Tanggal:      sql.NullString{String: time.Now().Format("02 Jan 2006"), Valid: true},
		Catatan:      sql.NullString{String: req.Catatan, Valid: req.Catatan != ""},
	}

	if err := s.repo.Create(ctx, r); err != nil {
		return nil, err
	}

	return s.toRotasiResponse(r), nil
}

func (s *rotasiService) GetAllRotasi(ctx context.Context) ([]RotasiResponse, error) {
	list, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []RotasiResponse
	for _, r := range list {
		responses = append(responses, *s.toRotasiResponse(&r))
	}
	return responses, nil
}

func (s *rotasiService) GetRotasiByUsername(ctx context.Context, username string) ([]RotasiResponse, error) {
	list, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	var responses []RotasiResponse
	for _, r := range list {
		responses = append(responses, *s.toRotasiResponse(&r))
	}
	return responses, nil
}

func (s *rotasiService) DeleteRotasi(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *rotasiService) toRotasiResponse(r *PendidikanRotasi) *RotasiResponse {
	return &RotasiResponse{
		ID:           r.ID,
		UserUsername: r.UserUsername.String,
		Stase:        r.Stase,
		Lokasi:       r.Lokasi.String,
		Periode:      r.Periode.String,
		Pembimbing:   r.Pembimbing.String,
		Kehadiran:    r.Kehadiran.String,
		Nilai:        r.Nilai.String,
		Status:       r.Status,
		Tanggal:      r.Tanggal.String,
		Catatan:      r.Catatan.String,
	}
}

// MiniCex Service
type miniCexService struct {
	repo MiniCexRepository
}

func NewMiniCexService(repo MiniCexRepository) MiniCexService {
	return &miniCexService{repo: repo}
}

func (s *miniCexService) CreateMiniCex(ctx context.Context, req CreateMiniCexRequest, username string) (*MiniCexResponse, error) {
	m := &PendidikanMiniCex{
		UserUsername: sql.NullString{String: username, Valid: username != ""},
		Pasien:       sql.NullString{String: req.Pasien, Valid: req.Pasien != ""},
		Fokus:        sql.NullString{String: req.Fokus, Valid: req.Fokus != ""},
		Kasus:        sql.NullString{String: req.Kasus, Valid: req.Kasus != ""},
		Evaluator:    sql.NullString{String: req.Evaluator, Valid: req.Evaluator != ""},
		Skor:         sql.NullString{String: "85.0", Valid: true},
		Status:       "Menunggu Validasi",
		Tanggal:      sql.NullString{String: time.Now().Format("02 Jan 2006"), Valid: true},
		Catatan:      sql.NullString{String: req.Catatan, Valid: req.Catatan != ""},
	}

	if err := s.repo.Create(ctx, m); err != nil {
		return nil, err
	}

	return s.toMiniCexResponse(m), nil
}

func (s *miniCexService) GetAllMiniCex(ctx context.Context) ([]MiniCexResponse, error) {
	list, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []MiniCexResponse
	for _, m := range list {
		responses = append(responses, *s.toMiniCexResponse(&m))
	}
	return responses, nil
}

func (s *miniCexService) GetMiniCexByUsername(ctx context.Context, username string) ([]MiniCexResponse, error) {
	list, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	var responses []MiniCexResponse
	for _, m := range list {
		responses = append(responses, *s.toMiniCexResponse(&m))
	}
	return responses, nil
}

func (s *miniCexService) DeleteMiniCex(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *miniCexService) toMiniCexResponse(m *PendidikanMiniCex) *MiniCexResponse {
	return &MiniCexResponse{
		ID:           m.ID,
		UserUsername: m.UserUsername.String,
		Pasien:       m.Pasien.String,
		Fokus:        m.Fokus.String,
		Kasus:        m.Kasus.String,
		Evaluator:    m.Evaluator.String,
		Skor:         m.Skor.String,
		Status:       m.Status,
		Tanggal:      m.Tanggal.String,
		Catatan:      m.Catatan.String,
	}
}

// Dops Service
type dopsService struct {
	repo DopsRepository
}

func NewDopsService(repo DopsRepository) DopsService {
	return &dopsService{repo: repo}
}

func (s *dopsService) CreateDops(ctx context.Context, req CreateDopsRequest, username string) (*DopsResponse, error) {
	d := &PendidikanDops{
		UserUsername: sql.NullString{String: username, Valid: username != ""},
		Prosedur:     sql.NullString{String: req.Prosedur, Valid: req.Prosedur != ""},
		Kategori:     sql.NullString{String: "Tindakan Klinis", Valid: true},
		Kesulitan:    sql.NullString{String: req.Kesulitan, Valid: req.Kesulitan != ""},
		Supervisor:   sql.NullString{String: req.Supervisor, Valid: req.Supervisor != ""},
		Skor:         sql.NullString{String: "88.0", Valid: true},
		Status:       "Menunggu Validasi",
		Tanggal:      sql.NullString{String: time.Now().Format("02 Jan 2006"), Valid: true},
		Catatan:      sql.NullString{String: req.Catatan, Valid: req.Catatan != ""},
	}

	if err := s.repo.Create(ctx, d); err != nil {
		return nil, err
	}

	return s.toDopsResponse(d), nil
}

func (s *dopsService) GetAllDops(ctx context.Context) ([]DopsResponse, error) {
	list, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []DopsResponse
	for _, d := range list {
		responses = append(responses, *s.toDopsResponse(&d))
	}
	return responses, nil
}

func (s *dopsService) GetDopsByUsername(ctx context.Context, username string) ([]DopsResponse, error) {
	list, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	var responses []DopsResponse
	for _, d := range list {
		responses = append(responses, *s.toDopsResponse(&d))
	}
	return responses, nil
}

func (s *dopsService) DeleteDops(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *dopsService) toDopsResponse(d *PendidikanDops) *DopsResponse {
	return &DopsResponse{
		ID:           d.ID,
		UserUsername: d.UserUsername.String,
		Prosedur:     d.Prosedur.String,
		Kategori:     d.Kategori.String,
		Kesulitan:    d.Kesulitan.String,
		Supervisor:   d.Supervisor.String,
		Skor:         d.Skor.String,
		Status:       d.Status,
		Tanggal:      d.Tanggal.String,
		Catatan:      d.Catatan.String,
	}
}

// Seminar Service
type seminarService struct {
	repo SeminarRepository
}

func NewSeminarService(repo SeminarRepository) SeminarService {
	return &seminarService{repo: repo}
}

func (s *seminarService) CreateSeminar(ctx context.Context, req CreateSeminarRequest, username string) (*SeminarResponse, error) {
	sem := &PendidikanSeminar{
		UserUsername: sql.NullString{String: username, Valid: username != ""},
		Judul:        sql.NullString{String: req.Judul, Valid: req.Judul != ""},
		Jenis:        sql.NullString{String: req.Jenis, Valid: req.Jenis != ""},
		Narasumber:   sql.NullString{String: req.Narasumber, Valid: req.Narasumber != ""},
		Ruang:        sql.NullString{String: "Ruang Pertemuan Departemen", Valid: true},
		Skor:         sql.NullString{String: "86.0 (A)", Valid: true},
		Status:       "Menunggu Validasi",
		Tanggal:      sql.NullString{String: time.Now().Format("02 Jan 2006, 15:04") + " WIB", Valid: true},
		Catatan:      sql.NullString{String: req.Catatan, Valid: req.Catatan != ""},
	}

	if err := s.repo.Create(ctx, sem); err != nil {
		return nil, err
	}

	return s.toSeminarResponse(sem), nil
}

func (s *seminarService) GetAllSeminar(ctx context.Context) ([]SeminarResponse, error) {
	list, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []SeminarResponse
	for _, sem := range list {
		responses = append(responses, *s.toSeminarResponse(&sem))
	}
	return responses, nil
}

func (s *seminarService) GetSeminarByUsername(ctx context.Context, username string) ([]SeminarResponse, error) {
	list, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	var responses []SeminarResponse
	for _, sem := range list {
		responses = append(responses, *s.toSeminarResponse(&sem))
	}
	return responses, nil
}

func (s *seminarService) DeleteSeminar(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *seminarService) toSeminarResponse(sem *PendidikanSeminar) *SeminarResponse {
	return &SeminarResponse{
		ID:           sem.ID,
		UserUsername: sem.UserUsername.String,
		Judul:        sem.Judul.String,
		Jenis:        sem.Jenis.String,
		Narasumber:   sem.Narasumber.String,
		Ruang:        sem.Ruang.String,
		Skor:         sem.Skor.String,
		Status:       sem.Status,
		Tanggal:      sem.Tanggal.String,
		Catatan:      sem.Catatan.String,
	}
}

// CBD Service
type cbdService struct {
	repo CbdRepository
}

func NewCbdService(repo CbdRepository) CbdService {
	return &cbdService{repo: repo}
}

func (s *cbdService) CreateCbd(ctx context.Context, req CreateCbdRequest, username string) (*CbdResponse, error) {
	c := &PendidikanCbd{
		UserUsername: sql.NullString{String: username, Valid: username != ""},
		Pasien:       sql.NullString{String: req.Pasien, Valid: req.Pasien != ""},
		Topik:        sql.NullString{String: req.Topik, Valid: req.Topik != ""},
		Kategori:     sql.NullString{String: req.Kategori, Valid: req.Kategori != ""},
		Pembimbing:   sql.NullString{String: req.Pembimbing, Valid: req.Pembimbing != ""},
		Kompleksitas: sql.NullString{String: "Sedang", Valid: true},
		Skor:         sql.NullString{String: "85.0", Valid: true},
		Status:       "Menunggu Validasi",
		Tanggal:      sql.NullString{String: time.Now().Format("02 Jan 2006"), Valid: true},
		Catatan:      sql.NullString{String: req.Catatan, Valid: req.Catatan != ""},
	}

	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}

	return s.toCbdResponse(c), nil
}

func (s *cbdService) GetAllCbd(ctx context.Context) ([]CbdResponse, error) {
	list, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []CbdResponse
	for _, c := range list {
		responses = append(responses, *s.toCbdResponse(&c))
	}
	return responses, nil
}

func (s *cbdService) GetCbdByUsername(ctx context.Context, username string) ([]CbdResponse, error) {
	list, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	var responses []CbdResponse
	for _, c := range list {
		responses = append(responses, *s.toCbdResponse(&c))
	}
	return responses, nil
}

func (s *cbdService) DeleteCbd(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *cbdService) toCbdResponse(c *PendidikanCbd) *CbdResponse {
	return &CbdResponse{
		ID:            c.ID,
		UserUsername:  c.UserUsername.String,
		Pasien:        c.Pasien.String,
		Topik:         c.Topik.String,
		Kategori:      c.Kategori.String,
		Pembimbing:    c.Pembimbing.String,
		Kompleksitas:  c.Kompleksitas.String,
		Skor:          c.Skor.String,
		Status:        c.Status,
		Tanggal:       c.Tanggal.String,
		Catatan:       c.Catatan.String,
	}
}
