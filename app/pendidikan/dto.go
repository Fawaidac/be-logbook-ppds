package pendidikan

type CreateKompetensiRequest struct {
	Kode        string `json:"kode" binding:"required"`
	LevelTarget string `json:"level_target" binding:"required"`
	AchievedLog int    `json:"achieved_log" binding:"required"`
	Evaluator   string `json:"evaluator" binding:"required"`
	Deskripsi   string `json:"deskripsi" binding:"required"`
}

type KompetensiResponse struct {
	ID            int    `json:"id"`
	UserUsername  string `json:"user_username,omitempty"`
	Kode          string `json:"kode"`
	Nama          string `json:"nama,omitempty"`
	Domain        string `json:"domain,omitempty"`
	LevelTarget   string `json:"level_target"`
	TargetLog     int    `json:"target_log"`
	AchievedLog   int    `json:"achieved_log"`
	Evaluator     string `json:"evaluator,omitempty"`
	Status        string `json:"status"`
	TglVerifikasi string `json:"tgl_verifikasi,omitempty"`
	Deskripsi     string `json:"deskripsi,omitempty"`
}

type CreateRotasiRequest struct {
	Shift       string `json:"shift"`
	Pembimbing  string `json:"pembimbing"`
	Catatan     string `json:"catatan"`
}

type RotasiResponse struct {
	ID           int    `json:"id"`
	UserUsername string `json:"user_username,omitempty"`
	Stase        string `json:"stase"`
	Lokasi       string `json:"lokasi,omitempty"`
	Periode      string `json:"periode,omitempty"`
	Pembimbing   string `json:"pembimbing,omitempty"`
	Kehadiran    string `json:"kehadiran,omitempty"`
	Nilai        string `json:"nilai,omitempty"`
	Status       string `json:"status"`
	Tanggal      string `json:"tanggal,omitempty"`
	Catatan      string `json:"catatan,omitempty"`
}

type CreateMiniCexRequest struct {
	Pasien    string `json:"pasien"`
	Fokus     string `json:"fokus"`
	Kasus     string `json:"kasus"`
	Evaluator string `json:"evaluator"`
	Catatan   string `json:"catatan"`
}

type MiniCexResponse struct {
	ID           int    `json:"id"`
	UserUsername string `json:"user_username,omitempty"`
	Pasien       string `json:"pasien,omitempty"`
	Fokus        string `json:"fokus,omitempty"`
	Kasus        string `json:"kasus,omitempty"`
	Evaluator    string `json:"evaluator,omitempty"`
	Skor         string `json:"skor,omitempty"`
	Status       string `json:"status"`
	Tanggal      string `json:"tanggal,omitempty"`
	Catatan      string `json:"catatan,omitempty"`
}

type CreateDopsRequest struct {
	Prosedur   string `json:"prosedur"`
	Kesulitan  string `json:"kesulitan"`
	Supervisor string `json:"supervisor"`
	Catatan    string `json:"catatan"`
}

type DopsResponse struct {
	ID           int    `json:"id"`
	UserUsername string `json:"user_username,omitempty"`
	Prosedur     string `json:"prosedur,omitempty"`
	Kategori     string `json:"kategori,omitempty"`
	Kesulitan    string `json:"kesulitan,omitempty"`
	Supervisor   string `json:"supervisor,omitempty"`
	Skor         string `json:"skor,omitempty"`
	Status       string `json:"status"`
	Tanggal      string `json:"tanggal,omitempty"`
	Catatan      string `json:"catatan,omitempty"`
}

type CreateSeminarRequest struct {
	Judul      string `json:"judul"`
	Jenis      string `json:"jenis"`
	Narasumber string `json:"narasumber"`
	Catatan    string `json:"catatan"`
}

type SeminarResponse struct {
	ID           int    `json:"id"`
	UserUsername string `json:"user_username,omitempty"`
	Judul        string `json:"judul,omitempty"`
	Jenis        string `json:"jenis,omitempty"`
	Narasumber   string `json:"narasumber,omitempty"`
	Ruang        string `json:"ruang,omitempty"`
	Skor         string `json:"skor,omitempty"`
	Status       string `json:"status"`
	Tanggal      string `json:"tanggal,omitempty"`
	Catatan      string `json:"catatan,omitempty"`
}

type CreateCbdRequest struct {
	Pasien       string `json:"pasien"`
	Topik        string `json:"topik"`
	Kategori     string `json:"kategori"`
	Pembimbing   string `json:"pembimbing"`
	Catatan      string `json:"catatan"`
}

type CbdResponse struct {
	ID            int    `json:"id"`
	UserUsername  string `json:"user_username,omitempty"`
	Pasien        string `json:"pasien,omitempty"`
	Topik         string `json:"topik,omitempty"`
	Kategori      string `json:"kategori,omitempty"`
	Pembimbing    string `json:"pembimbing,omitempty"`
	Kompleksitas  string `json:"kompleksitas,omitempty"`
	Skor          string `json:"skor,omitempty"`
	Status        string `json:"status"`
	Tanggal       string `json:"tanggal,omitempty"`
	Catatan       string `json:"catatan,omitempty"`
}
