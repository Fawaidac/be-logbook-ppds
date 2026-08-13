package kegiatan_ilmiah

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type Service interface {
	CreateKegiatan(ctx context.Context, req CreateKegiatanIlmiahRequest, username, programStudi, ppdsName, nimNip string) (*KegiatanIlmiahResponse, error)
	GetAllKegiatan(ctx context.Context, userID int) ([]KegiatanIlmiahResponse, error)
	GetByID(ctx context.Context, id int) (*KegiatanIlmiahResponse, error)
	GetByKategori(ctx context.Context, kategori string) ([]KegiatanIlmiahResponse, error)
	DeleteKegiatan(ctx context.Context, id int) error
	CreateBimbingan(ctx context.Context, req CreateBimbinganRequest, username string) (*BimbinganResponse, error)
	GetAllBimbingan(ctx context.Context) ([]BimbinganResponse, error)
	GetBimbinganByID(ctx context.Context, id int) (*BimbinganResponse, error)
}

type service struct {
	kegiatanRepo   KegiatanRepository
	bimbinganRepo  BimbinganRepository
}

func NewService(kegiatanRepo KegiatanRepository, bimbinganRepo BimbinganRepository) Service {
	return &service{
		kegiatanRepo:  kegiatanRepo,
		bimbinganRepo: bimbinganRepo,
	}
}

func parseDate(dateStr string) sql.NullTime {
	if dateStr == "" {
		return sql.NullTime{Valid: false}
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		t, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			return sql.NullTime{Valid: false}
		}
	}
	return sql.NullTime{Time: t, Valid: true}
}

func formatDate(nt sql.NullTime) string {
	if !nt.Valid {
		return ""
	}
	return nt.Time.Format("2006-01-02")
}

func (s *service) CreateKegiatan(ctx context.Context, req CreateKegiatanIlmiahRequest, username, programStudi, ppdsName, nimNip string) (*KegiatanIlmiahResponse, error) {
	k := &KegiatanIlmiah{
		UserUsername:  sql.NullString{String: username, Valid: username != ""},
		ProgramStudi:  sql.NullString{String: programStudi, Valid: programStudi != ""},
		PPDSName:      sql.NullString{String: ppdsName, Valid: ppdsName != ""},
		NIM_NIP:       sql.NullString{String: nimNip, Valid: nimNip != ""},
		Kategori:      req.Kategori,
		JenisKegiatan: req.JenisKegiatan,
		Topik:         req.Topik,
		TanggalMulai:  parseDate(req.TanggalMulai),
		TanggalSelesai: parseDate(req.TanggalSelesai),
		LokasiTipe:    req.LokasiTipe,
		LokasiDetail:  sql.NullString{String: req.LokasiDetail, Valid: req.LokasiDetail != ""},
		Sebagai:       sql.NullString{String: req.Sebagai, Valid: req.Sebagai != ""},
		Pembimbing1:   sql.NullString{String: req.Pembimbing1, Valid: req.Pembimbing1 != ""},
		Pembimbing2:   sql.NullString{String: req.Pembimbing2, Valid: req.Pembimbing2 != ""},
		Pembimbing3:   sql.NullString{String: req.Pembimbing3, Valid: req.Pembimbing3 != ""},
		Pembimbing4:   sql.NullString{String: req.Pembimbing4, Valid: req.Pembimbing4 != ""},
		Pembimbing5:   sql.NullString{String: req.Pembimbing5, Valid: req.Pembimbing5 != ""},
		Penguji1:      sql.NullString{String: req.Penguji1, Valid: req.Penguji1 != ""},
		Penguji2:      sql.NullString{String: req.Penguji2, Valid: req.Penguji2 != ""},
		Penguji3:      sql.NullString{String: req.Penguji3, Valid: req.Penguji3 != ""},
		Penguji4:      sql.NullString{String: req.Penguji4, Valid: req.Penguji4 != ""},
		Penguji5:      sql.NullString{String: req.Penguji5, Valid: req.Penguji5 != ""},
		Deskripsi:     sql.NullString{String: req.Deskripsi, Valid: req.Deskripsi != ""},
		Status:        "pending",
	}

	if err := s.kegiatanRepo.Create(ctx, k); err != nil {
		return nil, err
	}

	return s.toKegiatanResponse(k), nil
}

func (s *service) GetAllKegiatan(ctx context.Context, userID int) ([]KegiatanIlmiahResponse, error) {
	list, err := s.kegiatanRepo.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []KegiatanIlmiahResponse
	for _, item := range list {
		responses = append(responses, *s.toKegiatanResponse(&item))
	}
	return responses, nil
}

