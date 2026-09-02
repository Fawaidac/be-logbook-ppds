package kegiatan_ilmiah

import (
	"context"
	"testing"
)

type fakeKegiatanRepo struct {
	findAllUserName string
}

func (r *fakeKegiatanRepo) Create(ctx context.Context, k *KegiatanIlmiah) error { return nil }
func (r *fakeKegiatanRepo) FindAll(ctx context.Context, userName string) ([]KegiatanIlmiah, error) {
	r.findAllUserName = userName
	return []KegiatanIlmiah{}, nil
}
func (r *fakeKegiatanRepo) FindByID(ctx context.Context, id int) (*KegiatanIlmiah, error) {
	return nil, nil
}
func (r *fakeKegiatanRepo) FindByKategori(ctx context.Context, kategori string) ([]KegiatanIlmiah, error) {
	return []KegiatanIlmiah{}, nil
}
func (r *fakeKegiatanRepo) Delete(ctx context.Context, id int) error { return nil }

type fakeBimbinganRepo struct{}

func (r *fakeBimbinganRepo) Create(ctx context.Context, b *BimbinganPenelitian) error { return nil }
func (r *fakeBimbinganRepo) FindAll(ctx context.Context) ([]BimbinganPenelitian, error) {
	return []BimbinganPenelitian{}, nil
}
func (r *fakeBimbinganRepo) FindByID(ctx context.Context, id int) (*BimbinganPenelitian, error) {
	return nil, nil
}

func TestGetAllKegiatanFiltersByLoggedInUsername(t *testing.T) {
	repo := &fakeKegiatanRepo{}
	svc := NewService(repo, &fakeBimbinganRepo{})

	_, err := svc.GetAllKegiatan(context.Background(), "residen01")
	if err != nil {
		t.Fatal(err)
	}

	if repo.findAllUserName != "residen01" {
		t.Fatalf("FindAll username = %q, want residen01", repo.findAllUserName)
	}
}
