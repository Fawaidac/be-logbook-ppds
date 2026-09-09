package approval

import (
	"context"
)

type TindakanRepository interface {
	FindByStatus(ctx context.Context, status, supervisorName string) ([]TindakanApprovalItem, error)
	UpdateStatus(ctx context.Context, id int, status string) error
	UpdateStatusWithNote(ctx context.Context, id int, status string, catatan string) error
	IsOwnedBySupervisor(ctx context.Context, id int, supervisorName string) (bool, error)
}

type KegiatanIlmiahRepository interface {
	FindByStatus(ctx context.Context, status, supervisorName string) ([]KegiatanApprovalItem, error)
	UpdateStatus(ctx context.Context, id int, status string) error
	UpdateStatusWithNote(ctx context.Context, id int, status string, catatan string) error
	IsOwnedBySupervisor(ctx context.Context, id int, supervisorName string) (bool, error)
}


type PendidikanEvaluasiRepository interface {
	FindByStatus(ctx context.Context, status, supervisorName string) ([]PendidikanEvaluasiApprovalItem, error)
	UpdateStatus(ctx context.Context, id int, status string) error
	UpdateStatusWithNote(ctx context.Context, id int, status string, catatan string) error
	IsOwnedBySupervisor(ctx context.Context, id int, supervisorName string) (bool, error)
}
