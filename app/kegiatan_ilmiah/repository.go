package kegiatan_ilmiah

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type KegiatanRepository interface {
	Create(ctx context.Context, k *KegiatanIlmiah) error
	FindAll(ctx context.Context, userID int) ([]KegiatanIlmiah, error)
	FindByID(ctx context.Context, id int) (*KegiatanIlmiah, error)
	FindByKategori(ctx context.Context, kategori string) ([]KegiatanIlmiah, error)
	Delete(ctx context.Context, id int) error
}

type BimbinganRepository interface {
	Create(ctx context.Context, b *BimbinganPenelitian) error
	FindAll(ctx context.Context) ([]BimbinganPenelitian, error)
	FindByID(ctx context.Context, id int) (*BimbinganPenelitian, error)
}

type kegiatanRepo struct {
	db *sqlx.DB
}

type bimbinganRepo struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) KegiatanRepository {
	return &kegiatanRepo{db: db}
}

func NewBimbinganRepository(db *sqlx.DB) BimbinganRepository {
	return &bimbinganRepo{db: db}
}

func (r *kegiatanRepo) Create(ctx context.Context, k *KegiatanIlmiah) error {
	query := `
		INSERT INTO kegiatan_ilmiah (
			user_username, program_studi, ppds_name, nim_nip,
			kategori, jenis_kegiatan, topik, tanggal_mulai, tanggal_selesai,
			lokasi_tipe, lokasi_detail, sebagai, pembimbing_1, pembimbing_2, pembimbing_3,
			pembimbing_4, pembimbing_5, penguji_1, penguji_2, penguji_3, penguji_4, penguji_5,
			deskripsi, lampiran_path, status, updated_at
		) VALUES (
			$1, $2, $3, $4,
			$5, $6, $7, $8, $9,
			$10, $11, $12, $13, $14, $15,
			$16, $17, $18, $19, $20, $21, $22,
			$23, $24, $25, NOW()
		)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRowContext(ctx, query,
		k.UserUsername, k.ProgramStudi, k.PPDSName, k.NIM_NIP,
		k.Kategori, k.JenisKegiatan, k.Topik, k.TanggalMulai, k.TanggalSelesai,
		k.LokasiTipe, k.LokasiDetail, k.Sebagai, k.Pembimbing1, k.Pembimbing2, k.Pembimbing3,
		k.Pembimbing4, k.Pembimbing5, k.Penguji1, k.Penguji2, k.Penguji3, k.Penguji4, k.Penguji5,
		k.Deskripsi, k.LampiranPath, k.Status,
	).Scan(&k.ID, &k.CreatedAt, &k.UpdatedAt)
}

func (r *kegiatanRepo) FindAll(ctx context.Context, userID int) ([]KegiatanIlmiah, error) {
	var list []KegiatanIlmiah
	var err error

	if userID > 0 {
		query := `SELECT * FROM kegiatan_ilmiah WHERE residen_id = $1 ORDER BY id DESC`
		err = r.db.SelectContext(ctx, &list, query, userID)
	} else {
		query := `SELECT * FROM kegiatan_ilmiah ORDER BY id DESC`
		err = r.db.SelectContext(ctx, &list, query)
	}

	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *kegiatanRepo) FindByID(ctx context.Context, id int) (*KegiatanIlmiah, error) {
	var k KegiatanIlmiah
	query := `SELECT * FROM kegiatan_ilmiah WHERE id = $1`
	err := r.db.GetContext(ctx, &k, query, id)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *kegiatanRepo) FindByKategori(ctx context.Context, kategori string) ([]KegiatanIlmiah, error) {
	var list []KegiatanIlmiah
	query := `SELECT * FROM kegiatan_ilmiah WHERE kategori = $1 ORDER BY id DESC`
	err := r.db.SelectContext(ctx, &list, query, kategori)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *kegiatanRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM kegiatan_ilmiah WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *bimbinganRepo) Create(ctx context.Context, b *BimbinganPenelitian) error {
	query := `
		INSERT INTO bimbingan_penelitian (
			user_username, sesi_ke, tahap, topik_bimbingan, tanggal, pembimbing,
			catatan_pembimbing, status_acc, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, NOW()
		)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRowContext(ctx, query,
		b.UserUsername, b.SesiKe, b.Tahap, b.TopikBimbingan, b.Tanggal, b.Pembimbing,
		b.CatatanPembimbing, b.StatusAcc,
	).Scan(&b.ID, &b.CreatedAt, &b.UpdatedAt)
}

func (r *bimbinganRepo) FindAll(ctx context.Context) ([]BimbinganPenelitian, error) {
	var list []BimbinganPenelitian
	query := `SELECT * FROM bimbingan_penelitian ORDER BY id DESC`
	err := r.db.SelectContext(ctx, &list, query)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *bimbinganRepo) FindByID(ctx context.Context, id int) (*BimbinganPenelitian, error) {
	var b BimbinganPenelitian
	query := `SELECT * FROM bimbingan_penelitian WHERE id = $1`
	err := r.db.GetContext(ctx, &b, query, id)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
