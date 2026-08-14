package approval

// Response DTOs
type ApprovalListResponse struct {
	Tindakan         []TindakanApprovalItem         `json:"tindakan,omitempty"`
	KegiatanIlmiah   []KegiatanApprovalItem         `json:"kegiatan_ilmiah,omitempty"`
	AktivitasKlinik  []AktivitasKlinikApprovalItem  `json:"aktivitas_klinik,omitempty"`
	PendidikanEvaluasi []PendidikanEvaluasiApprovalItem `json:"pendidikan_evaluasi,omitempty"`
}

type TindakanApprovalItem struct {
	ID             int    `json:"id"`
	UserUsername   string `json:"user_username,omitempty"`
	MRNumber       string `json:"mr_number"`
	PatientName    string `json:"patient_name"`
	DiagnosisLabel string `json:"diagnosis_label,omitempty"`
	PlanProcedure  string `json:"plan_procedure,omitempty"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at,omitempty"`
}

type KegiatanApprovalItem struct {
	ID             int    `json:"id"`
	UserUsername   string `json:"user_username,omitempty"`
	Kategori       string `json:"kategori"`
	JenisKegiatan  string `json:"jenis_kegiatan,omitempty"`
	Topik          string `json:"topik,omitempty"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at,omitempty"`
}

type AktivitasKlinikApprovalItem struct {
	ID            int    `json:"id"`
	UserUsername  string `json:"user_username,omitempty"`
	NamaAktivitas string `json:"nama_aktivitas,omitempty"`
	Tanggal       string `json:"tanggal,omitempty"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at,omitempty"`
}

type PendidikanEvaluasiApprovalItem struct {
	ID           int    `json:"id"`
	UserUsername string `json:"user_username,omitempty"`
	JenisEvaluasi string `json:"jenis_evaluasi,omitempty"`
	Tanggal      string `json:"tanggal,omitempty"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at,omitempty"`
}

type ApprovalAction struct {
	ID     int    `uri:"id" binding:"required"`
	Action string `json:"action" binding:"required,oneof=approve reject"`
}

type ApprovalResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
