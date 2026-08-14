package approval

import (
	"context"
)

type TindakanRepository interface {
	FindByStatus(ctx context.Context, status string) ([]TindakanApprovalItem, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}

type KegiatanIlmiahRepository interface {
	FindByStatus(ctx context.Context, status string) ([]KegiatanApprovalItem, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}

type AktivitasKlinikRepository interface {
	FindByStatus(ctx context.Context, status string) ([]AktivitasKlinikApprovalItem, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}

type PendidikanEvaluasiRepository interface {
	FindByStatus(ctx context.Context, status string) ([]PendidikanEvaluasiApprovalItem, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}
