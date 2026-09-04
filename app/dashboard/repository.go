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

	// Admin dashboard
	GetCapaianStase(ctx context.Context, username, name string) (int, error)
	CountUsersByRole(ctx context.Context, role string) (int, error)
	CountPendingRegistrations(ctx context.Context) (int, error)
	GetPendingRegistrations(ctx context.Context, limit int) ([]AdminRegistrasiItem, error)
	CountPendidikanRecords(ctx context.Context) (int, error)
	CountKegiatanIlmiahRecords(ctx context.Context) (int, error)
	CountJadwalRecords(ctx context.Context) (int, error)
	CountActiveResidenThisYear(ctx context.Context, year int) (int, error)

	// Supervisor dashboard (scoped: supervisorName kosong = semua)
	CountTindakanByStatus(ctx context.Context, status, supervisorName string) (int, error)
	CountTindakanByStatusInMonth(ctx context.Context, status string, year, month int, supervisorName string) (int, error)
	CountResidenBySupervisor(ctx context.Context, supervisorName string) (int, error)
	GetResidenActivitySummary(ctx context.Context, limit int, supervisorName string) ([]SupervisorResidenItem, error)
	GetPendingTindakanQueue(ctx context.Context, limit int, supervisorName string) ([]SupervisorPendingItem, error)
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

// GetCapaianStase menghitung persentase capaian target kompetensi
// (SUM achieved_log / SUM target_log) untuk user yang login.
func (r *repository) GetCapaianStase(ctx context.Context, username, name string) (int, error) {
	var persen int
	query := `
		SELECT COALESCE(ROUND(SUM(achieved_log)::numeric / NULLIF(SUM(target_log), 0) * 100), 0)::int
		FROM pendidikan_kompetensi
		WHERE ($1 = '' OR user_username = $1 OR ($2 != '' AND user_username = $2))
	`
	err := r.db.GetContext(ctx, &persen, query, username, name)
	if err != nil {
		return 0, err
	}
	return persen, nil
}

func (r *repository) CountUsersByRole(ctx context.Context, role string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE role::text = $1`
	err := r.db.GetContext(ctx, &count, query, role)
	return count, err
}

func (r *repository) CountPendingRegistrations(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM user_registrations WHERE status = 'pending'`
	err := r.db.GetContext(ctx, &count, query)
	return count, err
}

func (r *repository) GetPendingRegistrations(ctx context.Context, limit int) ([]AdminRegistrasiItem, error) {
	var items []AdminRegistrasiItem
	query := `
		SELECT id,
		       name,
		       COALESCE(program_studi, '') AS program_studi,
		       COALESCE(university, '') AS university,
		       created_at
		FROM user_registrations
		WHERE status = 'pending'
		ORDER BY created_at DESC
		LIMIT $1
	`
	err := r.db.SelectContext(ctx, &items, query, limit)
	if items == nil {
		items = []AdminRegistrasiItem{}
	}
	return items, err
}

func (r *repository) CountPendidikanRecords(ctx context.Context) (int, error) {
	var count int
	query := `
		SELECT
			(SELECT COUNT(*) FROM pendidikan_kompetensi)
			+ (SELECT COUNT(*) FROM pendidikan_mini_cex)
			+ (SELECT COUNT(*) FROM pendidikan_dops)
			+ (SELECT COUNT(*) FROM pendidikan_seminar)
			+ (SELECT COUNT(*) FROM pendidikan_cbd)
	`
	err := r.db.GetContext(ctx, &count, query)
	return count, err
}

func (r *repository) CountKegiatanIlmiahRecords(ctx context.Context) (int, error) {
	var count int
	query := `
		SELECT
			(SELECT COUNT(*) FROM kegiatan_ilmiah)
			+ (SELECT COUNT(*) FROM bimbingan_penelitian)
	`
	err := r.db.GetContext(ctx, &count, query)
	return count, err
}

