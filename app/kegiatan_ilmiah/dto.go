package kegiatan_ilmiah

type CreateKegiatanIlmiahRequest struct {
	Kategori      string `json:"kategori" binding:"required,oneof=simposium workshop multidisiplin ilmiah_lain"`
	JenisKegiatan string `json:"jenis_kegiatan" binding:"required"`
	Topik         string `json:"topik" binding:"required"`
	TanggalMulai  string `json:"tanggal_mulai"`
	TanggalSelesai string `json:"tanggal_selesai"`
	LokasiTipe    string `json:"lokasi_tipe" binding:"oneof=rsds_fk_unair luar_rsds_fk_unair"`
	LokasiDetail  string `json:"lokasi_detail"`
	Sebagai       string `json:"sebagai"`
	Pembimbing1   string `json:"pembimbing_1"`
	Pembimbing2   string `json:"pembimbing_2"`
	Pembimbing3   string `json:"pembimbing_3"`
	Pembimbing4   string `json:"pembimbing_4"`
	Pembimbing5   string `json:"pembimbing_5"`
	Penguji1      string `json:"penguji_1"`
	Penguji2      string `json:"penguji_2"`
	Penguji3      string `json:"penguji_3"`
	Penguji4      string `json:"penguji_4"`
	Penguji5      string `json:"penguji_5"`
	Deskripsi     string `json:"deskripsi"`
}

type CreateBimbinganRequest struct {
	Tahap            string `json:"tahap" binding:"required"`
	TopikBimbingan   string `json:"topik_bimbingan" binding:"required"`
	Pembimbing       string `json:"pembimbing" binding:"required"`
	CatatanPembimbing string `json:"catatan_pembimbing" binding:"required"`
}

type KegiatanIlmiahResponse struct {
	ID             int    `json:"id"`
	UserUsername   string `json:"user_username,omitempty"`
	ProgramStudi   string `json:"program_studi,omitempty"`
	PPDSName       string `json:"ppds_name,omitempty"`
	NIM_NIP        string `json:"nim_nip,omitempty"`
	Kategori       string `json:"kategori"`
	JenisKegiatan  string `json:"jenis_kegiatan"`
	Topik          string `json:"topik"`
	TanggalMulai   string `json:"tanggal_mulai,omitempty"`
	TanggalSelesai string `json:"tanggal_selesai,omitempty"`
	LokasiTipe     string `json:"lokasi_tipe"`
	LokasiDetail   string `json:"lokasi_detail"`
	Sebagai        string `json:"sebagai"`
	Pembimbing1    string `json:"pembimbing_1"`
	Pembimbing2    string `json:"pembimbing_2"`
	Pembimbing3    string `json:"pembimbing_3"`
	Pembimbing4    string `json:"pembimbing_4"`
	Pembimbing5    string `json:"pembimbing_5"`
	Penguji1       string `json:"penguji_1"`
	Penguji2       string `json:"penguji_2"`
	Penguji3       string `json:"penguji_3"`
	Penguji4       string `json:"penguji_4"`
	Penguji5       string `json:"penguji_5"`
	Deskripsi      string `json:"deskripsi"`
	LampiranPath   string `json:"lampiran_path"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type BimbinganResponse struct {
	ID                int    `json:"id"`
	UserUsername      string `json:"user_username,omitempty"`
	SesiKe            string `json:"sesi_ke"`
	Tahap             string `json:"tahap"`
	TopikBimbingan    string `json:"topik_bimbingan"`
	Tanggal           string `json:"tanggal,omitempty"`
	Pembimbing        string `json:"pembimbing"`
	CatatanPembimbing string `json:"catatan_pembimbing"`
	StatusAcc         string `json:"status_acc"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}
