package jadwal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Service interface {
	Create(ctx context.Context, req CreateJadwalRequest, username string) (*EventResponse, error)
	GetEvents(ctx context.Context, startStr, endStr, typeStr string) ([]EventResponse, error)
	Update(ctx context.Context, id int, req UpdateJadwalRequest) (*EventResponse, error)
	UpdateDates(ctx context.Context, id int, req UpdateDatesRequest) (*EventResponse, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func getEventClassName(typeStr string) string {
	switch typeStr {
	case "bimbingan":
		return "fc-event-warning"
	case "rotasi":
		return "fc-event-success"
	case "ujian":
		return "fc-event-danger"
	default:
		return "fc-event-primary"
	}
}

func parseTime(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}
	// Coba format ISO8601 / RFC3339 terlebih dahulu
	t, err := time.Parse(time.RFC3339, timeStr)
	if err == nil {
		return t, nil
	}
	// Fallback format 'YYYY-MM-DD HH:mm:ss' atau 'YYYY-MM-DD'
	t, err = time.Parse("2006-01-02 15:04:05", timeStr)
	if err == nil {
		return t, nil
	}
	return time.Parse("2006-01-02", timeStr)
}

func (s *service) Create(ctx context.Context, req CreateJadwalRequest, username string) (*EventResponse, error) {
	startTime, err := parseTime(req.Start)
	if err != nil {
		return nil, errors.New("Format waktu 'start' tidak valid")
	}
	endTime, err := parseTime(req.End)
	if err != nil {
		return nil, errors.New("Format waktu 'end' tidak valid")
	}

	j := &Jadwal{
		Title:        req.Title,
		Description:  sql.NullString{String: req.Description, Valid: req.Description != ""},
		Location:     sql.NullString{String: req.Location, Valid: req.Location != ""},
		StartTime:    startTime,
		EndTime:      endTime,
		AllDay:       req.AllDay,
		Type:         req.Type,
		UserUsername: sql.NullString{String: username, Valid: username != ""},
	}

	if err := s.repo.Create(ctx, j); err != nil {
		return nil, err
	}

	return s.toEventResponse(j), nil
}

func (s *service) GetEvents(ctx context.Context, startStr, endStr, typeStr string) ([]EventResponse, error) {
	var startFilter, endFilter time.Time
	if startStr != "" {
		startFilter, _ = parseTime(startStr)
	}
	if endStr != "" {
		endFilter, _ = parseTime(endStr)
	}

	list, err := s.repo.FindAll(ctx, startFilter, endFilter, typeStr)
	if err != nil {
		return nil, err
	}

	var events []EventResponse
	for _, j := range list {
		events = append(events, *s.toEventResponse(&j))
	}
	return events, nil
}

func (s *service) Update(ctx context.Context, id int, req UpdateJadwalRequest) (*EventResponse, error) {
	j, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("Jadwal tidak ditemukan")
	}

	startTime, err := parseTime(req.Start)
	if err != nil {
		return nil, errors.New("Format waktu 'start' tidak valid")
	}
	endTime, err := parseTime(req.End)
	if err != nil {
		return nil, errors.New("Format waktu 'end' tidak valid")
	}

	j.Title = req.Title
	j.Description = sql.NullString{String: req.Description, Valid: req.Description != ""}
	j.Location = sql.NullString{String: req.Location, Valid: req.Location != ""}
	j.StartTime = startTime
	j.EndTime = endTime
	j.AllDay = req.AllDay
	j.Type = req.Type

	if err := s.repo.Update(ctx, j); err != nil {
		return nil, err
	}

	return s.toEventResponse(j), nil
}

func (s *service) UpdateDates(ctx context.Context, id int, req UpdateDatesRequest) (*EventResponse, error) {
	j, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("Jadwal tidak ditemukan")
	}

	startTime, err := parseTime(req.Start)
	if err != nil {
		return nil, errors.New("Format waktu 'start' tidak valid")
	}
	endTime, err := parseTime(req.End)
	if err != nil {
		return nil, errors.New("Format waktu 'end' tidak valid")
	}

	j.StartTime = startTime
	j.EndTime = endTime
	j.AllDay = req.AllDay

	if err := s.repo.UpdateDates(ctx, id, startTime, endTime, req.AllDay); err != nil {
		return nil, err
	}

	return s.toEventResponse(j), nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("Jadwal tidak ditemukan")
	}
	return s.repo.Delete(ctx, id)
}

func (s *service) toEventResponse(j *Jadwal) *EventResponse {
	return &EventResponse{
		ID:          fmt.Sprintf("%d", j.ID),
		Title:       j.Title,
		Description: j.Description.String,
		Location:    j.Location.String,
		Start:       j.StartTime.Format(time.RFC3339),
		End:         j.EndTime.Format(time.RFC3339),
		AllDay:      j.AllDay,
		ClassName:   getEventClassName(j.Type),
		Type:        j.Type,
		UserID:      0,
	}
}
