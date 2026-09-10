package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"be-logbook-ppds/app/approval"
	"be-logbook-ppds/configs"
	"be-logbook-ppds/pkg/database"
)

func main() {
	cfg := configs.LoadConfig()
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	tindakanRepo := &approval.TindakanRepoAdapter{DB: db}
	kegiatanRepo := &approval.KegiatanIlmiahRepoAdapter{DB: db}
	pendidikanRepo := &approval.PendidikanEvaluasiRepoAdapter{DB: db}

	service := approval.NewService(tindakanRepo, kegiatanRepo, pendidikanRepo)

	ctx := context.Background()

	// Test 1: Admin (supervisorName = "")
	resAdmin, err := service.GetMenunggu(ctx, "")
	if err != nil {
		log.Fatalf("GetMenunggu admin error: %v", err)
	}
	outAdmin, _ := json.MarshalIndent(resAdmin, "", "  ")
	fmt.Printf("=== GET MENUNGGU (ADMIN / EMPTY SUPERVISOR) ===\n%s\n\n", string(outAdmin))

	// Test 2: Supervisor "dr. Budi Santoso, Sp.B"
	resSup, err := service.GetMenunggu(ctx, "dr. Budi Santoso, Sp.B")
	if err != nil {
		log.Fatalf("GetMenunggu supervisor error: %v", err)
	}
	outSup, _ := json.MarshalIndent(resSup, "", "  ")
	fmt.Printf("=== GET MENUNGGU (SUPERVISOR: dr. Budi Santoso, Sp.B) ===\n%s\n\n", string(outSup))

	// Test 3: Supervisor partial match "budi"
	resSupPartial, err := service.GetMenunggu(ctx, "budi")
	if err != nil {
		log.Fatalf("GetMenunggu supervisor partial error: %v", err)
	}
	outSupPartial, _ := json.MarshalIndent(resSupPartial, "", "  ")
	fmt.Printf("=== GET MENUNGGU (SUPERVISOR: budi) ===\n%s\n", string(outSupPartial))
}
