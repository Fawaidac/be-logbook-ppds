package pendidikan

import (
	"database/sql"
	"time"
)

type PendidikanKompetensi struct {
	ID             int            `db:"id" json:"id"`
	UserUsername   sql.NullString `db:"user_username" json:"user_username"`
	Kode           string         `db:"kode" json:"kode"`
	Nama           sql.NullString `db:"nama" json:"nama"`
	Domain         sql.NullString `db:"domain" json:"domain"`
	LevelTarget    string         `db:"level_target" json:"level_target"`
	TargetLog      int            `db:"target_log" json:"target_log"`
	AchievedLog    int            `db:"achieved_log" json:"achieved_log"`
	Evaluator      sql.NullString `db:"evaluator" json:"evaluator"`
	Status         string         `db:"status" json:"status"`
	TglVerifikasi  sql.NullString `db:"tgl_verifikasi" json:"tgl_verifikasi"`
	Deskripsi      sql.NullString `db:"deskripsi" json:"deskripsi"`
	CreatedAt      time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at" json:"updated_at"`
}

type PendidikanRotasi struct {
	ID           int            `db:"id" json:"id"`
	UserUsername sql.NullString `db:"user_username" json:"user_username"`
	Stase        string         `db:"stase" json:"stase"`
	Lokasi       sql.NullString `db:"lokasi" json:"lokasi"`
	Periode      sql.NullString `db:"periode" json:"periode"`
	Pembimbing   sql.NullString `db:"pembimbing" json:"pembimbing"`
	Kehadiran    sql.NullString `db:"kehadiran" json:"kehadiran"`
	Nilai        sql.NullString `db:"nilai" json:"nilai"`
	Status       string         `db:"status" json:"status"`
	Tanggal      sql.NullString `db:"tanggal" json:"tanggal"`
	Catatan      sql.NullString `db:"catatan" json:"catatan"`
	CreatedAt    time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at" json:"updated_at"`
}

type PendidikanMiniCex struct {
	ID            int            `db:"id" json:"id"`
	UserUsername  sql.NullString `db:"user_username" json:"user_username"`
	Pasien        sql.NullString `db:"pasien" json:"pasien"`
	Fokus         sql.NullString `db:"fokus" json:"fokus"`
	Kasus         sql.NullString `db:"kasus" json:"kasus"`
	Evaluator     sql.NullString `db:"evaluator" json:"evaluator"`
	Skor          sql.NullString `db:"skor" json:"skor"`
	Status        string         `db:"status" json:"status"`
	Tanggal       sql.NullString `db:"tanggal" json:"tanggal"`
	Catatan       sql.NullString `db:"catatan" json:"catatan"`
	ReviziCatatan sql.NullString `db:"revisi_catatan" json:"revisi_catatan"`
	CreatedAt     time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updated_at"`
}

type PendidikanDops struct {
	ID            int            `db:"id" json:"id"`
	UserUsername  sql.NullString `db:"user_username" json:"user_username"`
	Prosedur      sql.NullString `db:"prosedur" json:"prosedur"`
	Kategori      sql.NullString `db:"kategori" json:"kategori"`
	Kesulitan     sql.NullString `db:"kesulitan" json:"kesulitan"`
	Supervisor    sql.NullString `db:"supervisor" json:"supervisor"`
	Skor          sql.NullString `db:"skor" json:"skor"`
	Status        string         `db:"status" json:"status"`
	Tanggal       sql.NullString `db:"tanggal" json:"tanggal"`
	Catatan       sql.NullString `db:"catatan" json:"catatan"`
	ReviziCatatan sql.NullString `db:"revisi_catatan" json:"revisi_catatan"`
	CreatedAt     time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updated_at"`
}

type PendidikanSeminar struct {
	ID            int            `db:"id" json:"id"`
	UserUsername  sql.NullString `db:"user_username" json:"user_username"`
	Judul         sql.NullString `db:"judul" json:"judul"`
	Jenis         sql.NullString `db:"jenis" json:"jenis"`
	Narasumber    sql.NullString `db:"narasumber" json:"narasumber"`
	Ruang         sql.NullString `db:"ruang" json:"ruang"`
	Skor          sql.NullString `db:"skor" json:"skor"`
	Status        string         `db:"status" json:"status"`
	Tanggal       sql.NullString `db:"tanggal" json:"tanggal"`
	Catatan       sql.NullString `db:"catatan" json:"catatan"`
	ReviziCatatan sql.NullString `db:"revisi_catatan" json:"revisi_catatan"`
	CreatedAt     time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updated_at"`
}

type PendidikanCbd struct {
	ID            int            `db:"id" json:"id"`
	UserUsername  sql.NullString `db:"user_username" json:"user_username"`
	Pasien        sql.NullString `db:"pasien" json:"pasien"`
	Topik         sql.NullString `db:"topik" json:"topik"`
	Kategori      sql.NullString `db:"kategori" json:"kategori"`
	Pembimbing    sql.NullString `db:"pembimbing" json:"pembimbing"`
	Kompleksitas  sql.NullString `db:"kompleksitas" json:"kompleksitas"`
	Skor          sql.NullString `db:"skor" json:"skor"`
	Status        string         `db:"status" json:"status"`
	Tanggal       sql.NullString `db:"tanggal" json:"tanggal"`
	Catatan       sql.NullString `db:"catatan" json:"catatan"`
	ReviziCatatan sql.NullString `db:"revisi_catatan" json:"revisi_catatan"`
	CreatedAt     time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updated_at"`
}
