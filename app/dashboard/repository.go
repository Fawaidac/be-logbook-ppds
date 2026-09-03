package dashboard

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetTotalTindakan(ctx context.Context, username, name string) (int, error)
	GetMenungguValidasi(ctx context.Context, username, name string) (int, error)
	GetPerluRevisi(ctx context.Context, username, name string) (int, error)
	GetUpcomingJadwals(ctx context.Context, username, name string) ([]UpcomingJadwalItem, error)
	GetRecentEntries(ctx context.Context, username, name string) ([]RecentTindakanItem, error)
	CountTindakanByMonth(ctx context.Context, year, month int, kemandirian string, username, name string) (int, error)
	CountKegiatanByMonth(ctx context.Context, year, month int, username, name string) (int, error)

	GetRekapTindakan(ctx context.Context, periode, stase string, username, name string) ([]RekapTindakanItem, error)
	GetKegiatanCountByKategori(ctx context.Context, kategori string, username, name string) (int, error)
	GetDPJPStats(ctx context.Context, username, name string) ([]DPJPStatItem, error)
	GetTotalBimbingan(ctx context.Context, username, name string) (int, error)
	GetMenungguBimbinganCount(ctx context.Context, username, name string) (int, error)
	GetDisetujuiBimbinganCount(ctx context.Context, username, name string) (int, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

const userFilterCond = `($1 = '' OR user_username = $1 OR ($2 != '' AND user_username = $2))`

func (r *repository) GetTotalTindakan(ctx context.Context, username, name string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM tindakans WHERE ` + userFilterCond
	err := r.db.GetContext(ctx, &count, query, username, name)
	return count, err
}

func (r *repository) GetMenungguValidasi(ctx context.Context, username, name string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM tindakans WHERE status = 'menunggu' AND ` + userFilterCond
	err := r.db.GetContext(ctx, &count, query, username, name)
	return count, err
}

func (r *repository) GetPerluRevisi(ctx context.Context, username, name string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM tindakans WHERE status = 'ditolak' AND ` + userFilterCond
	err := r.db.GetContext(ctx, &count, query, username, name)
	return count, err
}

func (r *repository) GetUpcomingJadwals(ctx context.Context, username, name string) ([]UpcomingJadwalItem, error) {
	var items []UpcomingJadwalItem
	query := `
		SELECT id, title, COALESCE(location, '') AS location, start_time, end_time, all_day, COALESCE(type, 'jaga') AS type
		FROM jadwals
		WHERE (` + userFilterCond + ` OR user_username IS NULL OR user_username = '')
		ORDER BY start_time ASC
		LIMIT 3
	`
	err := r.db.SelectContext(ctx, &items, query, username, name)
	if items == nil {
		items = []UpcomingJadwalItem{}
	}
	return items, err
}

func (r *repository) GetRecentEntries(ctx context.Context, username, name string) ([]RecentTindakanItem, error) {
	var items []RecentTindakanItem
	query := `
		SELECT id,
		       COALESCE(procedure_date, created_at) AS procedure_date,
		       COALESCE(diagnosis_label, '-') AS diagnosis_label,
		       COALESCE(plan_procedure, '-') AS plan_procedure,
		       COALESCE(kemandirian, 'dibimbing') AS kemandirian,
		       COALESCE(supervisor_name, '-') AS supervisor_name,
		       COALESCE(supervisor_name, '-') AS dpjp_name,
		       COALESCE(status, 'menunggu') AS status
		FROM tindakans
		WHERE ` + userFilterCond + `
		ORDER BY created_at DESC
		LIMIT 5
	`
	err := r.db.SelectContext(ctx, &items, query, username, name)
	if items == nil {
		items = []RecentTindakanItem{}
	}
	return items, err
}

func (r *repository) CountTindakanByMonth(ctx context.Context, year, month int, kemandirian string, username, name string) (int, error) {
	var count int
	query := `
		SELECT COUNT(*)
		FROM tindakans
		WHERE status = 'disetujui'
		  AND kemandirian = $1
		  AND EXTRACT(YEAR FROM COALESCE(procedure_date, created_at)) = $2
		  AND EXTRACT(MONTH FROM COALESCE(procedure_date, created_at)) = $3
		  AND ($4 = '' OR user_username = $4 OR ($5 != '' AND user_username = $5))
	`
	err := r.db.GetContext(ctx, &count, query, kemandirian, year, month, username, name)
	return count, err
}

func (r *repository) CountKegiatanByMonth(ctx context.Context, year, month int, username, name string) (int, error) {
	var count int
	query := `
		SELECT COUNT(*)
		FROM kegiatan_ilmiah
		WHERE status = 'disetujui'
		  AND EXTRACT(YEAR FROM COALESCE(tanggal_mulai, created_at)) = $1
		  AND EXTRACT(MONTH FROM COALESCE(tanggal_mulai, created_at)) = $2
		  AND ($3 = '' OR user_username = $3 OR ($4 != '' AND user_username = $4))
	`
	err := r.db.GetContext(ctx, &count, query, year, month, username, name)
	return count, err
}

func (r *repository) GetRekapTindakan(ctx context.Context, periode, stase string, username, name string) ([]RekapTindakanItem, error) {
	var items []RekapTindakanItem
	query := `
		SELECT COALESCE(plan_procedure, '-') AS plan_procedure,
		       COUNT(CASE WHEN kemandirian = 'mandiri' THEN 1 END) AS mandiri_count,
		       COUNT(CASE WHEN kemandirian = 'dibimbing' THEN 1 END) AS dibimbing_count,
		       COUNT(CASE WHEN kemandirian = 'observasi' THEN 1 END) AS observasi_count,
		       COUNT(*) AS total_count
		FROM tindakans
		WHERE ` + userFilterCond + `
		GROUP BY plan_procedure
	`
	err := r.db.SelectContext(ctx, &items, query, username, name)
	if items == nil {
		items = []RekapTindakanItem{}
	}
	return items, err
}

func (r *repository) GetKegiatanCountByKategori(ctx context.Context, kategori string, username, name string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM kegiatan_ilmiah WHERE kategori::text = $1 AND ($2 = '' OR user_username = $2 OR ($3 != '' AND user_username = $3))`
	err := r.db.GetContext(ctx, &count, query, kategori, username, name)
	return count, err
}

func (r *repository) GetDPJPStats(ctx context.Context, username, name string) ([]DPJPStatItem, error) {
	var items []DPJPStatItem
	query := `
		SELECT COALESCE(supervisor_name, 'dr. Andi Wijaya, Sp.OT') AS name,
		       COUNT(*) AS total_logs,
		       COUNT(CASE WHEN status = 'disetujui' THEN 1 END) AS approved_logs,
		       COUNT(CASE WHEN status = 'menunggu' THEN 1 END) AS pending_logs
		FROM tindakans
		WHERE ` + userFilterCond + `
		GROUP BY supervisor_name
	`
	err := r.db.SelectContext(ctx, &items, query, username, name)
	if items == nil {
		items = []DPJPStatItem{}
	}
	return items, err
}

func (r *repository) GetTotalBimbingan(ctx context.Context, username, name string) (int, error) {
	var tCount, mCount, dCount int
	_ = r.db.GetContext(ctx, &tCount, `SELECT COUNT(*) FROM tindakans WHERE `+userFilterCond, username, name)
	_ = r.db.GetContext(ctx, &mCount, `SELECT COUNT(*) FROM pendidikan_mini_cex WHERE `+userFilterCond, username, name)
	_ = r.db.GetContext(ctx, &dCount, `SELECT COUNT(*) FROM pendidikan_dops WHERE `+userFilterCond, username, name)
	return tCount + mCount + dCount, nil
}

func (r *repository) GetMenungguBimbinganCount(ctx context.Context, username, name string) (int, error) {
	var tCount, mCount int
	_ = r.db.GetContext(ctx, &tCount, `SELECT COUNT(*) FROM tindakans WHERE status = 'menunggu' AND `+userFilterCond, username, name)
	_ = r.db.GetContext(ctx, &mCount, `SELECT COUNT(*) FROM pendidikan_mini_cex WHERE status = 'Menunggu Validasi' AND `+userFilterCond, username, name)
	return tCount + mCount, nil
}

func (r *repository) GetDisetujuiBimbinganCount(ctx context.Context, username, name string) (int, error) {
	var tCount, mCount int
	_ = r.db.GetContext(ctx, &tCount, `SELECT COUNT(*) FROM tindakans WHERE status = 'disetujui' AND `+userFilterCond, username, name)
	_ = r.db.GetContext(ctx, &mCount, `SELECT COUNT(*) FROM pendidikan_mini_cex WHERE status = 'Disetujui' AND `+userFilterCond, username, name)
	return tCount + mCount, nil
}