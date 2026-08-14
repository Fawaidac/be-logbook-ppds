package approval

import (
	"context"
	"errors"
)

type Service interface {
	GetMenunggu(ctx context.Context) (*ApprovalListResponse, error)
	GetDisetujui(ctx context.Context) (*ApprovalListResponse, error)
	GetDitolak(ctx context.Context) (*ApprovalListResponse, error)

	ApproveTindakan(ctx context.Context, id int) error
	RejectTindakan(ctx context.Context, id int) error

	ApproveKegiatanIlmiah(ctx context.Context, id int) error
	RejectKegiatanIlmiah(ctx context.Context, id int) error

	ApproveAktivitasKlinik(ctx context.Context, id int) error
	RejectAktivitasKlinik(ctx context.Context, id int) error

	ApprovePendidikanEvaluasi(ctx context.Context, id int) error
	RejectPendidikanEvaluasi(ctx context.Context, id int) error
}

type service struct {
	tindakanRepo       TindakanRepository
	kegiatanRepo       KegiatanIlmiahRepository
	aktivitasRepo      AktivitasKlinikRepository
	pendidikanEvalRepo PendidikanEvaluasiRepository
}

func NewService(
	tindakanRepo TindakanRepository,
	kegiatanRepo KegiatanIlmiahRepository,
	aktivitasRepo AktivitasKlinikRepository,
	pendidikanEvalRepo PendidikanEvaluasiRepository,
) Service {
	return &service{
		tindakanRepo:       tindakanRepo,
		kegiatanRepo:       kegiatanRepo,
		aktivitasRepo:      aktivitasRepo,
		pendidikanEvalRepo: pendidikanEvalRepo,
	}
}

func (s *service) GetMenunggu(ctx context.Context) (*ApprovalListResponse, error) {
	tindakan, err := s.tindakanRepo.FindByStatus(ctx, "menunggu")
	if err != nil && err != context.Canceled {
		tindakan = []TindakanApprovalItem{}
	}

	kegiatan, err := s.kegiatanRepo.FindByStatus(ctx, "pending")
	if err != nil && err != context.Canceled {
		kegiatan = []KegiatanApprovalItem{}
	}

	aktivitas, err := s.aktivitasRepo.FindByStatus(ctx, "Menunggu Validasi")
	if err != nil && err != context.Canceled {
		aktivitas = []AktivitasKlinikApprovalItem{}
	}

	pendidikan, err := s.pendidikanEvalRepo.FindByStatus(ctx, "Menunggu Validasi")
	if err != nil && err != context.Canceled {
		pendidikan = []PendidikanEvaluasiApprovalItem{}
	}

	return &ApprovalListResponse{
		Tindakan:           tindakan,
		KegiatanIlmiah:     kegiatan,
		AktivitasKlinik:    aktivitas,
		PendidikanEvaluasi: pendidikan,
	}, nil
}

func (s *service) GetDisetujui(ctx context.Context) (*ApprovalListResponse, error) {
	tindakan, _ := s.tindakanRepo.FindByStatus(ctx, "disetujui")
	kegiatan, _ := s.kegiatanRepo.FindByStatus(ctx, "disetujui")
	aktivitas, _ := s.aktivitasRepo.FindByStatus(ctx, "Disetujui")
	pendidikan, _ := s.pendidikanEvalRepo.FindByStatus(ctx, "Disetujui")

	return &ApprovalListResponse{
		Tindakan:           tindakan,
		KegiatanIlmiah:     kegiatan,
		AktivitasKlinik:    aktivitas,
		PendidikanEvaluasi: pendidikan,
	}, nil
}

func (s *service) GetDitolak(ctx context.Context) (*ApprovalListResponse, error) {
	tindakan, _ := s.tindakanRepo.FindByStatus(ctx, "ditolak")
	kegiatan, _ := s.kegiatanRepo.FindByStatus(ctx, "ditolak")
	aktivitas, _ := s.aktivitasRepo.FindByStatus(ctx, "Perlu Revisi")
	pendidikan, _ := s.pendidikanEvalRepo.FindByStatus(ctx, "Perlu Revisi")

	return &ApprovalListResponse{
		Tindakan:           tindakan,
		KegiatanIlmiah:     kegiatan,
		AktivitasKlinik:    aktivitas,
		PendidikanEvaluasi: pendidikan,
	}, nil
}

func (s *service) ApproveTindakan(ctx context.Context, id int) error {
	if err := s.tindakanRepo.UpdateStatus(ctx, id, "disetujui"); err != nil {
		return errors.New("gagal menyetujui tindakan")
	}
	return nil
}

func (s *service) RejectTindakan(ctx context.Context, id int) error {
	if err := s.tindakanRepo.UpdateStatus(ctx, id, "ditolak"); err != nil {
		return errors.New("gagal menolak tindakan")
	}
	return nil
}

func (s *service) ApproveKegiatanIlmiah(ctx context.Context, id int) error {
	if err := s.kegiatanRepo.UpdateStatus(ctx, id, "disetujui"); err != nil {
		return errors.New("gagal menyetujui kegiatan ilmiah")
	}
	return nil
}

func (s *service) RejectKegiatanIlmiah(ctx context.Context, id int) error {
	if err := s.kegiatanRepo.UpdateStatus(ctx, id, "ditolak"); err != nil {
		return errors.New("gagal menolak kegiatan ilmiah")
	}
	return nil
}

func (s *service) ApproveAktivitasKlinik(ctx context.Context, id int) error {
	if err := s.aktivitasRepo.UpdateStatus(ctx, id, "Disetujui"); err != nil {
		return errors.New("gagal menyetujui aktivitas klinik")
	}
	return nil
}

func (s *service) RejectAktivitasKlinik(ctx context.Context, id int) error {
	if err := s.aktivitasRepo.UpdateStatus(ctx, id, "Perlu Revisi"); err != nil {
		return errors.New("gagal menolak aktivitas klinik")
	}
	return nil
}

func (s *service) ApprovePendidikanEvaluasi(ctx context.Context, id int) error {
	if err := s.pendidikanEvalRepo.UpdateStatus(ctx, id, "Disetujui"); err != nil {
		return errors.New("gagal menyetujui pendidikan evaluasi")
	}
	return nil
}

func (s *service) RejectPendidikanEvaluasi(ctx context.Context, id int) error {
	if err := s.pendidikanEvalRepo.UpdateStatus(ctx, id, "Perlu Revisi"); err != nil {
		return errors.New("gagal menolak pendidikan evaluasi")
	}
	return nil
}
