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
// Nama tabel di migration adalah "kegiatan_ilmiah" (tanpa 's')
type KegiatanIlmiahRepoAdapter struct {
	DB *sqlx.DB
}

func (r *KegiatanIlmiahRepoAdapter) FindByStatus(ctx context.Context, status string) ([]KegiatanApprovalItem, error) {
	query := `
        SELECT id, user_username, COALESCE(ppds_name, '') AS ppds_name, kategori, COALESCE(jenis_kegiatan, '') AS jenis_kegiatan, COALESCE(topik, '') AS topik, COALESCE(tanggal_mulai::TEXT, '') AS tanggal_mulai, COALESCE(lokasi_tipe, '') AS lokasi_tipe, COALESCE(lokasi_detail, '') AS lokasi_detail, COALESCE(sebagai, '') AS sebagai, status, created_at
        FROM kegiatan_ilmiah
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
	query := `UPDATE kegiatan_ilmiah SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id)
	return err
}

// AktivitasKlinikRepoAdapter implements AktivitasKlinikRepository
// Kolom yang tersedia: tindakan (bukan nama_aktivitas)
type AktivitasKlinikRepoAdapter struct {
	DB *sqlx.DB
}

func (r *AktivitasKlinikRepoAdapter) FindByStatus(ctx context.Context, status string) ([]AktivitasKlinikApprovalItem, error) {
	query := `
		SELECT id, user_username, tindakan AS nama_aktivitas, tanggal::TEXT AS tanggal, status, created_at
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
// Nama tabel di migration adalah "pendidikan_evaluasi" (tanpa 's')
// Kolom yang tersedia: kategori (bukan jenis_evaluasi)
type PendidikanEvaluasiRepoAdapter struct {
	DB *sqlx.DB
}

func (r *PendidikanEvaluasiRepoAdapter) FindByStatus(ctx context.Context, status string) ([]PendidikanEvaluasiApprovalItem, error) {
	query := `
		SELECT id, user_username, kategori AS jenis_evaluasi, tanggal, status, created_at
		FROM pendidikan_evaluasi
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
	query := `UPDATE pendidikan_evaluasi SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id)
	return err
}
