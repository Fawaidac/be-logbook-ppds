package jadwal

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Create(ctx context.Context, j *Jadwal) error
	FindAll(ctx context.Context, startFilter, endFilter time.Time, typeFilter string) ([]Jadwal, error)
	FindByID(ctx context.Context, id int) (*Jadwal, error)
	Update(ctx context.Context, j *Jadwal) error
	UpdateDates(ctx context.Context, id int, startTime, endTime time.Time, allDay bool) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, j *Jadwal) error {
	query := `
		INSERT INTO jadwals (title, description, location, start_time, end_time, all_day, type, user_id, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		j.Title, j.Description, j.Location, j.StartTime, j.EndTime, j.AllDay, j.Type, j.UserID,
	).Scan(&j.ID, &j.CreatedAt, &j.UpdatedAt)
}

func (r *repository) FindAll(ctx context.Context, startFilter, endFilter time.Time, typeFilter string) ([]Jadwal, error) {
	var list []Jadwal
	query := `SELECT id, title, description, location, start_time, end_time, all_day, type, user_id, created_at, updated_at FROM jadwals WHERE 1=1`
	args := []interface{}{}
	argIdx := 1

	if !startFilter.IsZero() {
		query += fmt.Sprintf(" AND start_time >= $%d", argIdx)
		args = append(args, startFilter)
		argIdx++
	}
	if !endFilter.IsZero() {
		query += fmt.Sprintf(" AND end_time <= $%d", argIdx)
		args = append(args, endFilter)
		argIdx++
	}
	if typeFilter != "" && typeFilter != "all" {
		query += fmt.Sprintf(" AND type = $%d", argIdx)
		args = append(args, typeFilter)
		argIdx++
	}

	query += ` ORDER BY start_time ASC`

	err := r.db.SelectContext(ctx, &list, query, args...)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *repository) FindByID(ctx context.Context, id int) (*Jadwal, error) {
	var j Jadwal
	query := `SELECT id, title, description, location, start_time, end_time, all_day, type, user_id, created_at, updated_at FROM jadwals WHERE id = $1`
	err := r.db.GetContext(ctx, &j, query, id)
	if err != nil {
		return nil, err
	}
	return &j, nil
}

func (r *repository) Update(ctx context.Context, j *Jadwal) error {
	query := `
		UPDATE jadwals
		SET title = $1, description = $2, location = $3, start_time = $4, end_time = $5, all_day = $6, type = $7, updated_at = NOW()
		WHERE id = $8`
	_, err := r.db.ExecContext(ctx, query,
		j.Title, j.Description, j.Location, j.StartTime, j.EndTime, j.AllDay, j.Type, j.ID,
	)
	return err
}

func (r *repository) UpdateDates(ctx context.Context, id int, startTime, endTime time.Time, allDay bool) error {
	query := `
		UPDATE jadwals
		SET start_time = $1, end_time = $2, all_day = $3, updated_at = NOW()
		WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, startTime, endTime, allDay, id)
	return err
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM jadwals WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
