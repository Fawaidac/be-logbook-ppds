package pendidikan

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type KompetensiRepository interface {
	Create(ctx context.Context, k *PendidikanKompetensi) error
	FindAll(ctx context.Context) ([]PendidikanKompetensi, error)
	FindByUsername(ctx context.Context, username string) ([]PendidikanKompetensi, error)
	FindByID(ctx context.Context, id int) (*PendidikanKompetensi, error)
	Delete(ctx context.Context, id int) error
}

type RotasiRepository interface {
	Create(ctx context.Context, r *PendidikanRotasi) error
	FindAll(ctx context.Context) ([]PendidikanRotasi, error)
	FindByUsername(ctx context.Context, username string) ([]PendidikanRotasi, error)
	FindByID(ctx context.Context, id int) (*PendidikanRotasi, error)
	Delete(ctx context.Context, id int) error
}

type MiniCexRepository interface {
	Create(ctx context.Context, m *PendidikanMiniCex) error
	FindAll(ctx context.Context) ([]PendidikanMiniCex, error)
	FindByUsername(ctx context.Context, username string) ([]PendidikanMiniCex, error)
	FindByID(ctx context.Context, id int) (*PendidikanMiniCex, error)
	Delete(ctx context.Context, id int) error
}

type DopsRepository interface {
	Create(ctx context.Context, d *PendidikanDops) error
	FindAll(ctx context.Context) ([]PendidikanDops, error)
	FindByUsername(ctx context.Context, username string) ([]PendidikanDops, error)
	FindByID(ctx context.Context, id int) (*PendidikanDops, error)
	Delete(ctx context.Context, id int) error
}

type SeminarRepository interface {
	Create(ctx context.Context, s *PendidikanSeminar) error
	FindAll(ctx context.Context) ([]PendidikanSeminar, error)
	FindByUsername(ctx context.Context, username string) ([]PendidikanSeminar, error)
	FindByID(ctx context.Context, id int) (*PendidikanSeminar, error)
	Delete(ctx context.Context, id int) error
}

type CbdRepository interface {
	Create(ctx context.Context, c *PendidikanCbd) error
	FindAll(ctx context.Context) ([]PendidikanCbd, error)
	FindByUsername(ctx context.Context, username string) ([]PendidikanCbd, error)
	FindByID(ctx context.Context, id int) (*PendidikanCbd, error)
	Delete(ctx context.Context, id int) error
}

// Kompetensi Repository Implementation
type kompetensiRepo struct {
	db *sqlx.DB
}

func NewKompetensiRepository(db *sqlx.DB) KompetensiRepository {
	return &kompetensiRepo{db: db}
}

func (r *kompetensiRepo) Create(ctx context.Context, k *PendidikanKompetensi) error {
	query := `
		INSERT INTO pendidikan_kompetensi (user_username, kode, nama, domain, level_target, target_log, achieved_log, evaluator, status, tgl_verifikasi, deskripsi, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW())
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		k.UserUsername, k.Kode, k.Nama, k.Domain, k.LevelTarget, k.TargetLog, k.AchievedLog, k.Evaluator, k.Status, k.TglVerifikasi, k.Deskripsi,
	).Scan(&k.ID, &k.CreatedAt, &k.UpdatedAt)
}

func (r *kompetensiRepo) FindAll(ctx context.Context) ([]PendidikanKompetensi, error) {
	var list []PendidikanKompetensi
	query := `SELECT * FROM pendidikan_kompetensi ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query)
	return list, err
}

