package approval

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// TindakanRepoAdapter implements TindakanRepository
type TindakanRepoAdapter struct {
	DB *sqlx.DB
}

func (r *TindakanRepoAdapter) FindByStatus(ctx context.Context, status, supervisorName string) ([]TindakanApprovalItem, error) {
	// supervisorName kosong = lihat semua (admin); jika terisi,
	// hanya tindakan dengan DPJP yang memilih supervisor tsb (relasi 1 residen -> 1 supervisor).
	// Subquery agar PostgreSQL mengembalikan tipe TEXT, bukan DATE/TIMESTAMP asli.
	query := `
		SELECT id, user_username, mr_number, visit_date, patient_name, gender, birth_date,
		       division, diagnosis_label, procedure_code, plan_procedure, activity,
		       procedure_date, room, role_label, kemandirian, clinical_note,
		       supervisor_name, status, created_at
		FROM (
			SELECT id,
			       COALESCE(user_username, '')::TEXT AS user_username,
			       COALESCE(mr_number, '')::TEXT AS mr_number,
			       COALESCE(visit_date::TEXT, '')::TEXT AS visit_date,
			       COALESCE(patient_name, '')::TEXT AS patient_name,
			       COALESCE(gender::TEXT, '')::TEXT AS gender,
			       COALESCE(birth_date::TEXT, '')::TEXT AS birth_date,
			       COALESCE(division, '')::TEXT AS division,
			       COALESCE(diagnosis_label, '')::TEXT AS diagnosis_label,
			       COALESCE(procedure_code, '')::TEXT AS procedure_code,
			       COALESCE(plan_procedure, '')::TEXT AS plan_procedure,
			       COALESCE(activity, '')::TEXT AS activity,
			       COALESCE(procedure_date::TEXT, '')::TEXT AS procedure_date,
			       COALESCE(room, '')::TEXT AS room,
			       COALESCE(role, '')::TEXT AS role_label,
			       COALESCE(kemandirian::TEXT, '')::TEXT AS kemandirian,
			       COALESCE(clinical_note, '')::TEXT AS clinical_note,
			       COALESCE(supervisor_name, '')::TEXT AS supervisor_name,
			       COALESCE(status::TEXT, '')::TEXT AS status,
			       COALESCE(created_at::TEXT, '')::TEXT AS created_at
			FROM tindakans
			WHERE status = $1::status_enum
			  AND ($2 = '' OR supervisor_name = $2)
		) sub
		ORDER BY created_at DESC
	`

	var items []TindakanApprovalItem
	if err := r.DB.SelectContext(ctx, &items, query, status, supervisorName); err != nil && err != sql.ErrNoRows {
		fmt.Printf("DEBUG FindByStatus error: %v\n", err)
		return []TindakanApprovalItem{}, err
	}
	fmt.Printf("DEBUG FindByStatus: status=%s supervisor=%s rows=%d\n", status, supervisorName, len(items))

	if items == nil {
		items = []TindakanApprovalItem{}
	}

	return items, nil
}

func (r *TindakanRepoAdapter) UpdateStatus(ctx context.Context, id int, status string) error {
	query := `UPDATE tindakans SET status = $1::status_enum, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id)
	return err
}

// IsOwnedBySupervisor memeriksa apakah supervisor terkait adalah DPJP
// yang dipilih pada tindakan tersebut.
func (r *TindakanRepoAdapter) IsOwnedBySupervisor(ctx context.Context, id int, supervisorName string) (bool, error) {
	var exists bool
	err := r.DB.GetContext(ctx, &exists,
		`SELECT EXISTS(SELECT 1 FROM tindakans WHERE id = $1 AND COALESCE(supervisor_name, '') = $2)`,
		id, supervisorName)
	return exists, err
}

// KegiatanIlmiahRepoAdapter implements KegiatanIlmiahRepository
// Nama tabel di migration adalah "kegiatan_ilmiah" (tanpa 's')
type KegiatanIlmiahRepoAdapter struct {
	DB *sqlx.DB
}

func (r *KegiatanIlmiahRepoAdapter) FindByStatus(ctx context.Context, status, supervisorName string) ([]KegiatanApprovalItem, error) {
	// Subquery agar PostgreSQL mengembalikan tipe TEXT, bukan DATE/TIMESTAMP asli.
	query := `
		SELECT id, user_username, ppds_name, kategori, jenis_kegiatan, topik,
		       tanggal_mulai, lokasi_tipe, lokasi_detail, sebagai, status, created_at
		FROM (
			SELECT id,
			       COALESCE(user_username, '')::TEXT AS user_username,
			       COALESCE(ppds_name, '')::TEXT AS ppds_name,
			       COALESCE(kategori, '')::TEXT AS kategori,
			       COALESCE(jenis_kegiatan, '')::TEXT AS jenis_kegiatan,
			       COALESCE(topik, '')::TEXT AS topik,
			       COALESCE(tanggal_mulai::TEXT, '')::TEXT AS tanggal_mulai,
			       COALESCE(lokasi_tipe, '')::TEXT AS lokasi_tipe,
			       COALESCE(lokasi_detail, '')::TEXT AS lokasi_detail,
			       COALESCE(sebagai, '')::TEXT AS sebagai,
			       COALESCE(status::TEXT, '')::TEXT AS status,
			       COALESCE(created_at::TEXT, '')::TEXT AS created_at
			FROM kegiatan_ilmiah
			WHERE status = $1::status_kegiatan_enum
			  AND ($2 = '' OR COALESCE(pembimbing_1, '') = $2)
		) sub
		ORDER BY created_at DESC
	`

	var items []KegiatanApprovalItem
	if err := r.DB.SelectContext(ctx, &items, query, status, supervisorName); err != nil && err != sql.ErrNoRows {
		return []KegiatanApprovalItem{}, err
	}

	if items == nil {
		items = []KegiatanApprovalItem{}
	}

	return items, nil
}

func (r *KegiatanIlmiahRepoAdapter) UpdateStatus(ctx context.Context, id int, status string) error {
	query := `UPDATE kegiatan_ilmiah SET status = $1::status_kegiatan_enum, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id)
	return err
}

