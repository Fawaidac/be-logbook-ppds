package tindakan

type CreateTindakanRequest struct {
	MRNumber       string `json:"mr_number" binding:"required"`
	VisitDate      string `json:"visit_date"`
	PatientName    string `json:"patient_name" binding:"required"`
	Gender         string `json:"gender"`
	BirthDate      string `json:"birth_date"`
	Division       string `json:"division"`
	DiagnosisLabel string `json:"diagnosis_label" binding:"required"`
	ProcedureCode  string `json:"procedure_code"`
	PlanProcedure  string `json:"plan_procedure" binding:"required"`
	Activity       string `json:"activity"`
	ProcedureDate  string `json:"procedure_date"`
	RoomLabel      string `json:"room_label"`
	RoleLabel      string `json:"role_label"`
	Kemandirian    string `json:"kemandirian"`
	ClinicalNote   string `json:"clinical_note"`
	SupervisorName string `json:"supervisor_name"`
}

type UpdateTindakanRequest struct {
	MRNumber       string `json:"mr_number" binding:"required"`
	PatientName    string `json:"patient_name" binding:"required"`
	PlanProcedure  string `json:"plan_procedure" binding:"required"`
	DiagnosisLabel string `json:"diagnosis_label"`
	RoomLabel      string `json:"room_label"`
	ClinicalNote   string `json:"clinical_note"`
}

type TindakanResponse struct {
	ID             int    `json:"id"`
	UserUsername   string `json:"user_username,omitempty"`
	MRNumber       string `json:"mr_number"`
	VisitDate      string `json:"visit_date,omitempty"`
	PatientName    string `json:"patient_name"`
	Gender         string `json:"gender"`
	BirthDate      string `json:"birth_date,omitempty"`
	Division       string `json:"division"`
	DiagnosisLabel string `json:"diagnosis_label"`
	ProcedureCode  string `json:"procedure_code"`
	PlanProcedure  string `json:"plan_procedure"`
	Activity       string `json:"activity"`
	ProcedureDate  string `json:"procedure_date,omitempty"`
	Room           string `json:"room"`
	Role           string `json:"role"`
	Kemandirian    string `json:"kemandirian"`
	ClinicalNote   string `json:"clinical_note"`
	SupervisorName string `json:"supervisor_name"`
	Status         string `json:"status"`
	Feedback       string `json:"feedback"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type SummaryResponse struct {
	TotalCount        int                `json:"total_count"`
	MandiriCount      int                `json:"mandiri_count"`
	DibimbingCount    int                `json:"dibimbing_count"`
	ObservasiCount    int                `json:"observasi_count"`
	DisetujuiCount    int                `json:"disetujui_count"`
	VerifikasiPercent int                `json:"verifikasi_percent"`
	Entries           []TindakanResponse `json:"entries"`
}
