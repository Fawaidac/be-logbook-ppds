package tindakan

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Create(ctx context.Context, t *Tindakan) error
	FindAll(ctx context.Context, userID int) ([]Tindakan, error)
	FindByID(ctx context.Context, id int) (*Tindakan, error)
	Update(ctx context.Context, t *Tindakan) error
	UpdateStatus(ctx context.Context, id int, status string) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, t *Tindakan) error {
	query := `
		INSERT INTO tindakans (
			user_username, mr_number, visit_date, patient_name, gender, birth_date,
			division, diagnosis_label, procedure_code, plan_procedure,
			activity, procedure_date, room, role, kemandirian, clinical_note,
			supervisor_name, status, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16,
			$17, $18, NOW()
		)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRowContext(ctx, query,
		t.UserUsername, t.MRNumber, t.VisitDate, t.PatientName, t.Gender, t.BirthDate,
		t.Division, t.DiagnosisLabel, t.ProcedureCode, t.PlanProcedure,
		t.Activity, t.ProcedureDate, t.Room, t.Role, t.Kemandirian, t.ClinicalNote,
		t.SupervisorName, t.Status,
	).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
}

func (r *repository) FindAll(ctx context.Context, userID int) ([]Tindakan, error) {
	var list []Tindakan
	var err error

	if userID > 0 {
		query := `SELECT * FROM tindakans WHERE user_username = $1 ORDER BY id DESC`
		err = r.db.SelectContext(ctx, &list, query, userID)
	} else {
		query := `SELECT * FROM tindakans ORDER BY id DESC`
		err = r.db.SelectContext(ctx, &list, query)
	}

	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *repository) FindByID(ctx context.Context, id int) (*Tindakan, error) {
	var t Tindakan
	query := `SELECT * FROM tindakans WHERE id = $1`
	err := r.db.GetContext(ctx, &t, query, id)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *repository) Update(ctx context.Context, t *Tindakan) error {
	query := `
		UPDATE tindakans SET
			mr_number = $1, patient_name = $2, plan_procedure = $3,
			diagnosis_label = $4, room = $5, clinical_note = $6, updated_at = NOW()
		WHERE id = $7`

	_, err := r.db.ExecContext(ctx, query,
		t.MRNumber, t.PatientName, t.PlanProcedure,
		t.DiagnosisLabel, t.Room, t.ClinicalNote, t.ID,
	)
	return err
}

func (r *repository) UpdateStatus(ctx context.Context, id int, status string) error {
	query := `UPDATE tindakans SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM tindakans WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}