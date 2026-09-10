package approval

// Response DTOs
type ApprovalListResponse struct {
	Tindakan           []TindakanApprovalItem           `json:"tindakan,omitempty"`
	KegiatanIlmiah     []KegiatanApprovalItem           `json:"kegiatan_ilmiah,omitempty"`
	PendidikanEvaluasi []PendidikanEvaluasiApprovalItem `json:"pendidikan_evaluasi,omitempty"`
}

type TindakanApprovalItem struct {
	ID             int    `db:"id" json:"id"`
	UserUsername   string `db:"user_username" json:"user_username,omitempty"`
	MRNumber       string `db:"mr_number" json:"mr_number"`
	VisitDate      string `db:"visit_date" json:"visit_date,omitempty"`
	PatientName    string `db:"patient_name" json:"patient_name"`
	Gender         string `db:"gender" json:"gender,omitempty"`
	BirthDate      string `db:"birth_date" json:"birth_date,omitempty"`
	Division       string `db:"division" json:"division,omitempty"`
	DiagnosisLabel string `db:"diagnosis_label" json:"diagnosis_label,omitempty"`
	ProcedureCode  string `db:"procedure_code" json:"procedure_code,omitempty"`
	PlanProcedure  string `db:"plan_procedure" json:"plan_procedure,omitempty"`
	Activity       string `db:"activity" json:"activity,omitempty"`
	ProcedureDate  string `db:"procedure_date" json:"procedure_date,omitempty"`
	Room           string `db:"room" json:"room,omitempty"`
	RoleLabel      string `db:"role_label" json:"role_label,omitempty"`
	Kemandirian    string `db:"kemandirian" json:"kemandirian,omitempty"`
	ClinicalNote   string `db:"clinical_note" json:"clinical_note,omitempty"`
	SupervisorName string `db:"supervisor_name" json:"supervisor_name,omitempty"`
	Feedback       string `db:"feedback" json:"feedback,omitempty"`
	Status         string `db:"status" json:"status"`
	CreatedAt      string `db:"created_at" json:"created_at,omitempty"`
}

type KegiatanApprovalItem struct {
	ID                int    `db:"id" json:"id"`
	UserUsername      string `db:"user_username" json:"user_username,omitempty"`
	PPDSName          string `db:"ppds_name" json:"ppds_name,omitempty"`
	Kategori          string `db:"kategori" json:"kategori"`
	JenisKegiatan     string `db:"jenis_kegiatan" json:"jenis_kegiatan,omitempty"`
	Topik             string `db:"topik" json:"topik,omitempty"`
	TanggalMulai      string `db:"tanggal_mulai" json:"tanggal_mulai,omitempty"`
	LokasiTipe        string `db:"lokasi_tipe" json:"lokasi_tipe,omitempty"`
	LokasiDetail      string `db:"lokasi_detail" json:"lokasi_detail,omitempty"`
	Sebagai           string `db:"sebagai" json:"sebagai,omitempty"`
	Pembimbing1       string `db:"pembimbing_1" json:"pembimbing_1,omitempty"`
	CatatanPembimbing string `db:"catatan_pembimbing" json:"catatan_pembimbing,omitempty"`
	Status            string `db:"status" json:"status"`
	CreatedAt         string `db:"created_at" json:"created_at,omitempty"`
}


type PendidikanEvaluasiApprovalItem struct {
	ID            int    `db:"id" json:"id"`
	UserUsername  string `db:"user_username" json:"user_username,omitempty"`
	JenisEvaluasi string `db:"jenis_evaluasi" json:"jenis_evaluasi,omitempty"`
	Tanggal       string `db:"tanggal" json:"tanggal,omitempty"`
	RevisiCatatan string `db:"revisi_catatan" json:"revisi_catatan,omitempty"`
	Status        string `db:"status" json:"status"`
	CreatedAt     string `db:"created_at" json:"created_at,omitempty"`
}

type RejectRequest struct {
	Catatan string `json:"catatan,omitempty" form:"catatan,omitempty"`
}

type ApprovalAction struct {
	ID     int    `uri:"id" binding:"required"`
	Action string `json:"action" binding:"required,oneof=approve reject"`
}

type SupervisorFilter struct {
	SupervisorName string `form:"supervisor_name"`
}

type ApprovalResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