func (s *service) GetByID(ctx context.Context, id int) (*KegiatanIlmiahResponse, error) {
	k, err := s.kegiatanRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("kegiatan ilmiah tidak ditemukan")
	}
	return s.toKegiatanResponse(k), nil
}

func (s *service) GetByKategori(ctx context.Context, kategori string) ([]KegiatanIlmiahResponse, error) {
	list, err := s.kegiatanRepo.FindByKategori(ctx, kategori)
	if err != nil {
		return nil, err
	}

	var responses []KegiatanIlmiahResponse
	for _, item := range list {
		responses = append(responses, *s.toKegiatanResponse(&item))
	}
	return responses, nil
}

func (s *service) DeleteKegiatan(ctx context.Context, id int) error {
	_, err := s.kegiatanRepo.FindByID(ctx, id)
	if err != nil {
		return errors.New("kegiatan ilmiah tidak ditemukan")
	}
	return s.kegiatanRepo.Delete(ctx, id)
}

func (s *service) CreateBimbingan(ctx context.Context, req CreateBimbinganRequest, username string) (*BimbinganResponse, error) {
	list, err := s.bimbinganRepo.FindAll(ctx)
	if err != nil {
		list = []BimbinganPenelitian{}
	}

	sesiNum := len(list) + 1
	sesiKe := "Sesi " + strconv.Itoa(sesiNum)

	b := &BimbinganPenelitian{
		UserUsername:     sql.NullString{String: username, Valid: username != ""},
		SesiKe:           sesiKe,
		Tahap:            req.Tahap,
		TopikBimbingan:   req.TopikBimbingan,
		Tanggal:          sql.NullTime{Time: time.Now(), Valid: true},
		Pembimbing:       sql.NullString{String: req.Pembimbing, Valid: true},
		CatatanPembimbing: sql.NullString{String: req.CatatanPembimbing, Valid: true},
		StatusAcc:        "Dalam Proses",
	}

	if err := s.bimbinganRepo.Create(ctx, b); err != nil {
		return nil, err
	}

	return s.toBimbinganResponse(b), nil
}

func (s *service) GetAllBimbingan(ctx context.Context) ([]BimbinganResponse, error) {
	list, err := s.bimbinganRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []BimbinganResponse
	for _, item := range list {
		responses = append(responses, *s.toBimbinganResponse(&item))
	}
	return responses, nil
}

func (s *service) GetBimbinganByID(ctx context.Context, id int) (*BimbinganResponse, error) {
	b, err := s.bimbinganRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("bimbingan penelitian tidak ditemukan")
	}
	return s.toBimbinganResponse(b), nil
}



func (s *service) toKegiatanResponse(k *KegiatanIlmiah) *KegiatanIlmiahResponse {
	return &KegiatanIlmiahResponse{
		ID:             k.ID,
		UserUsername:   k.UserUsername.String,
		ProgramStudi:   k.ProgramStudi.String,
		PPDSName:       k.PPDSName.String,
		NIM_NIP:        k.NIM_NIP.String,
		Kategori:       k.Kategori,
		JenisKegiatan:  k.JenisKegiatan,
		Topik:          k.Topik,
		TanggalMulai:   formatDate(k.TanggalMulai),
		TanggalSelesai: formatDate(k.TanggalSelesai),
		LokasiTipe:     k.LokasiTipe,
		LokasiDetail:   k.LokasiDetail.String,
		Sebagai:        k.Sebagai.String,
		Pembimbing1:    k.Pembimbing1.String,
		Pembimbing2:    k.Pembimbing2.String,
		Pembimbing3:    k.Pembimbing3.String,
		Pembimbing4:    k.Pembimbing4.String,
		Pembimbing5:    k.Pembimbing5.String,
		Penguji1:       k.Penguji1.String,
		Penguji2:       k.Penguji2.String,
		Penguji3:       k.Penguji3.String,
		Penguji4:       k.Penguji4.String,
		Penguji5:       k.Penguji5.String,
		Deskripsi:      k.Deskripsi.String,
		LampiranPath:   k.LampiranPath.String,
		Status:         k.Status,
		CreatedAt:      k.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      k.UpdatedAt.Format(time.RFC3339),
	}
}

func (s *service) toBimbinganResponse(b *BimbinganPenelitian) *BimbinganResponse {
	return &BimbinganResponse{
		ID:                b.ID,
		UserUsername:      b.UserUsername.String,
		SesiKe:            b.SesiKe,
		Tahap:             b.Tahap,
		TopikBimbingan:    b.TopikBimbingan,
		Tanggal:           formatDate(b.Tanggal),
		Pembimbing:        b.Pembimbing.String,
		CatatanPembimbing: b.CatatanPembimbing.String,
		StatusAcc:         b.StatusAcc,
		CreatedAt:         b.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         b.UpdatedAt.Format(time.RFC3339),
	}
}
