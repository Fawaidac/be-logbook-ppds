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
	query := `
		SELECT id, user_username, mr_number, visit_date, patient_name, gender, birth_date,
		       division, diagnosis_label, procedure_code, plan_procedure, activity,
		       procedure_date, room, role_label, kemandirian, clinical_note,
		       supervisor_name, feedback, status, created_at
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
			       COALESCE(feedback, '')::TEXT AS feedback,
			       COALESCE(status::TEXT, '')::TEXT AS status,
			       COALESCE(created_at::TEXT, '')::TEXT AS created_at
			FROM tindakans
			WHERE status = $1::status_enum
			  AND (
			    $2 = ''
			    OR LOWER(TRIM(COALESCE(supervisor_name, ''))) = LOWER(TRIM($2))
			    OR (LENGTH(TRIM($2)) >= 3 AND LOWER(TRIM(COALESCE(supervisor_name, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%')
			    OR (LENGTH(TRIM(COALESCE(supervisor_name, ''))) >= 3 AND LOWER(TRIM($2)) ILIKE '%' || LOWER(TRIM(COALESCE(supervisor_name, ''))) || '%')
			  )
		) sub
		ORDER BY created_at DESC
	`

	var items []TindakanApprovalItem
	if err := r.DB.SelectContext(ctx, &items, query, status, supervisorName); err != nil && err != sql.ErrNoRows {
		fmt.Printf("DEBUG Tindakan FindByStatus error: %v\n", err)
		return []TindakanApprovalItem{}, err
	}

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

func (r *TindakanRepoAdapter) UpdateStatusWithNote(ctx context.Context, id int, status string, catatan string) error {
	query := `UPDATE tindakans SET status = $1::status_enum, feedback = $3, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id, catatan)
	return err
}

func (r *TindakanRepoAdapter) IsOwnedBySupervisor(ctx context.Context, id int, supervisorName string) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS(
			SELECT 1 FROM tindakans
			WHERE id = $1 AND (
			    $2 = ''
			    OR LOWER(TRIM(COALESCE(supervisor_name, ''))) = LOWER(TRIM($2))
			    OR (LENGTH(TRIM($2)) >= 3 AND LOWER(TRIM(COALESCE(supervisor_name, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%')
			    OR (LENGTH(TRIM(COALESCE(supervisor_name, ''))) >= 3 AND LOWER(TRIM($2)) ILIKE '%' || LOWER(TRIM(COALESCE(supervisor_name, ''))) || '%')
			)
		)`
	err := r.DB.GetContext(ctx, &exists, query, id, supervisorName)
	return exists, err
}

// KegiatanIlmiahRepoAdapter implements KegiatanIlmiahRepository
type KegiatanIlmiahRepoAdapter struct {
	DB *sqlx.DB
}