// IsOwnedBySupervisor memeriksa apakah supervisor terkait adalah
// pembimbing utama (pembimbing_1) pada kegiatan ilmiah tersebut.
func (r *KegiatanIlmiahRepoAdapter) IsOwnedBySupervisor(ctx context.Context, id int, supervisorName string) (bool, error) {
	var exists bool
	err := r.DB.GetContext(ctx, &exists,
		`SELECT EXISTS(SELECT 1 FROM kegiatan_ilmiah WHERE id = $1 AND COALESCE(pembimbing_1, '') = $2)`,
		id, supervisorName)
	return exists, err
}

// AktivitasKlinikRepoAdapter implements AktivitasKlinikRepository
// Kolom yang tersedia: tindakan (bukan nama_aktivitas)
type AktivitasKlinikRepoAdapter struct {
	DB *sqlx.DB
}

func (r *AktivitasKlinikRepoAdapter) FindByStatus(ctx context.Context, status, supervisorName string) ([]AktivitasKlinikApprovalItem, error) {
	// Subquery agar PostgreSQL mengembalikan tipe TEXT, bukan DATE/TIMESTAMP asli.
	query := `
		SELECT id, user_username, nama_aktivitas, tanggal, status, created_at
		FROM (
			SELECT id,
			       COALESCE(user_username, '')::TEXT AS user_username,
			       COALESCE(tindakan, '')::TEXT AS nama_aktivitas,
			       COALESCE(tanggal::TEXT, '')::TEXT AS tanggal,
			       COALESCE(status::TEXT, '')::TEXT AS status,
			       COALESCE(created_at::TEXT, '')::TEXT AS created_at
			FROM aktivitas_kliniks
			WHERE status = $1::enum_status_aktivitas
			  AND ($2 = '' OR COALESCE(supervisor, '') = $2)
		) sub
		ORDER BY created_at DESC
	`

	var items []AktivitasKlinikApprovalItem
	if err := r.DB.SelectContext(ctx, &items, query, status, supervisorName); err != nil && err != sql.ErrNoRows {
		return []AktivitasKlinikApprovalItem{}, err
	}

	if items == nil {
		items = []AktivitasKlinikApprovalItem{}
	}

	return items, nil
}

func (r *AktivitasKlinikRepoAdapter) UpdateStatus(ctx context.Context, id int, status string) error {
	query := `UPDATE aktivitas_kliniks SET status = $1::enum_status_aktivitas, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id)
	return err
}

// IsOwnedBySupervisor memeriksa apakah supervisor terkait tercatat
// pada aktivitas klinik tersebut.
func (r *AktivitasKlinikRepoAdapter) IsOwnedBySupervisor(ctx context.Context, id int, supervisorName string) (bool, error) {
	var exists bool
	err := r.DB.GetContext(ctx, &exists,
		`SELECT EXISTS(SELECT 1 FROM aktivitas_kliniks WHERE id = $1 AND COALESCE(supervisor, '') = $2)`,
		id, supervisorName)
	return exists, err
}

// PendidikanEvaluasiRepoAdapter implements PendidikanEvaluasiRepository
// Nama tabel di migration adalah "pendidikan_evaluasi" (tanpa 's')
// Kolom yang tersedia: kategori (bukan jenis_evaluasi)
type PendidikanEvaluasiRepoAdapter struct {
	DB *sqlx.DB
}

func (r *PendidikanEvaluasiRepoAdapter) FindByStatus(ctx context.Context, status, supervisorName string) ([]PendidikanEvaluasiApprovalItem, error) {
	// Subquery agar PostgreSQL mengembalikan tipe TEXT, bukan TIMESTAMP asli.
	query := `
		SELECT id, user_username, jenis_evaluasi, tanggal, status, created_at
		FROM (
			SELECT id,
			       COALESCE(user_username, '')::TEXT AS user_username,
			       COALESCE(kategori, '')::TEXT AS jenis_evaluasi,
			       COALESCE(tanggal, '')::TEXT AS tanggal,
			       COALESCE(status::TEXT, '')::TEXT AS status,
			       COALESCE(created_at::TEXT, '')::TEXT AS created_at
			FROM pendidikan_evaluasi
			WHERE status = $1
			  AND ($2 = '' OR COALESCE(supervisor, '') = $2 OR COALESCE(evaluator, '') = $2 OR COALESCE(pembimbing, '') = $2)
		) sub
		ORDER BY created_at DESC
	`

	var items []PendidikanEvaluasiApprovalItem
	if err := r.DB.SelectContext(ctx, &items, query, status, supervisorName); err != nil && err != sql.ErrNoRows {
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

// IsOwnedBySupervisor memeriksa apakah supervisor terkait tercatat sebagai
// supervisor, evaluator, atau pembimbing pada evaluasi pendidikan tersebut.
func (r *PendidikanEvaluasiRepoAdapter) IsOwnedBySupervisor(ctx context.Context, id int, supervisorName string) (bool, error) {
	var exists bool
	err := r.DB.GetContext(ctx, &exists,
		`SELECT EXISTS(
			SELECT 1 FROM pendidikan_evaluasi
			WHERE id = $1
			  AND (COALESCE(supervisor, '') = $2 OR COALESCE(evaluator, '') = $2 OR COALESCE(pembimbing, '') = $2)
		)`,
		id, supervisorName)
	return exists, err
}