func (r *kompetensiRepo) FindByUsername(ctx context.Context, username string) ([]PendidikanKompetensi, error) {
	var list []PendidikanKompetensi
	query := `SELECT * FROM pendidikan_kompetensi WHERE user_username = $1 ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query, username)
	return list, err
}

func (r *kompetensiRepo) FindByID(ctx context.Context, id int) (*PendidikanKompetensi, error) {
	var k PendidikanKompetensi
	query := `SELECT * FROM pendidikan_kompetensi WHERE id = $1`
	err := r.db.GetContext(ctx, &k, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &k, err
}

func (r *kompetensiRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM pendidikan_kompetensi WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// Rotasi Repository Implementation
type rotasiRepo struct {
	db *sqlx.DB
}

func NewRotasiRepository(db *sqlx.DB) RotasiRepository {
	return &rotasiRepo{db: db}
}

func (r *rotasiRepo) Create(ctx context.Context, ro *PendidikanRotasi) error {
	query := `
		INSERT INTO pendidikan_rotasi (user_username, stase, lokasi, periode, pembimbing, kehadiran, nilai, status, tanggal, catatan, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW())
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		ro.UserUsername, ro.Stase, ro.Lokasi, ro.Periode, ro.Pembimbing, ro.Kehadiran, ro.Nilai, ro.Status, ro.Tanggal, ro.Catatan,
	).Scan(&ro.ID, &ro.CreatedAt, &ro.UpdatedAt)
}

func (r *rotasiRepo) FindAll(ctx context.Context) ([]PendidikanRotasi, error) {
	var list []PendidikanRotasi
	query := `SELECT * FROM pendidikan_rotasi ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query)
	return list, err
}

func (r *rotasiRepo) FindByUsername(ctx context.Context, username string) ([]PendidikanRotasi, error) {
	var list []PendidikanRotasi
	query := `SELECT * FROM pendidikan_rotasi WHERE user_username = $1 ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query, username)
	return list, err
}

func (r *rotasiRepo) FindByID(ctx context.Context, id int) (*PendidikanRotasi, error) {
	var ro PendidikanRotasi
	query := `SELECT * FROM pendidikan_rotasi WHERE id = $1`
	err := r.db.GetContext(ctx, &ro, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &ro, err
}

func (r *rotasiRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM pendidikan_rotasi WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// MiniCex Repository Implementation
type miniCexRepo struct {
	db *sqlx.DB
}

func NewMiniCexRepository(db *sqlx.DB) MiniCexRepository {
	return &miniCexRepo{db: db}
}

func (r *miniCexRepo) Create(ctx context.Context, m *PendidikanMiniCex) error {
	query := `
		INSERT INTO pendidikan_mini_cex (user_username, pasien, fokus, kasus, evaluator, skor, status, tanggal, catatan, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		m.UserUsername, m.Pasien, m.Fokus, m.Kasus, m.Evaluator, m.Skor, m.Status, m.Tanggal, m.Catatan,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
}

func (r *miniCexRepo) FindAll(ctx context.Context) ([]PendidikanMiniCex, error) {
	var list []PendidikanMiniCex
	query := `SELECT * FROM pendidikan_mini_cex ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query)
	return list, err
}

func (r *miniCexRepo) FindByUsername(ctx context.Context, username string) ([]PendidikanMiniCex, error) {
	var list []PendidikanMiniCex
	query := `SELECT * FROM pendidikan_mini_cex WHERE user_username = $1 ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query, username)
	return list, err
}

func (r *miniCexRepo) FindByID(ctx context.Context, id int) (*PendidikanMiniCex, error) {
	var m PendidikanMiniCex
	query := `SELECT * FROM pendidikan_mini_cex WHERE id = $1`
	err := r.db.GetContext(ctx, &m, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

func (r *miniCexRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM pendidikan_mini_cex WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// Dops Repository Implementation
type dopsRepo struct {
	db *sqlx.DB
}

func NewDopsRepository(db *sqlx.DB) DopsRepository {
	return &dopsRepo{db: db}
}

func (r *dopsRepo) Create(ctx context.Context, d *PendidikanDops) error {
	query := `
		INSERT INTO pendidikan_dops (user_username, prosedur, kategori, kesulitan, supervisor, skor, status, tanggal, catatan, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		d.UserUsername, d.Prosedur, d.Kategori, d.Kesulitan, d.Supervisor, d.Skor, d.Status, d.Tanggal, d.Catatan,
	).Scan(&d.ID, &d.CreatedAt, &d.UpdatedAt)
}

func (r *dopsRepo) FindAll(ctx context.Context) ([]PendidikanDops, error) {
	var list []PendidikanDops
	query := `SELECT * FROM pendidikan_dops ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query)
	return list, err
}

func (r *dopsRepo) FindByUsername(ctx context.Context, username string) ([]PendidikanDops, error) {
	var list []PendidikanDops
	query := `SELECT * FROM pendidikan_dops WHERE user_username = $1 ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query, username)
	return list, err
}