func (r *repository) CountJadwalRecords(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM jadwals`
	err := r.db.GetContext(ctx, &count, query)
	return count, err
}

func (r *repository) CountActiveResidenThisYear(ctx context.Context, year int) (int, error) {
	var count int
	query := `
		SELECT COUNT(DISTINCT t.user_username)
		FROM tindakans t
		JOIN users u ON u.username = t.user_username
		WHERE u.role::text = 'residen'
		  AND EXTRACT(YEAR FROM COALESCE(t.procedure_date, t.created_at)) = $1
	`
	err := r.db.GetContext(ctx, &count, query, year)
	return count, err
}

func (r *repository) CountTindakanByStatus(ctx context.Context, status, supervisorName string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM tindakans WHERE status = $1 AND ($2 = '' OR supervisor_name = $2)`
	err := r.db.GetContext(ctx, &count, query, status, supervisorName)
	return count, err
}

func (r *repository) CountTindakanByStatusInMonth(ctx context.Context, status string, year, month int, supervisorName string) (int, error) {
	var count int
	query := `
		SELECT COUNT(*) FROM tindakans
		WHERE status = $1
		  AND EXTRACT(YEAR FROM COALESCE(procedure_date, created_at)) = $2
		  AND EXTRACT(MONTH FROM COALESCE(procedure_date, created_at)) = $3
		  AND ($4 = '' OR supervisor_name = $4)
	`
	err := r.db.GetContext(ctx, &count, query, status, year, month, supervisorName)
	return count, err
}

// CountResidenBySupervisor menghitung jumlah ppds yang dibimbing supervisor
// (relasi 1 residen -> 1 supervisor via users.supervisor_name).
func (r *repository) CountResidenBySupervisor(ctx context.Context, supervisorName string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE role::text = 'residen' AND ($1 = '' OR supervisor_name = $1)`
	err := r.db.GetContext(ctx, &count, query, supervisorName)
	return count, err
}

// GetResidenActivitySummary merangkum aktivitas validasi per residen:
// jumlah logbook berstatus menunggu dan disetujui, diurutkan dari yang
// paling banyak mengantre. Saat supervisorName terisi, hanya ppds yang
// dibimbing supervisor tersebut yang ditampilkan.
func (r *repository) GetResidenActivitySummary(ctx context.Context, limit int, supervisorName string) ([]SupervisorResidenItem, error) {
	var items []SupervisorResidenItem
	query := `
		SELECT u.username,
		       u.name,
		       COALESCE(u.program_studi, '') AS prodi,
		       COALESCE(SUM(CASE WHEN t.status = 'menunggu' THEN 1 ELSE 0 END), 0)::int AS pending_count,
		       COALESCE(SUM(CASE WHEN t.status = 'disetujui' THEN 1 ELSE 0 END), 0)::int AS disetujui_count
		FROM users u
		LEFT JOIN tindakans t ON t.user_username = u.username
		WHERE u.role::text = 'residen'
		  AND ($1 = '' OR u.supervisor_name = $1)
		GROUP BY u.username, u.name, u.program_studi
		ORDER BY pending_count DESC, disetujui_count DESC
		LIMIT $2
	`
	err := r.db.SelectContext(ctx, &items, query, supervisorName, limit)
	if items == nil {
		items = []SupervisorResidenItem{}
	}
	return items, err
}

// GetPendingTindakanQueue mengambil antrian logbook berstatus menunggu
// terbaru beserta nama residen pemiliknya, difilter berdasarkan DPJP
// yang memilih supervisor (supervisorName kosong = semua).
func (r *repository) GetPendingTindakanQueue(ctx context.Context, limit int, supervisorName string) ([]SupervisorPendingItem, error) {
	var items []SupervisorPendingItem
	query := `
		SELECT t.id,
		       COALESCE(NULLIF(u.name, ''), COALESCE(t.user_username, '')) AS residen_name,
		       t.plan_procedure AS prosedur,
		       COALESCE(t.division, '') AS stase,
		       t.kemandirian,
		       COALESCE(t.procedure_date, t.created_at) AS procedure_date,
		       t.created_at
		FROM tindakans t
		LEFT JOIN users u ON u.username = t.user_username
		WHERE t.status = 'menunggu'
		  AND ($1 = '' OR t.supervisor_name = $1)
		ORDER BY t.created_at DESC
		LIMIT $2
	`
	err := r.db.SelectContext(ctx, &items, query, supervisorName, limit)
	if items == nil {
		items = []SupervisorPendingItem{}
	}
	return items, err
}
