package approval

import (
	"context"
	"errors"
)

type Service interface {
	GetMenunggu(ctx context.Context, supervisorName string) (*ApprovalListResponse, error)
	GetDisetujui(ctx context.Context, supervisorName string) (*ApprovalListResponse, error)
	GetDitolak(ctx context.Context, supervisorName string) (*ApprovalListResponse, error)

	ApproveTindakan(ctx context.Context, id int, supervisorName string) error
	RejectTindakan(ctx context.Context, id int, supervisorName string) error

	ApproveKegiatanIlmiah(ctx context.Context, id int, supervisorName string) error
	RejectKegiatanIlmiah(ctx context.Context, id int, supervisorName string) error

	ApproveAktivitasKlinik(ctx context.Context, id int, supervisorName string) error
	RejectAktivitasKlinik(ctx context.Context, id int, supervisorName string) error

	ApprovePendidikanEvaluasi(ctx context.Context, id int, supervisorName string) error
	RejectPendidikanEvaluasi(ctx context.Context, id int, supervisorName string) error
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

func (s *service) GetMenunggu(ctx context.Context, supervisorName string) (*ApprovalListResponse, error) {
	tindakan, err := s.tindakanRepo.FindByStatus(ctx, "menunggu", supervisorName)
	if err != nil && err != context.Canceled {
		tindakan = []TindakanApprovalItem{}
	}

	kegiatan, err := s.kegiatanRepo.FindByStatus(ctx, "pending", supervisorName)
	if err != nil && err != context.Canceled {
		kegiatan = []KegiatanApprovalItem{}
	}

	aktivitas, err := s.aktivitasRepo.FindByStatus(ctx, "Menunggu Validasi", supervisorName)
	if err != nil && err != context.Canceled {
		aktivitas = []AktivitasKlinikApprovalItem{}
	}

	pendidikan, err := s.pendidikanEvalRepo.FindByStatus(ctx, "Menunggu Validasi", supervisorName)
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

func (s *service) GetDisetujui(ctx context.Context, supervisorName string) (*ApprovalListResponse, error) {
	tindakan, _ := s.tindakanRepo.FindByStatus(ctx, "disetujui", supervisorName)
	kegiatan, _ := s.kegiatanRepo.FindByStatus(ctx, "disetujui", supervisorName)
	aktivitas, _ := s.aktivitasRepo.FindByStatus(ctx, "Disetujui", supervisorName)
	pendidikan, _ := s.pendidikanEvalRepo.FindByStatus(ctx, "Disetujui", supervisorName)

	return &ApprovalListResponse{
		Tindakan:           tindakan,
		KegiatanIlmiah:     kegiatan,
		AktivitasKlinik:    aktivitas,
		PendidikanEvaluasi: pendidikan,
	}, nil
}

func (s *service) GetDitolak(ctx context.Context, supervisorName string) (*ApprovalListResponse, error) {
	tindakan, _ := s.tindakanRepo.FindByStatus(ctx, "ditolak", supervisorName)
	kegiatan, _ := s.kegiatanRepo.FindByStatus(ctx, "ditolak", supervisorName)
	aktivitas, _ := s.aktivitasRepo.FindByStatus(ctx, "Perlu Revisi", supervisorName)
	pendidikan, _ := s.pendidikanEvalRepo.FindByStatus(ctx, "Perlu Revisi", supervisorName)

	return &ApprovalListResponse{
		Tindakan:           tindakan,
		KegiatanIlmiah:     kegiatan,
		AktivitasKlinik:    aktivitas,
		PendidikanEvaluasi: pendidikan,
	}, nil
}

// ensureSupervisorOwnership memvalidasi bahwa supervisor hanya boleh
// memproses data yang dia tercatat sebagai DPJP/pembimbingnya.
func ensureSupervisorOwnership(owned bool, err error) error {
	if err != nil {
		return errors.New("gagal memverifikasi kepemilikan data")
	}
	if !owned {
		return errors.New("data tidak ditemukan atau bukan di bawah supervisi Anda")
	}
	return nil
}

func (s *service) ApproveTindakan(ctx context.Context, id int, supervisorName string) error {
	if supervisorName != "" {
		owned, err := s.tindakanRepo.IsOwnedBySupervisor(ctx, id, supervisorName)
		if e := ensureSupervisorOwnership(owned, err); e != nil {
			return e
		}
	}
	if err := s.tindakanRepo.UpdateStatus(ctx, id, "disetujui"); err != nil {
		return errors.New("gagal menyetujui tindakan")
	}
	return nil
}

func (s *service) RejectTindakan(ctx context.Context, id int, supervisorName string) error {
	if supervisorName != "" {
		owned, err := s.tindakanRepo.IsOwnedBySupervisor(ctx, id, supervisorName)
		if e := ensureSupervisorOwnership(owned, err); e != nil {
			return e
		}
	}
	if err := s.tindakanRepo.UpdateStatus(ctx, id, "ditolak"); err != nil {
		return errors.New("gagal menolak tindakan")
	}
	return nil
}

func (s *service) ApproveKegiatanIlmiah(ctx context.Context, id int, supervisorName string) error {
	if supervisorName != "" {
		owned, err := s.kegiatanRepo.IsOwnedBySupervisor(ctx, id, supervisorName)
		if e := ensureSupervisorOwnership(owned, err); e != nil {
			return e
		}
	}
	if err := s.kegiatanRepo.UpdateStatus(ctx, id, "disetujui"); err != nil {
		return errors.New("gagal menyetujui kegiatan ilmiah")
	}
	return nil
}

func (s *service) RejectKegiatanIlmiah(ctx context.Context, id int, supervisorName string) error {
	if supervisorName != "" {
		owned, err := s.kegiatanRepo.IsOwnedBySupervisor(ctx, id, supervisorName)
		if e := ensureSupervisorOwnership(owned, err); e != nil {
			return e
		}
	}
	if err := s.kegiatanRepo.UpdateStatus(ctx, id, "ditolak"); err != nil {
		return errors.New("gagal menolak kegiatan ilmiah")
	}
	return nil
}

func (s *service) ApproveAktivitasKlinik(ctx context.Context, id int, supervisorName string) error {
	if supervisorName != "" {
		owned, err := s.aktivitasRepo.IsOwnedBySupervisor(ctx, id, supervisorName)
		if e := ensureSupervisorOwnership(owned, err); e != nil {
			return e
		}
	}
	if err := s.aktivitasRepo.UpdateStatus(ctx, id, "Disetujui"); err != nil {
		return errors.New("gagal menyetujui aktivitas klinik")
	}
	return nil
}

func (s *service) RejectAktivitasKlinik(ctx context.Context, id int, supervisorName string) error {
	if supervisorName != "" {
		owned, err := s.aktivitasRepo.IsOwnedBySupervisor(ctx, id, supervisorName)
		if e := ensureSupervisorOwnership(owned, err); e != nil {
			return e
		}
	}
	if err := s.aktivitasRepo.UpdateStatus(ctx, id, "Perlu Revisi"); err != nil {
		return errors.New("gagal menolak aktivitas klinik")
	}
	return nil
}

func (s *service) ApprovePendidikanEvaluasi(ctx context.Context, id int, supervisorName string) error {
	if supervisorName != "" {
		owned, err := s.pendidikanEvalRepo.IsOwnedBySupervisor(ctx, id, supervisorName)
		if e := ensureSupervisorOwnership(owned, err); e != nil {
			return e
		}
	}
	if err := s.pendidikanEvalRepo.UpdateStatus(ctx, id, "Disetujui"); err != nil {
		return errors.New("gagal menyetujui pendidikan evaluasi")
	}
	return nil
}

func (s *service) RejectPendidikanEvaluasi(ctx context.Context, id int, supervisorName string) error {
	if supervisorName != "" {
		owned, err := s.pendidikanEvalRepo.IsOwnedBySupervisor(ctx, id, supervisorName)
		if e := ensureSupervisorOwnership(owned, err); e != nil {
			return e
		}
	}
	if err := s.pendidikanEvalRepo.UpdateStatus(ctx, id, "Perlu Revisi"); err != nil {
		return errors.New("gagal menolak pendidikan evaluasi")
	}
	return nil
}