func (r *dopsRepo) FindByID(ctx context.Context, id int) (*PendidikanDops, error) {
	var d PendidikanDops
	query := `SELECT * FROM pendidikan_dops WHERE id = $1`
	err := r.db.GetContext(ctx, &d, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &d, err
}

func (r *dopsRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM pendidikan_dops WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// Seminar Repository Implementation
type seminarRepo struct {
	db *sqlx.DB
}

func NewSeminarRepository(db *sqlx.DB) SeminarRepository {
	return &seminarRepo{db: db}
}

func (r *seminarRepo) Create(ctx context.Context, s *PendidikanSeminar) error {
	query := `
		INSERT INTO pendidikan_seminar (user_username, judul, jenis, narasumber, ruang, skor, status, tanggal, catatan, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		s.UserUsername, s.Judul, s.Jenis, s.Narasumber, s.Ruang, s.Skor, s.Status, s.Tanggal, s.Catatan,
	).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)
}

func (r *seminarRepo) FindAll(ctx context.Context) ([]PendidikanSeminar, error) {
	var list []PendidikanSeminar
	query := `SELECT * FROM pendidikan_seminar ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query)
	return list, err
}

func (r *seminarRepo) FindByUsername(ctx context.Context, username string) ([]PendidikanSeminar, error) {
	var list []PendidikanSeminar
	query := `SELECT * FROM pendidikan_seminar WHERE user_username = $1 ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query, username)
	return list, err
}

func (r *seminarRepo) FindByID(ctx context.Context, id int) (*PendidikanSeminar, error) {
	var s PendidikanSeminar
	query := `SELECT * FROM pendidikan_seminar WHERE id = $1`
	err := r.db.GetContext(ctx, &s, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &s, err
}

func (r *seminarRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM pendidikan_seminar WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// CBD Repository Implementation
type cbdRepo struct {
	db *sqlx.DB
}

func NewCbdRepository(db *sqlx.DB) CbdRepository {
	return &cbdRepo{db: db}
}

func (r *cbdRepo) Create(ctx context.Context, c *PendidikanCbd) error {
	query := `
		INSERT INTO pendidikan_cbd (user_username, pasien, topik, kategori, pembimbing, kompleksitas, skor, status, tanggal, catatan, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW())
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		c.UserUsername, c.Pasien, c.Topik, c.Kategori, c.Pembimbing, c.Kompleksitas, c.Skor, c.Status, c.Tanggal, c.Catatan,
	).Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt)
}

func (r *cbdRepo) FindAll(ctx context.Context) ([]PendidikanCbd, error) {
	var list []PendidikanCbd
	query := `SELECT * FROM pendidikan_cbd ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query)
	return list, err
}

func (r *cbdRepo) FindByUsername(ctx context.Context, username string) ([]PendidikanCbd, error) {
	var list []PendidikanCbd
	query := `SELECT * FROM pendidikan_cbd WHERE user_username = $1 ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &list, query, username)
	return list, err
}

func (r *cbdRepo) FindByID(ctx context.Context, id int) (*PendidikanCbd, error) {
	var c PendidikanCbd
	query := `SELECT * FROM pendidikan_cbd WHERE id = $1`
	err := r.db.GetContext(ctx, &c, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &c, err
}

func (r *cbdRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM pendidikan_cbd WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
