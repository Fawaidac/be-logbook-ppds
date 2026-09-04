package dashboard

import (
	"context"
	"math"
	"strings"
	"time"

	"be-logbook-ppds/app/pendidikan"
)

type Service interface {
	GetDashboardSummary(ctx context.Context, year int, username, name string) (*DashboardSummaryResponse, error)
	GetLaporanSummary(ctx context.Context, periode, stase string, username, name string) (*LaporanSummaryResponse, error)
	GetAdminDashboard(ctx context.Context) (*AdminDashboardResponse, error)
	GetSupervisorDashboard(ctx context.Context, supervisorName string) (*SupervisorDashboardResponse, error)
}

type service struct {
	repo           Repository
	kompetensiRepo pendidikan.KompetensiRepository
}

func NewService(repo Repository, kompetensiRepo pendidikan.KompetensiRepository) Service {
	return &service{
		repo:           repo,
		kompetensiRepo: kompetensiRepo,
	}
}

func (s *service) GetDashboardSummary(ctx context.Context, year int, username, name string) (*DashboardSummaryResponse, error) {
	if year == 0 {
		year = time.Now().Year()
	}

	totalTindakan, _ := s.repo.GetTotalTindakan(ctx, username, name)
	menungguValidasi, _ := s.repo.GetMenungguValidasi(ctx, username, name)
	perluRevisi, _ := s.repo.GetPerluRevisi(ctx, username, name)
	capaianStase, _ := s.repo.GetCapaianStase(ctx, username, name)
	upcomingJadwals, _ := s.repo.GetUpcomingJadwals(ctx, username, name)
	recentEntries, _ := s.repo.GetRecentEntries(ctx, username, name)

	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	mandiriData := make([]int, 12)
	dibimbingData := make([]int, 12)
	observasiData := make([]int, 12)
	seminarData := make([]int, 12)

	var mandiriCum, dibimbingCum, observasiCum, seminarCum int

	for m := 1; m <= 12; m++ {
		mCount, _ := s.repo.CountTindakanByMonth(ctx, year, m, "mandiri", username, name)
		dCount, _ := s.repo.CountTindakanByMonth(ctx, year, m, "dibimbing", username, name)
		oCount, _ := s.repo.CountTindakanByMonth(ctx, year, m, "observasi", username, name)
		sCount, _ := s.repo.CountKegiatanByMonth(ctx, year, m, username, name)

		mandiriCum += mCount
		dibimbingCum += dCount
		observasiCum += oCount
		seminarCum += sCount

		mandiriData[m-1] = mandiriCum
		dibimbingData[m-1] = dibimbingCum
		observasiData[m-1] = observasiCum
		seminarData[m-1] = seminarCum
	}

	return &DashboardSummaryResponse{
		TotalTindakan:    totalTindakan,
		MenungguValidasi: menungguValidasi,
		PerluRevisi:      perluRevisi,
		CapaianStase:     capaianStase,
		UpcomingJadwals:  upcomingJadwals,
		RecentEntries:    recentEntries,
		ChartData: ChartDataResponse{
			Months:    months,
			Mandiri:   mandiriData,
			Dibimbing: dibimbingData,
			Observasi: observasiData,
			Seminar:   seminarData,
		},
	}, nil
}

