package kegiatan_ilmiah

import (
	"database/sql"
	"time"
)

type KegiatanIlmiah struct {
	ID            int            `db:"id" json:"id"`
	UserUsername  sql.NullString `db:"user_username" json:"user_username"`
	ProgramStudi  sql.NullString `db:"program_studi" json:"program_studi"`
	PPDSName      sql.NullString `db:"ppds_name" json:"ppds_name"`
	NIM_NIP       sql.NullString `db:"nim_nip" json:"nim_nip"`
	Kategori      string         `db:"kategori" json:"kategori"`
	JenisKegiatan string         `db:"jenis_kegiatan" json:"jenis_kegiatan"`
	Topik         string         `db:"topik" json:"topik"`
	TanggalMulai  sql.NullTime   `db:"tanggal_mulai" json:"tanggal_mulai"`
	TanggalSelesai sql.NullTime  `db:"tanggal_selesai" json:"tanggal_selesai"`
	LokasiTipe    string         `db:"lokasi_tipe" json:"lokasi_tipe"`
	LokasiDetail  sql.NullString `db:"lokasi_detail" json:"lokasi_detail"`
	Sebagai       sql.NullString `db:"sebagai" json:"sebagai"`
	Pembimbing1   sql.NullString `db:"pembimbing_1" json:"pembimbing_1"`
	Pembimbing2   sql.NullString `db:"pembimbing_2" json:"pembimbing_2"`
	Pembimbing3   sql.NullString `db:"pembimbing_3" json:"pembimbing_3"`
	Pembimbing4   sql.NullString `db:"pembimbing_4" json:"pembimbing_4"`
	Pembimbing5   sql.NullString `db:"pembimbing_5" json:"pembimbing_5"`
	Penguji1      sql.NullString `db:"penguji_1" json:"penguji_1"`
	Penguji2      sql.NullString `db:"penguji_2" json:"penguji_2"`
	Penguji3      sql.NullString `db:"penguji_3" json:"penguji_3"`
	Penguji4      sql.NullString `db:"penguji_4" json:"penguji_4"`
	Penguji5      sql.NullString `db:"penguji_5" json:"penguji_5"`
	Deskripsi     sql.NullString `db:"deskripsi" json:"deskripsi"`
	LampiranPath  sql.NullString `db:"lampiran_path" json:"lampiran_path"`
	Status        string         `db:"status" json:"status"`
	CreatedAt     time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updated_at"`
}

type BimbinganPenelitian struct {
	ID               int            `db:"id" json:"id"`
	UserUsername     sql.NullString `db:"user_username" json:"user_username"`
	SesiKe           string         `db:"sesi_ke" json:"sesi_ke"`
	Tahap            string         `db:"tahap" json:"tahap"`
	TopikBimbingan   string         `db:"topik_bimbingan" json:"topik_bimbingan"`
	Tanggal          sql.NullTime   `db:"tanggal" json:"tanggal"`
	Pembimbing       sql.NullString `db:"pembimbing" json:"pembimbing"`
	CatatanPembimbing sql.NullString `db:"catatan_pembimbing" json:"catatan_pembimbing"`
	StatusAcc        string         `db:"status_acc" json:"status_acc"`
	CreatedAt        time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time      `db:"updated_at" json:"updated_at"`
}
