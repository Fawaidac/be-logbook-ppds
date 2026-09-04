package dashboard

import "time"

type DashboardSummaryResponse struct {
	TotalTindakan    int                  `json:"total_tindakan"`
	MenungguValidasi int                  `json:"menunggu_validasi"`
	PerluRevisi      int                  `json:"perlu_revisi"`
	CapaianStase     int                  `json:"capaian_stase"`
	UpcomingJadwals  []UpcomingJadwalItem `json:"upcoming_jadwals"`
	RecentEntries    []RecentTindakanItem `json:"recent_entries"`
	ChartData        ChartDataResponse    `json:"chart_data"`
}

type UpcomingJadwalItem struct {
	ID       int       `db:"id" json:"id"`
	Title    string    `db:"title" json:"title"`
	Start    time.Time `db:"start_time" json:"start"`
	End      time.Time `db:"end_time" json:"end"`
	AllDay   bool      `db:"all_day" json:"all_day"`
	Type     string    `db:"type" json:"type"`
	Location string    `db:"location" json:"location"`
}

type RecentTindakanItem struct {
	ID             int       `db:"id" json:"id"`
	ProcedureDate  time.Time `db:"procedure_date" json:"procedure_date"`
	DiagnosisLabel string    `db:"diagnosis_label" json:"diagnosis_label"`
	PlanProcedure  string    `db:"plan_procedure" json:"plan_procedure"`
	Kemandirian    string    `db:"kemandirian" json:"kemandirian"`
	SupervisorName string    `db:"supervisor_name" json:"supervisor_name"`
	DpjpName       string    `db:"dpjp_name" json:"dpjp_name"`
	Status         string    `db:"status" json:"status"`
}

type ChartDataResponse struct {
	Months    []string `json:"months"`
	Mandiri   []int    `json:"mandiri"`
	Dibimbing []int    `json:"dibimbing"`
	Observasi []int    `json:"observasi"`
	Seminar   []int    `json:"seminar"`
}

type RekapTindakanItem struct {
	PlanProcedure  string `db:"plan_procedure" json:"plan_procedure"`
	MandiriCount   int    `db:"mandiri_count" json:"mandiri_count"`
	DibimbingCount int    `db:"dibimbing_count" json:"dibimbing_count"`
	ObservasiCount int    `db:"observasi_count" json:"observasi_count"`
	TotalCount     int    `db:"total_count" json:"total_count"`
}

type DPJPStatItem struct {
	Name         string `db:"name" json:"name"`
	TotalLogs    int    `db:"total_logs" json:"total_logs"`
	ApprovedLogs int    `db:"approved_logs" json:"approved_logs"`
	PendingLogs  int    `db:"pending_logs" json:"pending_logs"`
	Initials     string `json:"initials"`
	Pct          int    `json:"pct"`
}

type LaporanSummaryResponse struct {
	RekapTindakan      []RekapTindakanItem `json:"rekap_tindakan"`
	KompetensiList     interface{}         `json:"kompetensi_list"`
	SimposiumCount     int                 `json:"simposium_count"`
	WorkshopCount      int                 `json:"workshop_count"`
	MultidisiplinCount int                 `json:"multidisiplin_count"`
	IlmiahLainCount    int                 `json:"ilmiah_lain_count"`
	TotalBimbingan     int                 `json:"total_bimbingan"`
	MenungguCount      int                 `json:"menunggu_count"`
	DisetujuiCount     int                 `json:"disetujui_count"`
	ResponRate         int                 `json:"respon_rate"`
	DPJPStats          []DPJPStatItem      `json:"dpjp_stats"`
}

type AdminDashboardResponse struct {
	TotalResiden     int                      `json:"total_residen"`
	TotalSupervisor  int                      `json:"total_supervisor"`
	TotalAdmin       int                      `json:"total_admin"`
	TotalRegistrasi  int                      `json:"total_registrasi"`
	TotalTindakan    int                      `json:"total_tindakan"`
	PersenAktivasi   int                      `json:"persen_aktivasi"`
	RegistrasiList   []AdminRegistrasiItem    `json:"registrasi_list"`
	AktivitasList    []AdminAktivitasItem     `json:"aktivitas_list"`
	RoleDistribution RoleDistributionResponse `json:"role_distribution"`
}

type AdminRegistrasiItem struct {
	ID           int       `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	ProgramStudi string    `db:"program_studi" json:"program"`
	University   string    `db:"university" json:"university"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type AdminAktivitasItem struct {
	Kategori string `json:"kategori"`
	Jumlah   int64  `json:"jumlah"`
}

type RoleDistributionResponse struct {
	Labels []string `json:"labels"`
	Values []int    `json:"values"`
}

type SupervisorDashboardResponse struct {
	TotalMenunggu   int                         `json:"total_menunggu"`
	TotalDisetujui  int                         `json:"total_disetujui"`
	TotalRevisi     int                         `json:"total_revisi"`
	TotalResiden    int                         `json:"total_residen"`
	DisetujuiGrowth int                         `json:"disetujui_growth"`
	Chart           SupervisorChartDataResponse `json:"chart"`
	ResidenList     []SupervisorResidenItem     `json:"residen_list"`
	PendingList     []SupervisorPendingItem     `json:"pending_list"`
}

type SupervisorChartDataResponse struct {
	Labels    []string `json:"labels"`
	Disetujui []int    `json:"disetujui"`
	Menunggu  []int    `json:"menunggu"`
	Revisi    []int    `json:"revisi"`
}

type SupervisorResidenItem struct {
	Username  string `db:"username" json:"username"`
	Name      string `db:"name" json:"name"`
	Prodi     string `db:"prodi" json:"prodi"`
	Pending   int    `db:"pending_count" json:"pending"`
	Disetujui int    `db:"disetujui_count" json:"disetujui"`
}

type SupervisorPendingItem struct {
	ID            int       `db:"id" json:"id"`
	Residen       string    `db:"residen_name" json:"residen"`
	Prosedur      string    `db:"prosedur" json:"prosedur"`
	Stase         string    `db:"stase" json:"stase"`
	Kemandirian   string    `db:"kemandirian" json:"kemandirian"`
	ProcedureDate time.Time `db:"procedure_date" json:"procedure_date"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}
