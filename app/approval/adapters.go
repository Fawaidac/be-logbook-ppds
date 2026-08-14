package approval

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// TindakanRepoAdapter implements TindakanRepository
type TindakanRepoAdapter struct {
	DB *sqlx.DB
}

func (r *TindakanRepoAdapter) FindByStatus(ctx context.Context, status string) ([]TindakanApprovalItem, error) {
	query := `
		SELECT id, user_username, mr_number, patient_name, diagnosis_label, plan_procedure, status, created_at
		FROM tindakans
		WHERE status = $1
		ORDER BY created_at DESC
	`

	var items []TindakanApprovalItem
	if err := r.DB.SelectContext(ctx, &items, query, status); err != nil && err != sql.ErrNoRows {
		return []TindakanApprovalItem{}, err
	}

	if items == nil {
		items = []TindakanApprovalItem{}
	}

	return items, nil
}

func (r *TindakanRepoAdapter) UpdateStatus(ctx context.Context, id int, status string) error {
	query := `UPDATE tindakans SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id)
	return err
}

// KegiatanIlmiahRepoAdapter implements KegiatanIlmiahRepository
type KegiatanIlmiahRepoAdapter struct {
	DB *sqlx.DB
}

func (r *KegiatanIlmiahRepoAdapter) FindByStatus(ctx context.Context, status string) ([]KegiatanApprovalItem, error) {
	query := `
		SELECT id, user_username, kategori, jenis_kegiatan, topik, status, created_at
		FROM kegiatan_ilmiahs
		WHERE status = $1
		ORDER BY created_at DESC
	`

	var items []KegiatanApprovalItem
	if err := r.DB.SelectContext(ctx, &items, query, status); err != nil && err != sql.ErrNoRows {
		return []KegiatanApprovalItem{}, err
	}

	if items == nil {
		items = []KegiatanApprovalItem{}
	}

	return items, nil
}

func (r *KegiatanIlmiahRepoAdapter) UpdateStatus(ctx context.Context, id int, status string) error {
	query := `UPDATE kegiatan_ilmiahs SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id)
	return err
}

// AktivitasKlinikRepoAdapter implements AktivitasKlinikRepository
type AktivitasKlinikRepoAdapter struct {
	DB *sqlx.DB
}

func (r *AktivitasKlinikRepoAdapter) FindByStatus(ctx context.Context, status string) ([]AktivitasKlinikApprovalItem, error) {
	query := `
		SELECT id, user_username, nama_aktivitas, tanggal, status, created_at
		FROM aktivitas_kliniks
		WHERE status = $1
		ORDER BY created_at DESC
	`

	var items []AktivitasKlinikApprovalItem
	if err := r.DB.SelectContext(ctx, &items, query, status); err != nil && err != sql.ErrNoRows {
		return []AktivitasKlinikApprovalItem{}, err
	}

	if items == nil {
		items = []AktivitasKlinikApprovalItem{}
	}

	return items, nil
}

func (r *AktivitasKlinikRepoAdapter) UpdateStatus(ctx context.Context, id int, status string) error {
	query := `UPDATE aktivitas_kliniks SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id)
	return err
}

// PendidikanEvaluasiRepoAdapter implements PendidikanEvaluasiRepository
type PendidikanEvaluasiRepoAdapter struct {
	DB *sqlx.DB
}

func (r *PendidikanEvaluasiRepoAdapter) FindByStatus(ctx context.Context, status string) ([]PendidikanEvaluasiApprovalItem, error) {
	query := `
		SELECT id, user_username, jenis_evaluasi, tanggal, status, created_at
		FROM pendidikan_evaluasis
		WHERE status = $1
		ORDER BY created_at DESC
	`

	var items []PendidikanEvaluasiApprovalItem
	if err := r.DB.SelectContext(ctx, &items, query, status); err != nil && err != sql.ErrNoRows {
		return []PendidikanEvaluasiApprovalItem{}, err
	}

	if items == nil {
		items = []PendidikanEvaluasiApprovalItem{}
	}

	return items, nil
}

func (r *PendidikanEvaluasiRepoAdapter) UpdateStatus(ctx context.Context, id int, status string) error {
	query := `UPDATE pendidikan_evaluasis SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id)
	return err
}
