package tindakan

import (
	"database/sql"
	"time"
)

type Tindakan struct {
	ID             int            `db:"id" json:"id"`
	UserUsername   sql.NullString `db:"user_username" json:"user_username"`
	MRNumber       string         `db:"mr_number" json:"mr_number"`
	VisitDate      sql.NullTime   `db:"visit_date" json:"visit_date"`
	PatientName    string         `db:"patient_name" json:"patient_name"`
	Gender         sql.NullString `db:"gender" json:"gender"`
	BirthDate      sql.NullTime   `db:"birth_date" json:"birth_date"`
	Division       sql.NullString `db:"division" json:"division"`
	DiagnosisLabel sql.NullString `db:"diagnosis_label" json:"diagnosis_label"`
	ProcedureCode  sql.NullString `db:"procedure_code" json:"procedure_code"`
	PlanProcedure  string         `db:"plan_procedure" json:"plan_procedure"`
	Activity       sql.NullString `db:"activity" json:"activity"`
	ProcedureDate  sql.NullTime   `db:"procedure_date" json:"procedure_date"`
	Room           sql.NullString `db:"room" json:"room"`
	Role           sql.NullString `db:"role" json:"role"`
	Kemandirian    string         `db:"kemandirian" json:"kemandirian"`
	ClinicalNote   sql.NullString `db:"clinical_note" json:"clinical_note"`
	SupervisorName sql.NullString `db:"supervisor_name" json:"supervisor_name"`
	Status         string         `db:"status" json:"status"`
	Feedback       sql.NullString `db:"feedback" json:"feedback"`
	CreatedAt      time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at" json:"updated_at"`
}