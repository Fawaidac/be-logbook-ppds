package tindakan

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"time"
)

type Service interface {
	Create(ctx context.Context, req CreateTindakanRequest, userID int) (*TindakanResponse, error)
	GetSummary(ctx context.Context, userID int) (*SummaryResponse, error)
	GetByID(ctx context.Context, id int) (*TindakanResponse, error)
	Update(ctx context.Context, id int, req UpdateTindakanRequest) (*TindakanResponse, error)
	Send(ctx context.Context, id int) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func parseDate(dateStr string) sql.NullTime {
	if dateStr == "" {
		return sql.NullTime{Valid: false}
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		t, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			return sql.NullTime{Valid: false}
		}
	}
	return sql.NullTime{Time: t, Valid: true}
}

func formatDate(nt sql.NullTime) string {
	if !nt.Valid {
		return ""
	}
	return nt.Time.Format("2006-01-02")
}

func defaultString(val, fallback string) string {
	if val != "" {
		return val
	}
	return fallback
}

func (s *service) Create(ctx context.Context, req CreateTindakanRequest, userID int) (*TindakanResponse, error) {
	procDate := parseDate(req.ProcedureDate)
	if !procDate.Valid {
		procDate = sql.NullTime{Time: time.Now(), Valid: true}
	}

	t := &Tindakan{
		MRNumber:       req.MRNumber,
		VisitDate:      parseDate(req.VisitDate),
		PatientName:    req.PatientName,
		Gender:         sql.NullString{String: defaultString(req.Gender, "L"), Valid: true},
		BirthDate:      parseDate(req.BirthDate),
		Division:       sql.NullString{String: defaultString(req.Division, "Ortopedi & Traumatologi"), Valid: true},
		DiagnosisLabel: sql.NullString{String: defaultString(req.DiagnosisLabel, "-"), Valid: true},
		ProcedureCode:  sql.NullString{String: defaultString(req.ProcedureCode, "PROC-001"), Valid: true},
		PlanProcedure:  req.PlanProcedure,
		Activity:       sql.NullString{String: defaultString(req.Activity, "Operasi Elektif"), Valid: true},
		ProcedureDate:  procDate,
		Room:           sql.NullString{String: defaultString(req.RoomLabel, "OK Sentral RSU"), Valid: true},
		Role:           sql.NullString{String: defaultString(req.RoleLabel, "Operator Utama"), Valid: true},
		Kemandirian:    defaultString(req.Kemandirian, "dibimbing"),
		ClinicalNote:   sql.NullString{String: defaultString(req.ClinicalNote, "-"), Valid: true},
		SupervisorName: sql.NullString{String: defaultString(req.SupervisorName, "dr. Andi Wijaya, Sp.OT"), Valid: true},
		Status:         "draft",
	}

	if err := s.repo.Create(ctx, t); err != nil {
		return nil, err
	}

	return s.toResponse(t), nil
}

func (s *service) GetSummary(ctx context.Context, userID int) (*SummaryResponse, error) {
	list, err := s.repo.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	var entries []TindakanResponse
	totalCount := len(list)
	mandiriCount := 0
	dibimbingCount := 0
	observasiCount := 0
	disetujuiCount := 0

	for _, item := range list {
		if item.Kemandirian == "mandiri" {
			mandiriCount++
		} else if item.Kemandirian == "dibimbing" {
			dibimbingCount++
		} else if item.Kemandirian == "observasi" {
			observasiCount++
		}

		if item.Status == "disetujui" {
			disetujuiCount++
		}

		entries = append(entries, *s.toResponse(&item))
	}

	verifikasiPercent := 0
	if totalCount > 0 {
		verifikasiPercent = int(math.Round((float64(disetujuiCount) / float64(totalCount)) * 100))
	}

	return &SummaryResponse{
		TotalCount:        totalCount,
		MandiriCount:      mandiriCount,
		DibimbingCount:    dibimbingCount,
		ObservasiCount:    observasiCount,
		DisetujuiCount:    disetujuiCount,
		VerifikasiPercent: verifikasiPercent,
		Entries:           entries,
	}, nil
}

func (s *service) GetByID(ctx context.Context, id int) (*TindakanResponse, error) {
	t, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("Tindakan tidak ditemukan")
	}
	return s.toResponse(t), nil
}

func (s *service) Update(ctx context.Context, id int, req UpdateTindakanRequest) (*TindakanResponse, error) {
	t, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("Tindakan tidak ditemukan")
	}

	t.MRNumber = req.MRNumber
	t.PatientName = req.PatientName
	t.PlanProcedure = req.PlanProcedure

	if req.DiagnosisLabel != "" {
		t.DiagnosisLabel = sql.NullString{String: req.DiagnosisLabel, Valid: true}
	}
	if req.RoomLabel != "" {
		t.Room = sql.NullString{String: req.RoomLabel, Valid: true}
	}
	if req.ClinicalNote != "" {
		t.ClinicalNote = sql.NullString{String: req.ClinicalNote, Valid: true}
	}

	if err := s.repo.Update(ctx, t); err != nil {
		return nil, err
	}

	return s.toResponse(t), nil
}

func (s *service) Send(ctx context.Context, id int) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("Tindakan tidak ditemukan")
	}
	return s.repo.UpdateStatus(ctx, id, "menunggu")
}

func (s *service) Delete(ctx context.Context, id int) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("Tindakan tidak ditemukan")
	}
	return s.repo.Delete(ctx, id)
}

func (s *service) toResponse(t *Tindakan) *TindakanResponse {
	return &TindakanResponse{
		ID:             t.ID,
		UserUsername:   t.UserUsername.String,
		MRNumber:       t.MRNumber,
		VisitDate:      formatDate(t.VisitDate),
		PatientName:    t.PatientName,
		Gender:         t.Gender.String,
		BirthDate:      formatDate(t.BirthDate),
		Division:       t.Division.String,
		DiagnosisLabel: t.DiagnosisLabel.String,
		ProcedureCode:  t.ProcedureCode.String,
		PlanProcedure:  t.PlanProcedure,
		Activity:       t.Activity.String,
		ProcedureDate:  formatDate(t.ProcedureDate),
		Room:           t.Room.String,
		Role:           t.Role.String,
		Kemandirian:    t.Kemandirian,
		ClinicalNote:   t.ClinicalNote.String,
		SupervisorName: t.SupervisorName.String,
		Status:         t.Status,
		Feedback:       t.Feedback.String,
		CreatedAt:      t.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      t.UpdatedAt.Format(time.RFC3339),
	}
}