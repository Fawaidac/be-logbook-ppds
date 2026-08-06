package jadwal

import (
	"database/sql"
	"time"
)

type Jadwal struct {
	ID          int            `db:"id" json:"id"`
	Title       string         `db:"title" json:"title"`
	Description sql.NullString `db:"description" json:"description"`
	Location    sql.NullString `db:"location" json:"location"`
	StartTime   time.Time      `db:"start_time" json:"start_time"`
	EndTime     time.Time      `db:"end_time" json:"end_time"`
	AllDay      bool           `db:"all_day" json:"all_day"`
	Type        string         `db:"type" json:"type"`
	UserID      sql.NullInt64  `db:"user_id" json:"user_id"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
}