func (s *service) GetLaporanSummary(ctx context.Context, periode, stase string, username, name string) (*LaporanSummaryResponse, error) {
	rekapTindakan, _ := s.repo.GetRekapTindakan(ctx, periode, stase, username, name)

	var kompetensiList interface{}
	if s.kompetensiRepo != nil {
		if username != "" || name != "" {
			list, err := s.kompetensiRepo.FindByUsername(ctx, username)
			if (err != nil || len(list) == 0) && name != "" {
				list, err = s.kompetensiRepo.FindByUsername(ctx, name)
			}
			if err == nil {
				kompetensiList = list
			} else {
				kompetensiList = []string{}
			}
		} else {
			list, err := s.kompetensiRepo.FindAll(ctx)
			if err == nil {
				kompetensiList = list
			} else {
				kompetensiList = []string{}
			}
		}
	} else {
		kompetensiList = []string{}
	}

	simposium, _ := s.repo.GetKegiatanCountByKategori(ctx, "simposium", username, name)
	workshop, _ := s.repo.GetKegiatanCountByKategori(ctx, "workshop", username, name)
	multidisiplin, _ := s.repo.GetKegiatanCountByKategori(ctx, "multidisiplin", username, name)
	ilmiahLain, _ := s.repo.GetKegiatanCountByKategori(ctx, "ilmiah_lain", username, name)

	totalBimbingan, _ := s.repo.GetTotalBimbingan(ctx, username, name)
	menungguCount, _ := s.repo.GetMenungguBimbinganCount(ctx, username, name)
	disetujuiCount, _ := s.repo.GetDisetujuiBimbinganCount(ctx, username, name)

	responRate := 100
	if totalBimbingan > 0 {
		responRate = int(math.Round(float64(disetujuiCount) / float64(totalBimbingan) * 100))
	}

	rawDpjpStats, _ := s.repo.GetDPJPStats(ctx, username, name)
	for i := range rawDpjpStats {
		n := rawDpjpStats[i].Name
		cleaned := strings.TrimSpace(n)
		for _, prefix := range []string{"dr.", "Sp.OT", "Sp.B", "Sp.An-KIC", ","} {
			cleaned = strings.ReplaceAll(cleaned, prefix, "")
		}
		words := strings.Fields(cleaned)
		initials := ""
		if len(words) > 0 {
			initials += strings.ToUpper(words[0][:1])
		}
		if len(words) > 1 {
			initials += strings.ToUpper(words[1][:1])
		}
		if initials == "" {
			initials = "DP"
		}
		rawDpjpStats[i].Initials = initials

		pct := 0
		if rawDpjpStats[i].TotalLogs > 0 {
			pct = int(math.Round(float64(rawDpjpStats[i].ApprovedLogs) / float64(rawDpjpStats[i].TotalLogs) * 100))
		}
		rawDpjpStats[i].Pct = pct
	}

	return &LaporanSummaryResponse{
		RekapTindakan:      rekapTindakan,
		KompetensiList:     kompetensiList,
		SimposiumCount:     simposium,
		WorkshopCount:      workshop,
		MultidisiplinCount: multidisiplin,
		IlmiahLainCount:    ilmiahLain,
		TotalBimbingan:     totalBimbingan,
		MenungguCount:      menungguCount,
		DisetujuiCount:     disetujuiCount,
		ResponRate:         responRate,
		DPJPStats:          rawDpjpStats,
	}, nil
}

func (s *service) GetAdminDashboard(ctx context.Context) (*AdminDashboardResponse, error) {
	totalResiden, _ := s.repo.CountUsersByRole(ctx, "residen")
	totalSupervisor, _ := s.repo.CountUsersByRole(ctx, "supervisor")
	totalAdmin, _ := s.repo.CountUsersByRole(ctx, "admin")

	totalRegistrasi, _ := s.repo.CountPendingRegistrations(ctx)
	totalTindakan, _ := s.repo.GetTotalTindakan(ctx, "", "")

	registrasiList, _ := s.repo.GetPendingRegistrations(ctx, 5)

	aktifTahunIni, _ := s.repo.CountActiveResidenThisYear(ctx, time.Now().Year())
	persenAktivasi := 0
	if totalResiden > 0 {
		persenAktivasi = int(math.Round(float64(aktifTahunIni) / float64(totalResiden) * 100))
	}

	aktivitasList := make([]AdminAktivitasItem, 0, 4)
	tindakanCount, _ := s.repo.GetTotalTindakan(ctx, "", "")
	aktivitasList = append(aktivitasList, AdminAktivitasItem{Kategori: "Tindakan Klinik", Jumlah: int64(tindakanCount)})

	pendidikanCount, _ := s.repo.CountPendidikanRecords(ctx)
	aktivitasList = append(aktivitasList, AdminAktivitasItem{Kategori: "Pendidikan & Kompetensi", Jumlah: int64(pendidikanCount)})

	kegiatanCount, _ := s.repo.CountKegiatanIlmiahRecords(ctx)
	aktivitasList = append(aktivitasList, AdminAktivitasItem{Kategori: "Penelitian & Karya Ilmiah", Jumlah: int64(kegiatanCount)})

	jadwalCount, _ := s.repo.CountJadwalRecords(ctx)
	aktivitasList = append(aktivitasList, AdminAktivitasItem{Kategori: "Jadwal & Bimbingan", Jumlah: int64(jadwalCount)})

	if registrasiList == nil {
		registrasiList = []AdminRegistrasiItem{}
	}

	return &AdminDashboardResponse{
		TotalResiden:    totalResiden,
		TotalSupervisor: totalSupervisor,
		TotalAdmin:      totalAdmin,
		TotalRegistrasi: totalRegistrasi,
		TotalTindakan:   totalTindakan,
		PersenAktivasi:  persenAktivasi,
		RegistrasiList:  registrasiList,
		AktivitasList:   aktivitasList,
		RoleDistribution: RoleDistributionResponse{
			Labels: []string{"Residen", "Supervisor", "Admin"},
			Values: []int{totalResiden, totalSupervisor, totalAdmin},
		},
	}, nil
}