func (r *KegiatanIlmiahRepoAdapter) FindByStatus(ctx context.Context, status, supervisorName string) ([]KegiatanApprovalItem, error) {
	query := `
		SELECT id, user_username, ppds_name, kategori, jenis_kegiatan, topik,
		       tanggal_mulai, lokasi_tipe, lokasi_detail, sebagai, pembimbing_1, catatan_pembimbing, status, created_at
		FROM (
			SELECT id,
			       COALESCE(user_username, '')::TEXT AS user_username,
			       COALESCE(ppds_name, '')::TEXT AS ppds_name,
			       COALESCE(kategori::TEXT, '')::TEXT AS kategori,
			       COALESCE(jenis_kegiatan, '')::TEXT AS jenis_kegiatan,
			       COALESCE(topik, '')::TEXT AS topik,
			       COALESCE(tanggal_mulai::TEXT, '')::TEXT AS tanggal_mulai,
			       COALESCE(lokasi_tipe::TEXT, '')::TEXT AS lokasi_tipe,
			       COALESCE(lokasi_detail, '')::TEXT AS lokasi_detail,
			       COALESCE(sebagai, '')::TEXT AS sebagai,
			       COALESCE(pembimbing_1, '')::TEXT AS pembimbing_1,
			       COALESCE(catatan_pembimbing, '')::TEXT AS catatan_pembimbing,
			       COALESCE(status::TEXT, '')::TEXT AS status,
			       COALESCE(created_at::TEXT, '')::TEXT AS created_at
			FROM kegiatan_ilmiah
			WHERE status = $1::status_kegiatan_enum
			  AND (
			    $2 = ''
			    OR LOWER(TRIM(COALESCE(pembimbing_1, ''))) = LOWER(TRIM($2))
			    OR LOWER(TRIM(COALESCE(pembimbing_2, ''))) = LOWER(TRIM($2))
			    OR LOWER(TRIM(COALESCE(pembimbing_3, ''))) = LOWER(TRIM($2))
			    OR LOWER(TRIM(COALESCE(pembimbing_4, ''))) = LOWER(TRIM($2))
			    OR LOWER(TRIM(COALESCE(pembimbing_5, ''))) = LOWER(TRIM($2))
			    OR (LENGTH(TRIM($2)) >= 3 AND (
			         LOWER(TRIM(COALESCE(pembimbing_1, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%'
			      OR LOWER(TRIM(COALESCE(pembimbing_2, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%'
			      OR LOWER(TRIM(COALESCE(pembimbing_3, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%'
			      OR LOWER(TRIM(COALESCE(pembimbing_4, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%'
			      OR LOWER(TRIM(COALESCE(pembimbing_5, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%'
			    ))
			    OR (LENGTH(TRIM(COALESCE(pembimbing_1, ''))) >= 3 AND LOWER(TRIM($2)) ILIKE '%' || LOWER(TRIM(COALESCE(pembimbing_1, ''))) || '%')
			  )
		) sub
		ORDER BY created_at DESC
	`

	var items []KegiatanApprovalItem
	if err := r.DB.SelectContext(ctx, &items, query, status, supervisorName); err != nil && err != sql.ErrNoRows {
		fmt.Printf("DEBUG KegiatanIlmiah FindByStatus error: %v\n", err)
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

func (r *KegiatanIlmiahRepoAdapter) UpdateStatusWithNote(ctx context.Context, id int, status string, catatan string) error {
	query := `UPDATE kegiatan_ilmiah SET status = $1::status_kegiatan_enum, catatan_pembimbing = $3, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id, catatan)
	return err
}

func (r *KegiatanIlmiahRepoAdapter) IsOwnedBySupervisor(ctx context.Context, id int, supervisorName string) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS(
			SELECT 1 FROM kegiatan_ilmiah
			WHERE id = $1 AND (
			    $2 = ''
			    OR LOWER(TRIM(COALESCE(pembimbing_1, ''))) = LOWER(TRIM($2))
			    OR LOWER(TRIM(COALESCE(pembimbing_2, ''))) = LOWER(TRIM($2))
			    OR LOWER(TRIM(COALESCE(pembimbing_3, ''))) = LOWER(TRIM($2))
			    OR LOWER(TRIM(COALESCE(pembimbing_4, ''))) = LOWER(TRIM($2))
			    OR LOWER(TRIM(COALESCE(pembimbing_5, ''))) = LOWER(TRIM($2))
			    OR (LENGTH(TRIM($2)) >= 3 AND (
			         LOWER(TRIM(COALESCE(pembimbing_1, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%'
			      OR LOWER(TRIM(COALESCE(pembimbing_2, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%'
			      OR LOWER(TRIM(COALESCE(pembimbing_3, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%'
			      OR LOWER(TRIM(COALESCE(pembimbing_4, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%'
			      OR LOWER(TRIM(COALESCE(pembimbing_5, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%'
			    ))
			    OR (LENGTH(TRIM(COALESCE(pembimbing_1, ''))) >= 3 AND LOWER(TRIM($2)) ILIKE '%' || LOWER(TRIM(COALESCE(pembimbing_1, ''))) || '%')
			)
		)`
	err := r.DB.GetContext(ctx, &exists, query, id, supervisorName)
	return exists, err
}

// PendidikanEvaluasiRepoAdapter implements PendidikanEvaluasiRepository
type PendidikanEvaluasiRepoAdapter struct {
	DB *sqlx.DB
}

func (r *PendidikanEvaluasiRepoAdapter) FindByStatus(ctx context.Context, status, supervisorName string) ([]PendidikanEvaluasiApprovalItem, error) {
	query := `
		SELECT id, user_username, jenis_evaluasi, tanggal, revisi_catatan, status, created_at
		FROM (
			SELECT id,
			       COALESCE(user_username, '')::TEXT AS user_username,
			       COALESCE(NULLIF(nama, ''), NULLIF(jenis, ''), NULLIF(kategori, ''), '')::TEXT AS jenis_evaluasi,
			       COALESCE(tanggal::TEXT, '')::TEXT AS tanggal,
			       COALESCE(revisi_catatan, '')::TEXT AS revisi_catatan,
			       COALESCE(status::TEXT, '')::TEXT AS status,
			       COALESCE(created_at::TEXT, '')::TEXT AS created_at
			FROM pendidikan_evaluasi
			WHERE status = $1
			  AND (
			    $2 = ''
			    OR LOWER(TRIM(COALESCE(supervisor, evaluator, pembimbing, ''))) = LOWER(TRIM($2))
			    OR (LENGTH(TRIM($2)) >= 3 AND LOWER(TRIM(COALESCE(supervisor, evaluator, pembimbing, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%')
			    OR (LENGTH(TRIM(COALESCE(supervisor, evaluator, pembimbing, ''))) >= 3 AND LOWER(TRIM($2)) ILIKE '%' || LOWER(TRIM(COALESCE(supervisor, evaluator, pembimbing, ''))) || '%')
			  )
		) sub
		ORDER BY created_at DESC
	`

	var items []PendidikanEvaluasiApprovalItem
	if err := r.DB.SelectContext(ctx, &items, query, status, supervisorName); err != nil && err != sql.ErrNoRows {
		fmt.Printf("DEBUG PendidikanEvaluasi FindByStatus error: %v\n", err)
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

func (r *PendidikanEvaluasiRepoAdapter) UpdateStatusWithNote(ctx context.Context, id int, status string, catatan string) error {
	query := `UPDATE pendidikan_evaluasi SET status = $1, revisi_catatan = $3, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id, catatan)
	return err
}

func (r *PendidikanEvaluasiRepoAdapter) IsOwnedBySupervisor(ctx context.Context, id int, supervisorName string) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS(
			SELECT 1 FROM pendidikan_evaluasi
			WHERE id = $1 AND (
			    $2 = ''
			    OR LOWER(TRIM(COALESCE(supervisor, evaluator, pembimbing, ''))) = LOWER(TRIM($2))
			    OR (LENGTH(TRIM($2)) >= 3 AND LOWER(TRIM(COALESCE(supervisor, evaluator, pembimbing, ''))) ILIKE '%' || LOWER(TRIM($2)) || '%')
			    OR (LENGTH(TRIM(COALESCE(supervisor, evaluator, pembimbing, ''))) >= 3 AND LOWER(TRIM($2)) ILIKE '%' || LOWER(TRIM(COALESCE(supervisor, evaluator, pembimbing, ''))) || '%')
			)
		)`
	err := r.DB.GetContext(ctx, &exists, query, id, supervisorName)
	return exists, err
}

