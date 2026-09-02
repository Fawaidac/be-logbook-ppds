package tindakan

import (
	"context"
	"testing"
)

type fakeRepo struct {
	created *Tindakan
}

func (r *fakeRepo) Create(ctx context.Context, t *Tindakan) error {
	r.created = t
	return nil
}

func (r *fakeRepo) FindDPJP(ctx context.Context) ([]string, error) { return nil, nil }
func (r *fakeRepo) FindAll(ctx context.Context, userName string) ([]Tindakan, error) {
	return []Tindakan{}, nil
}
func (r *fakeRepo) FindByID(ctx context.Context, id int) (*Tindakan, error) { return nil, nil }
func (r *fakeRepo) Update(ctx context.Context, t *Tindakan) error           { return nil }
func (r *fakeRepo) UpdateStatus(ctx context.Context, id int, status string) error {
	return nil
}
func (r *fakeRepo) Delete(ctx context.Context, id int) error { return nil }

func TestCreateSetsLoggedInUserName(t *testing.T) {
	repo := &fakeRepo{}
	svc := NewService(repo)

	res, err := svc.Create(context.Background(), CreateTindakanRequest{
		MRNumber:       "MR001",
		PatientName:    "Pasien",
		DiagnosisLabel: "Dx",
		PlanProcedure:  "Op",
	}, "dr. Ratna Puspita")
	if err != nil {
		t.Fatal(err)
	}

	if repo.created.UserUsername.String != "dr. Ratna Puspita" || !repo.created.UserUsername.Valid {
		t.Fatalf("created user_username = %#v, want valid dr. Ratna Puspita", repo.created.UserUsername)
	}
	if res.UserUsername != "dr. Ratna Puspita" {
		t.Fatalf("response user_username = %q, want dr. Ratna Puspita", res.UserUsername)
	}
}