func (s *service) GetSupervisorDashboard(ctx context.Context, supervisorName string) (*SupervisorDashboardResponse, error) {
	now := time.Now()

	totalMenunggu, _ := s.repo.CountTindakanByStatus(ctx, "menunggu", supervisorName)
	totalRevisi, _ := s.repo.CountTindakanByStatus(ctx, "ditolak", supervisorName)
	totalDisetujui, _ := s.repo.CountTindakanByStatusInMonth(ctx, "disetujui", now.Year(), int(now.Month()), supervisorName)
	totalResiden, _ := s.repo.CountResidenBySupervisor(ctx, supervisorName)

	// Pertumbuhan disetujui: bulan ini vs bulan lalu
	prev := now.AddDate(0, -1, 0)
	disetujuiBulanLalu, _ := s.repo.CountTindakanByStatusInMonth(ctx, "disetujui", prev.Year(), int(prev.Month()), supervisorName)
	growth := 0
	if disetujuiBulanLalu > 0 {
		growth = int(math.Round(float64(totalDisetujui-disetujuiBulanLalu) / float64(disetujuiBulanLalu) * 100))
	}

	// Chart 6 bulan terakhir (label bulan dalam bahasa Indonesia)
	indoMonths := []string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"}
	labels := make([]string, 0, 6)
	chartDisetujui := make([]int, 0, 6)
	chartMenunggu := make([]int, 0, 6)
	chartRevisi := make([]int, 0, 6)

	for i := 5; i >= 0; i-- {
		m := now.AddDate(0, -i, 0)
		labels = append(labels, indoMonths[int(m.Month())-1])

		d, _ := s.repo.CountTindakanByStatusInMonth(ctx, "disetujui", m.Year(), int(m.Month()), supervisorName)
		w, _ := s.repo.CountTindakanByStatusInMonth(ctx, "menunggu", m.Year(), int(m.Month()), supervisorName)
		r, _ := s.repo.CountTindakanByStatusInMonth(ctx, "ditolak", m.Year(), int(m.Month()), supervisorName)

		chartDisetujui = append(chartDisetujui, d)
		chartMenunggu = append(chartMenunggu, w)
		chartRevisi = append(chartRevisi, r)
	}

	residenList, _ := s.repo.GetResidenActivitySummary(ctx, 5, supervisorName)
	pendingList, _ := s.repo.GetPendingTindakanQueue(ctx, 5, supervisorName)

	return &SupervisorDashboardResponse{
		TotalMenunggu:   totalMenunggu,
		TotalDisetujui:  totalDisetujui,
		TotalRevisi:     totalRevisi,
		TotalResiden:    totalResiden,
		DisetujuiGrowth: growth,
		Chart: SupervisorChartDataResponse{
			Labels:    labels,
			Disetujui: chartDisetujui,
			Menunggu:  chartMenunggu,
			Revisi:    chartRevisi,
		},
		ResidenList: residenList,
		PendingList: pendingList,
	}, nil
}
